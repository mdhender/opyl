package mapgen

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mdhender/opyl/internal/infra/prng"
	"github.com/mdhender/ottomap"
	"github.com/mdhender/ottomap/hex"
	"github.com/mdhender/ottomap/wog"
)

// Options configures a Generator.
type Options struct {
	// InputDir is the directory containing the input files
	// (Map, Regions, Land, Cities, randseed). Defaults to ".".
	InputDir      string
	InputMap      string // Defaults to "olympia-g3-source.wxx".
	InputRandSeed string // Defaults to "randseed".
	// OutputDir is the directory the output files (loc, gate, road,
	// randseed) are written to. Defaults to ".".
	OutputDir      string
	OutputMap      string // Defaults to "olympia-g3-map.wxx".
	OutputRandSeed string // Defaults to "randseed".
	// Log receives the diagnostic output the C program writes to stderr.
	// Defaults to os.Stderr. Use io.Discard to silence it.
	Log io.Writer
	// Origin holds the x,y from the input map and the q,r from the translated map
	Origin Origin
}

// Generator holds the entire state of a single map-generation run. It is the
// direct analogue of the global state in the legacy C program.
type Generator struct {
	RNG  *RNG
	PRNG *prng.PRNG

	InputDir       string
	InputMap       string
	InputRandSeed  string
	OutputDir      string
	OutputMap      string
	OutputRandSeed string
	Log            io.Writer

	Origin Origin

	allocFlag [MaxBox]bool
	dirVector [MaxDir]int

	InsideNames     [MaxInside]string
	InsideList      [MaxInside][]*Tile
	insideGatesTo   [MaxInside]int
	insideGatesFrom [MaxInside]int
	insideNumCities [MaxInside]int
	InsideTop       int

	MaxColUsed int
	MaxRowUsed int

	WaterCount   int
	LandCount    int
	UnknownCount int
	NumIslands   int

	Map    [MaxRow][MaxCol]*Tile
	HexMap map[hex.Axial]*Tile

	Subloc    []*Tile // 1-indexed; index 0 unused
	TopSubloc int

	// output buffers
	loc  bytes.Buffer
	gate bytes.Buffer
	road bytes.Buffer

	// Cities file reader state (matches the C static FILE *fp)
	citiesReader *bufio.Reader
	citiesOpened bool
	citiesFailed bool
	cityNames    *CityNameGenerator

	// bridge_corner_sup / bridge_map_hole_sup road name counters
	cornerRoadNameCnt int
	holeRoadNameCnt   int
}

// New returns a Generator configured with opts.
func New(opts Options) *Generator {
	g := &Generator{
		RNG:            NewRNG(),
		InputDir:       opts.InputDir,
		InputMap:       opts.InputMap,
		InputRandSeed:  opts.InputRandSeed,
		OutputDir:      opts.OutputDir,
		OutputMap:      opts.OutputMap,
		OutputRandSeed: opts.OutputRandSeed,
		Log:            opts.Log,
		Subloc:         make([]*Tile, MaxSubloc),
		HexMap:         map[hex.Axial]*Tile{},
	}
	if g.InputDir == "" {
		g.InputDir = "."
	}
	if g.InputMap == "" {
		g.InputMap = "olympia-g3-source.wxx"
	}
	if g.InputRandSeed == "" {
		g.InputRandSeed = "randseed"
	}
	if g.OutputDir == "" {
		g.OutputDir = "."
	}
	if g.OutputMap == "" {
		g.OutputMap = "olympia-g3-map.wxx"
	}
	if g.OutputRandSeed == "" {
		g.OutputRandSeed = "randseed"
	}
	if g.Log == nil {
		g.Log = os.Stderr
	}
	return g
}

func (g *Generator) logf(format string, a ...any) {
	fmt.Fprintf(g.Log, format, a...)
}

func (g *Generator) rnd(low, high int) int { return g.RNG.Rnd(low, high) }

func (g *Generator) inPath(name string) string  { return filepath.Join(g.InputDir, name) }
func (g *Generator) outPath(name string) string { return filepath.Join(g.OutputDir, name) }

// Run executes the full map-generation pipeline and writes the loc, gate,
// road, and randseed output files. It mirrors the C main() function.
func (g *Generator) Run() error {
	g.dirAssert()

	if err := g.RNG.LoadSeed(g.inPath(g.InputRandSeed)); err != nil {
		g.logf("%s could not be opened.\n", g.inPath(g.InputRandSeed))
	}

	if err := g.readMap(); err != nil {
		return err
	}
	g.fixTerrainLand()
	if err := g.setRegions(); err != nil {
		return err
	}
	if err := g.setProvinceClumps(); err != nil {
		return err
	}
	g.unnamedProvinceClumps()
	g.makeIslands()
	g.makeGraveyards()
	g.placeSublocations()
	g.makeGates()
	g.makeRoads()

	g.printMap()
	g.printSublocs()
	g.dumpContinents()
	g.countCities()
	g.countContinents()
	g.countSublocs()
	g.countSublocCoverage()
	g.dumpRoads()
	g.dumpGates()

	if err := g.writeFile("loc", g.loc.Bytes()); err != nil {
		return err
	}
	if err := g.writeFile("road", g.road.Bytes()); err != nil {
		return err
	}
	if err := g.writeFile("gate", g.gate.Bytes()); err != nil {
		return err
	}

	g.countTiles()
	g.logf("\nhighest province = %d\n\n", g.Map[g.MaxRowUsed][g.MaxColUsed].Region)

	return g.RNG.SaveSeed(g.outPath(g.OutputRandSeed))
}

// RunHex executes the full map-generation pipeline and writes the loc, gate,
// road, and randseed output files. It mirrors the C main() function.
func (g *Generator) RunHex() error {
	g.LandCount, g.WaterCount, g.UnknownCount = 0, 0, 0

	g.dirAssert()

	if err := g.RNG.LoadSeed(g.inPath(g.InputRandSeed)); err != nil {
		g.logf("%s could not be opened.\n", g.inPath(g.InputRandSeed))
	}

	// use the 16 bytes to seed our PRNG
	seed1 := binary.LittleEndian.Uint64(g.RNG.Digest[0:8])
	seed2 := binary.LittleEndian.Uint64(g.RNG.Digest[8:16])
	g.PRNG = prng.NewFromSeed(seed1, seed2)
	g.cityNames = NewCityNameGenerator(g.PRNG)

	if err := g.readHexMap(); err != nil {
		return err
	}
	g.logf("\nland / water / unknown = %6d / %d6 / %6d (%6d)\n\n", g.LandCount, g.WaterCount, g.UnknownCount, g.LandCount+g.WaterCount+g.UnknownCount)

	return nil
}

func (g *Generator) writeFile(name string, data []byte) error {
	return os.WriteFile(g.outPath(name), data, 0644)
}

// ---------------------------------------------------------------------------
// entity allocation
// ---------------------------------------------------------------------------

func (g *Generator) rndAllocNum(low, high int) int {
	n := g.rnd(low, high)

	for i := n; i <= high; i++ {
		if !g.allocFlag[i] {
			g.allocFlag[i] = true
			return i
		}
	}
	for i := low; i < n; i++ {
		if !g.allocFlag[i] {
			g.allocFlag[i] = true
			return i
		}
	}

	g.logf("rnd_alloc_num(%d,%d) failed\n", low, high)
	return -1
}

// ---------------------------------------------------------------------------
// region / coordinate helpers
// ---------------------------------------------------------------------------

func rcToRegion(row, col int) int {
	return 10000 + (row * 100) + col
}

func regionRow(where int) int { return (where / 100) % 100 }
func regionCol(where int) int { return where % 100 }

func (g *Generator) dirAssert() {
	if rcToRegion(1, 1) != 10101 || regionRow(10101) != 1 || regionCol(10101) != 1 {
		panic("dir_assert failed")
	}
	if rcToRegion(99, 99) != 19999 {
		panic("dir_assert failed")
	}
}

// ---------------------------------------------------------------------------
// adjacency
// ---------------------------------------------------------------------------

func (g *Generator) adjacentTileSup(row, col, dir int) *Tile {
	switch dir {
	case DirN:
		row--
	case DirNE:
		row--
		col++
	case DirE:
		col++
	case DirSE:
		row++
		col++
	case DirS:
		row++
	case DirSW:
		row++
		col--
	case DirW:
		col--
	case DirNW:
		row--
		col--
	default:
		panic(fmt.Sprintf("location_direction: bad dir %d", dir))
	}

	if col < 0 {
		col = g.MaxColUsed
	}
	if col > g.MaxColUsed {
		col = 0
	}

	if row < 0 || row > 99 || col < 0 || col > 99 {
		return nil
	}

	return g.Map[row][col]
}

func (g *Generator) provDest(t *Tile, dir int) int {
	row, col := t.Row, t.Col

	switch dir {
	case DirN:
		row--
	case DirE:
		col++
	case DirS:
		row++
	case DirW:
		col--
	default:
		panic(fmt.Sprintf("location_direction: bad dir %d", dir))
	}

	if row < 0 || row > 99 {
		return 0
	}
	if col < 0 {
		col = g.MaxColUsed
	}
	if col > g.MaxColUsed {
		col = 0
	}

	if g.Map[row][col] == nil {
		return 0
	}
	return g.Map[row][col].Region
}

func (g *Generator) randomizeDirVector() {
	g.dirVector[0] = 0
	for i := 1; i < MaxDir; i++ {
		g.dirVector[i] = i
	}
	for i := 1; i < MaxDir; i++ {
		swap := g.rnd(i, MaxDir-1)
		if i != swap {
			g.dirVector[i], g.dirVector[swap] = g.dirVector[swap], g.dirVector[i]
		}
	}
}

func (g *Generator) adjacentTileWater(row, col int) *Tile {
	var p *Tile
	g.randomizeDirVector()

	i := 1
	for !(p != nil && p.Terrain == TerrOcean) && i < MaxDir {
		p = g.adjacentTileSup(row, col, g.dirVector[i])
		i++
	}
	if i < MaxDir {
		return p
	}
	return nil
}

func (g *Generator) adjacentTileTerr(row, col int) *Tile {
	var p *Tile
	g.randomizeDirVector()

	i := 1
	for !(p != nil && p.Terrain != TerrLand && p.Terrain != TerrOcean) && i < MaxDir {
		p = g.adjacentTileSup(row, col, g.dirVector[i])
		i++
	}
	if i < MaxDir {
		return p
	}
	return nil
}

func (g *Generator) isPortCity(row, col int) bool {
	n := g.adjacentTileSup(row, col, DirN)
	s := g.adjacentTileSup(row, col, DirS)
	e := g.adjacentTileSup(row, col, DirE)
	w := g.adjacentTileSup(row, col, DirW)

	if (n != nil && n.Terrain == TerrOcean) ||
		(s != nil && s.Terrain == TerrOcean) ||
		(e != nil && e.Terrain == TerrOcean) ||
		(w != nil && w.Terrain == TerrOcean) {
		return true
	}
	return false
}

// ---------------------------------------------------------------------------
// Cities file
// ---------------------------------------------------------------------------

// nextCityName mimics the C create_a_city() use of getlin(fp): it returns the
// next line of the Cities file, or ("", false) once the file is exhausted or
// cannot be opened.
func (g *Generator) nextCityName() (string, bool) {
	if g.citiesFailed {
		return "", false
	}
	if !g.citiesOpened {
		g.citiesOpened = true
		f, err := os.Open(g.inPath("Cities"))
		if err != nil {
			g.logf("can't open Cities: %v\n", err)
			g.citiesFailed = true
			return "", false
		}
		g.citiesReader = bufio.NewReader(f)
	}

	line, err := g.citiesReader.ReadString('\n')
	if line == "" && err != nil {
		return "", false
	}
	line = strings.TrimRight(line, "\n")
	return line, true
}

// ---------------------------------------------------------------------------
// subloc / city / building creation
// ---------------------------------------------------------------------------

func (g *Generator) createASubloc(row, col, hidden, kind int) int {
	g.TopSubloc++
	if g.TopSubloc >= MaxSubloc {
		panic("top_subloc overflow")
	}

	t := &Tile{}
	g.Subloc[g.TopSubloc] = t
	if kind == TerrCity {
		t.Region = g.rndAllocNum(CityLow, CityHigh)
	} else {
		t.Region = g.rndAllocNum(SublocLow, SublocHigh)
	}
	t.Inside = g.Map[row][col].Region
	t.Row = row
	t.Col = col
	t.Hidden = hidden
	t.Terrain = kind
	t.Depth = 3

	if kind == TerrCity {
		g.Map[row][col].City = 2
	}

	g.Map[row][col].Subs = append(g.Map[row][col].Subs, t.Region)

	return g.TopSubloc
}

func (g *Generator) createACity(row, col int, name string, hasName bool, major int) int {
	if !hasName {
		if n, ok := g.nextCityName(); ok {
			name = n
		}
	}

	n := g.createASubloc(row, col, 0, TerrCity)
	g.Subloc[n].Name = name
	g.Subloc[n].MajorCity = major
	return n
}

func (g *Generator) createAGraveyard(row, col int) {
	hidden := g.rnd(0, 1)
	n := g.createASubloc(row, col, hidden, TerrGrave)
	s := g.nameGuild(TerrGrave)
	if s != "" {
		g.Subloc[n].Name = s
	}
}

// ---------------------------------------------------------------------------
// map reading
// ---------------------------------------------------------------------------

func (g *Generator) readHexMap() error {
	// use ottomap to load the map file
	fmt.Printf("ottomap: %s\n", ottomap.Version().Short())
	fmt.Printf("  input: %s\n", g.inPath(g.InputMap))
	fp, err := os.Open(g.inPath(g.InputMap))
	if err != nil {
		return fmt.Errorf("can't open %s: %w", g.InputMap, err)
	}
	defer func() {
		_ = fp.Close()
	}()
	om, ov, err := wog.Read(fp)
	if err != nil {
		fmt.Printf("\t%v\n", err)
		return err
	}
	fmt.Printf("ottomap: %s\n", ov.String())
	amin, amax, explicit := om.Bounds()
	fmt.Printf("bounds: %v..%v (explicit=%v)\n", amin, amax, explicit)
	omin, omax, explicit := om.BoundsOffset()
	cols := omax.Col - omin.Col + 1
	rows := omax.Row - omin.Row + 1
	fmt.Printf("%q %dx%d\n", om.Layout(), cols, rows)
	fmt.Printf("\t%8d tiles high\n", rows)
	fmt.Printf("\t%8d tiles wide\n", cols)

	// range over the tiles in the source map
	for coord, tile := range om.Tiles() {
		if tile.Terrain == "Blank" {
			// hole in the map
			continue
		}

		rc := coord.ToOffset(hex.OddQ)
		t := &Tile{
			Coords: coord,
			Row:    rc.Row,
			Col:    rc.Col,
			Region: rcToRegion(rc.Row, rc.Col),
			Depth:  2,
		}
		g.HexMap[coord] = t

		color := 0
		terrain := 0

		// map the Worldographer terrain to a map glyph
		var glyph byte
		switch tile.Terrain {
		case "Classic/Water Ocean Deep":
			glyph = ';'
			t.SeaLane = 1
			terrain = TerrOcean
			color = 1
			g.WaterCount++
		case "Classic/Water Ocean":
			glyph = ','
			terrain = TerrOcean
			color = 1
			g.WaterCount++
		case "Classic/Water Sea Deep":
			glyph = ':'
			t.SeaLane = 1
			terrain = TerrOcean
			color = 2
			g.WaterCount++
		case "Classic/Water Sea":
			glyph = '.'
			terrain = TerrOcean
			color = 2
			g.WaterCount++
		case "Classic/Water Kelp":
			glyph = '~'
			t.SeaLane = 1
			terrain = TerrOcean
			color = 3
			g.WaterCount++
		case "Classic/Water Kelp Heavy":
			glyph = ' '
			terrain = TerrOcean
			color = 3
			g.WaterCount++
		case "Classic/Water Shoals":
			glyph = '"'
			t.SeaLane = 1
			terrain = TerrOcean
			color = 4
			g.WaterCount++
		case "Classic/Water Reef":
			glyph = '\''
			terrain = TerrOcean
			color = 4
			g.WaterCount++
		case "Classic/Flat Grassland":
			glyph = 'p'
			color = 5
			terrain = TerrPlain
			g.LandCount++
		case "Classic/Hills":
			glyph = 'P'
			color = 6
			terrain = TerrPlain
			g.LandCount++
		case "Classic/Flat Desert Rocky":
			glyph = 'd'
			color = 7
			terrain = TerrDesert
			g.LandCount++
		case "Classic/Flat Desert Sandy":
			glyph = 'D'
			color = 8
			terrain = TerrDesert
			g.LandCount++
		case "Classic/Mountain":
			glyph = 'm'
			color = 9
			terrain = TerrMountain
			g.LandCount++
		case "Classic/Mountains":
			glyph = 'M'
			color = 10
			terrain = TerrMountain
			g.LandCount++
		case "Classic/Flat Swamp":
			glyph = 's'
			color = 11
			terrain = TerrSwamp
			g.LandCount++
		case "Classic/Flat Wetlands Jungle":
			glyph = 'S'
			color = 12
			terrain = TerrSwamp
			g.LandCount++
		case "Classic/Flat Forest Deciduous":
			glyph = 'f'
			color = 13
			terrain = TerrForest
			g.LandCount++
		case "Classic/Flat Forest Deciduous Heavy":
			glyph = 'F'
			color = 14
			terrain = TerrForest
			g.LandCount++
		case "Classic/Underdark Open":
			glyph = 'o'
			switch g.PRNG.Roll(1, 10) {
			case 1, 2, 3:
				terrain = TerrForest
				g.LandCount++
			case 4, 5, 6:
				terrain = TerrPlain
				g.LandCount++
			case 7, 8:
				terrain = TerrMountain
				g.LandCount++
			case 9:
				terrain = TerrSwamp
				g.LandCount++
			case 10:
				terrain = TerrDesert
				g.LandCount++
			}
			color = -1
		case "Classic/Mountains Forest Jungle":
			glyph = '^'
			color = 9
			terrain = TerrMountain
			t.UldimFlag = 1
			t.RegionBoundary = 1
			g.LandCount++
		case "Classic/Mountains Forest Dead":
			glyph = 'v'
			color = 9
			terrain = TerrMountain
			t.UldimFlag = 2
			t.RegionBoundary = 1
			g.LandCount++
		case "Classic/Mountains Forest Deciduous":
			glyph = '{'
			color = 16
			terrain = TerrMountain
			t.UldimFlag = 3
			t.Name = "Uldim pass"
			t.RegionBoundary = 1
			g.LandCount++
		case "Classic/Mountains Forest Evergreen":
			glyph = '}'
			color = 16
			terrain = TerrMountain
			t.UldimFlag = 4
			t.Name = "Uldim pass"
			t.RegionBoundary = 1
			g.LandCount++
		case "Classic/Flat Marsh":
			glyph = ']'
			terrain = TerrSwamp
			t.SummerbridgeFlag = 1
			t.Name = "Summerbridge"
			t.RegionBoundary = 1
			g.LandCount++
		case "Classic/Flat Moor":
			glyph = '['
			terrain = TerrSwamp
			t.SummerbridgeFlag = 2
			t.Name = "Summerbridge"
			t.RegionBoundary = 1
			g.LandCount++
		case "Classic/Mountain Snowcapped":
			glyph = '0'
			terrain = TerrMountain
			color = -1
			t.Name = "Mt. Olympus"
			g.LandCount++

		case "Classic/Flat Farmland":
			g.UnknownCount++
		case "Classic/Flat Forest Jungle":
			g.UnknownCount++
		case "Classic/Flat Forest Jungle Heavy":
			g.UnknownCount++
		case "Classic/Hills Forest Deciduous":
			g.UnknownCount++
		case "Classic/Hills Forest Jungle":
			g.UnknownCount++
		case "Classic/Other Badlands":
			g.UnknownCount++
		case "Classic/Natural Volcano Dormant":
			g.UnknownCount++

		default:
			g.logf("unknown terrain %q\n", tile.Terrain)
			panic("readHexMap: unknown terrain")
		}

		t.SaveChar = glyph
		t.Terrain = terrain
		t.Color = color
	}

	// range over the features in the source map
	for _, v := range om.Features() {
		t, ok := g.HexMap[v.Location]
		if !ok {
			panic(fmt.Sprintf("assert(city.location == %q)", v.Location))
		}

		isSafeHaven, isMajorCity := false, false
		switch v.Kind {
		case "Classic/Building Cathedral":
			t.SafeHaven = 1
			isSafeHaven, isMajorCity = true, true
		case "Classic/Military Castle":
			isSafeHaven, isMajorCity = false, true
		case "Classic/Building Church":
			t.SafeHaven = 1
			isSafeHaven, isMajorCity = true, false
		case "Classic/Military Camp":
			isSafeHaven, isMajorCity = false, false
		default:
			panic(fmt.Sprintf("assert(kind != %q)", v.Kind))
		}
		cityName := v.Label
		if cityName == "" { // randomly assign a name
			cityName = g.cityNames.Name()
		}

		fmt.Printf("%-30s, %-65q safe %-8v major %-8v (%03d.%03d) %v\n", cityName, v.Kind, isSafeHaven, isMajorCity, t.Col, t.Row, t.Coords)
	}

	return nil
}

func (g *Generator) readMap() error {
	f, err := os.Open(g.inPath("Map"))
	if err != nil {
		return fmt.Errorf("can't open Map: %w", err)
	}
	defer func() {
		_ = f.Close()
	}()

	r := bufio.NewReader(f)
	row := 0

	for {
		line, rerr := r.ReadString('\n')
		if line == "" {
			break
		}

		for col := 0; col < len(line) && line[col] != '\n'; col++ {
			ch := line[col]
			if ch == '#' { // hole in map
				continue
			}

			if row > g.MaxRowUsed {
				g.MaxRowUsed = row
			}
			if col > g.MaxColUsed {
				g.MaxColUsed = col
			}

			t := &Tile{
				Row:    row,
				Col:    col,
				Region: rcToRegion(row, col),
				Depth:  2,
			}
			g.Map[row][col] = t

			color := 0
			terrain := 0

			switch ch {
			case ';':
				t.SeaLane = 1
				terrain = TerrOcean
				color = 1
			case ',':
				terrain = TerrOcean
				color = 1
			case ':':
				t.SeaLane = 1
				terrain = TerrOcean
				color = 2
			case '.':
				terrain = TerrOcean
				color = 2
			case '~':
				t.SeaLane = 1
				terrain = TerrOcean
				color = 3
			case ' ':
				terrain = TerrOcean
				color = 3
			case '"':
				t.SeaLane = 1
				terrain = TerrOcean
				color = 4
			case '\'':
				terrain = TerrOcean
				color = 4
			case 'p':
				color = 5
				terrain = TerrPlain
			case 'P':
				color = 6
				terrain = TerrPlain
			case 'd':
				color = 7
				terrain = TerrDesert
			case 'D':
				color = 8
				terrain = TerrDesert
			case 'm':
				color = 9
				terrain = TerrMountain
			case 'M':
				color = 10
				terrain = TerrMountain
			case 's':
				color = 11
				terrain = TerrSwamp
			case 'S':
				color = 12
				terrain = TerrSwamp
			case 'f':
				color = 13
				terrain = TerrForest
			case 'F':
				color = 14
				terrain = TerrForest
			case 'o':
				switch g.rnd(1, 10) {
				case 1, 2, 3:
					terrain = TerrForest
				case 4, 5, 6:
					terrain = TerrPlain
				case 7, 8:
					terrain = TerrMountain
				case 9:
					terrain = TerrSwamp
				case 10:
					terrain = TerrDesert
				}
				color = -1
			case '^':
				color = 9
				terrain = TerrMountain
				t.UldimFlag = 1
				t.RegionBoundary = 1
			case 'v':
				color = 9
				terrain = TerrMountain
				t.UldimFlag = 2
				t.RegionBoundary = 1
			case '{':
				color = 16
				terrain = TerrMountain
				t.UldimFlag = 3
				t.Name = "Uldim pass"
				t.RegionBoundary = 1
			case '}':
				color = 16
				terrain = TerrMountain
				t.UldimFlag = 4
				t.Name = "Uldim pass"
				t.RegionBoundary = 1
			case ']':
				terrain = TerrSwamp
				t.SummerbridgeFlag = 1
				t.Name = "Summerbridge"
				t.RegionBoundary = 1
			case '[':
				terrain = TerrSwamp
				t.SummerbridgeFlag = 2
				t.Name = "Summerbridge"
				t.RegionBoundary = 1
			case 'O':
				terrain = TerrMountain
				color = -1
				t.Name = "Mt. Olympus"
			case '1':
				terrain = TerrForest
				color = 19
				t.SafeHaven = 1
				n := g.createACity(row, col, "Drassa", true, 1)
				g.Subloc[n].SafeHaven = 1
				g.logf("Start city #%c %s at (%d,%d)\n", ch, g.Subloc[n].Name, row, col)
			case '2':
				terrain = TerrForest
				color = 19
				t.SafeHaven = 1
				n := g.createACity(row, col, "Rimmon", true, 1)
				g.Subloc[n].SafeHaven = 1
				g.logf("Start city #%c %s at (%d,%d)\n", ch, g.Subloc[n].Name, row, col)
			case '3':
				terrain = TerrForest
				color = 19
				t.SafeHaven = 1
				n := g.createACity(row, col, "Harn", true, 1)
				g.Subloc[n].SafeHaven = 1
				g.logf("Start city #%c %s at (%d,%d)\n", ch, g.Subloc[n].Name, row, col)
			case '4':
				terrain = TerrForest
				color = 19
				t.SafeHaven = 1
				n := g.createACity(row, col, "Imperial City", true, 1)
				g.Subloc[n].SafeHaven = 1
				g.logf("Imperical City #%c %s at (%d,%d)\n", ch, g.Subloc[n].Name, row, col)
			case '5':
				terrain = TerrForest
				color = 19
				t.SafeHaven = 1
				n := g.createACity(row, col, "Port Aurnos", true, 1)
				g.Subloc[n].SafeHaven = 1
				g.logf("Start city #%c %s at (%d,%d)\n", ch, g.Subloc[n].Name, row, col)
			case '6':
				terrain = TerrForest
				color = 19
				t.SafeHaven = 1
				n := g.createACity(row, col, "Greyfell", true, 1)
				g.Subloc[n].SafeHaven = 1
				g.logf("Start city #%c %s at (%d,%d)\n", ch, g.Subloc[n].Name, row, col)
			case '7':
				terrain = TerrForest
				color = 19
				t.SafeHaven = 1
				n := g.createACity(row, col, "Yellowleaf", true, 1)
				g.Subloc[n].SafeHaven = 1
				g.logf("Start city #%c %s at (%d,%d)\n", ch, g.Subloc[n].Name, row, col)
			case '8':
				terrain = TerrForest
				color = 19
				n := g.createACity(row, col, "Golden City", true, 1)
				g.logf("Golden City #%c %s at (%d,%d)\n", ch, g.Subloc[n].Name, row, col)
			case '9', '0':
				terrain = TerrForest
				color = 19
				t.SafeHaven = 1
				n := g.createACity(row, col, "", false, 1)
				g.Subloc[n].SafeHaven = 1
				g.logf("Start city #%c %s at (%d,%d)\n", ch, g.Subloc[n].Name, row, col)
			case '*':
				terrain = TerrLand
				g.createACity(row, col, "", false, 1)
			case '%':
				terrain = TerrLand
				g.createACity(row, col, "", false, 0)
			default:
				g.logf("unknown terrain %c\n", ch)
				panic("read_map: unknown terrain")
			}

			t.SaveChar = ch
			t.Terrain = terrain
			t.Color = color

			if terrain == TerrWater || terrain == TerrOcean {
				g.WaterCount++
			} else {
				g.LandCount++
			}
		}

		row++
		if rerr != nil {
			break
		}
	}

	return nil
}

// ---------------------------------------------------------------------------
// terrain inference
// ---------------------------------------------------------------------------

func (g *Generator) fixTerrainLand() {
	for row := 0; row < MaxRow; row++ {
		for col := 0; col < MaxCol; col++ {
			t := g.Map[row][col]
			if t == nil || t.Terrain != TerrLand {
				continue
			}
			p := g.adjacentTileTerr(row, col)
			if p != nil && p.Terrain != TerrLand && p.Terrain != TerrOcean {
				t.Terrain = p.Terrain
				t.Color = p.Color
			} else {
				g.logf("fix_terrain: could not infer type of (%d,%d)\n", row, col)
				g.logf("    assuming 'forest'\n")
				t.Terrain = TerrForest
			}
		}
	}
}
