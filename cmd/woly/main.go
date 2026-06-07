// Command woly builds opyl's authored map artifact from a Worldographer
// (.wxx) source. It is a separate composition root from cmd/opyl: it emits
// the single versioned JSON document that the engine's planned MapSource
// port loads as immutable input (docs/adr ADR 0004).
//
// SOUSA: this file is Runtime. It reuses internal/domain for coordinate
// math — the single source of truth for axial (q, r) — and will wire a
// Worldographer-parsing infra adapter once one exists under internal/infra.
// Keep wiring and flag parsing here; no map-import logic in this file.
//
// The full importer is not built yet. This scaffold parses the pin flags,
// resolves them through the domain hex types, and reports what it would do.
//
// Flags are parsed with peterbourgon/ff: long names take a double dash
// (--source), matching the invocation forms shown in the reference docs.
//
// See docs/content/reference/model/map-artifact.md for the artifact shape
// and the offset→axial conversion woly performs.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/mdhender/opyl/internal/domain"
	"github.com/mdhender/opyl/internal/infra/prng"
	"github.com/mdhender/ottomap"
	"github.com/mdhender/ottomap/hex"
	"github.com/mdhender/ottomap/wog"
	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffhelp"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "woly:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := ff.NewFlagSet("woly")
	var (
		source        = fs.StringLong("source", "", "path to the Worldographer .wxx source (required)")
		key           = fs.StringLong("key", "", "path to the opyl-key.json file (required)")
		out           = fs.StringLong("out", "", "path to write the map artifact JSON (default: stdout)")
		showHistogram = fs.BoolLong("show-histogram", "show histogram of terrain names")
		xy            coordPair
		qr            coordPair
	)
	// pinFlag registers a coordPair flag with an explicit placeholder; like
	// ff's typed definers, a static misconfiguration here is a programming bug,
	// so panic rather than thread an error through run.
	pinFlag := func(long, placeholder, usage string, v *coordPair) {
		if _, err := fs.AddFlag(ff.FlagConfig{LongName: long, Value: v, Placeholder: placeholder, Usage: usage}); err != nil {
			panic(err)
		}
	}
	pinFlag("x-y", "X,Y", "Worldographer hex to pin (e.g. 5,8)", &xy)
	pinFlag("q-r", "Q,R", "axial coordinate to pin the hex to (e.g. 0,0)", &qr)
	if err := ff.Parse(fs, args); err != nil {
		if errors.Is(err, ff.ErrHelp) {
			usage(os.Stderr, fs)
			return nil
		}
		return err
	}
	if *source == "" {
		usage(os.Stderr, fs)
		return fmt.Errorf("missing required --source")
	}
	if *key == "" {
		usage(os.Stderr, fs)
		return fmt.Errorf("missing required --key")
	}

	artifact := Artifact{
		SchemaVersion: CurrentSchemaVersion,
	}

	// The pin re-centres the world in axial space: convert both the pinned
	// Worldographer hex and the target coordinate to axial, then translate
	// every hex by their difference. Raw offset subtraction is wrong because
	// r depends nonlinearly on the column (map-artifact.md).
	delta := domain.Coord{Q: 0, R: 0}
	if xy.set != qr.set {
		return fmt.Errorf("--x-y and --q-r must be given together")
	}
	if xy.set {
		pinHex := domain.Offset{Col: xy.a, Row: xy.b}
		pinTo := domain.Coord{Q: qr.a, R: qr.b}
		delta = pinTo.Sub(pinHex.ToAxial())
	}
	artifact.Origin = Origin{
		WXY: OffsetXY{X: xy.a, Y: xy.b},
		QR:  AxialQR{Q: qr.a, R: qr.b},
	}

	dest := "stdout"
	if *out != "" {
		dest = *out
	}
	_, _ = fmt.Fprintf(os.Stderr, "woly: source %s . %s; axial pin offset %s\n", *source, dest, delta)

	// load the key file
	mapKeys := MapKey{
		Regions:  map[string]*RegionData{},
		Terrains: map[string]*TerrainData{},
	}
	if data, err := os.ReadFile(*key); err != nil {
		fmt.Printf("\t%v\n", err)
		return err
	} else if err = json.Unmarshal(data, &mapKeys); err != nil {
		fmt.Printf("\t%v\n", err)
		return err
	} else { // normalize
		for k, v := range mapKeys.Regions {
			v.Name = k
			mapKeys.RegionList = append(mapKeys.RegionList, v)
		}
		sort.Slice(mapKeys.RegionList, func(i, j int) bool {
			return mapKeys.RegionList[i].Name < mapKeys.RegionList[j].Name
		})
		for k, v := range mapKeys.Terrains {
			v.Name = k
			mapKeys.TerrainList = append(mapKeys.TerrainList, v)
		}
		sort.Slice(mapKeys.TerrainList, func(i, j int) bool {
			return mapKeys.TerrainList[i].Name < mapKeys.TerrainList[j].Name
		})
	}

	for _, v := range mapKeys.Regions {
		axial := (domain.Offset{Col: v.Coords.X, Row: v.Coords.Y}).ToAxial().Add(delta)
		artifact.Regions = append(artifact.Regions, Region{
			ID:   fmt.Sprintf("%d,%d", axial.Q, axial.R),
			Q:    axial.Q,
			R:    axial.R,
			Name: v.Name,
			Kind: "normal",
		})
	}
	sort.Slice(artifact.Regions, func(i, j int) bool {
		if artifact.Regions[i].Q < artifact.Regions[j].Q {
			return true
		}
		if artifact.Regions[i].Q == artifact.Regions[j].Q {
			return artifact.Regions[i].R < artifact.Regions[j].R
		}
		return false
	})

	// TODO: parse the .wxx source behind an internal/infra adapter, apply
	// `delta` to every province's axial coordinate, mint entity numbers for
	// sub-locations, and emit the deterministic artifact to dest.

	fmt.Printf("info:\t%s\n", *source)

	// use ottomap to load the map file
	fmt.Printf("ottomap: %s\n", ottomap.Version().Short())
	fp, err := os.Open(*source)
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

	// force a seed; later the GM will be able to specify it.
	rnd := prng.NewFromSeed(42, 43)

	// create a histogram for the terrain types in the source
	terrainHistogram := map[string]int{}

	// create provinces from the map data
	provinces := map[hex.Axial]*InputProvince{}
	for coord, tile := range om.Tiles() {
		terrainHistogram[string(tile.Terrain)]++
		tileData := mapKeys.Terrains[string(tile.Terrain)]
		tileData.Count++
		provinces[coord] = DecodeTile(tile, coord, tileData.Glyph, rnd)
	}

	// flood fill regions.
	landCount, waterCount := 0, 0
	for _, v := range mapKeys.RegionList {
		// find the province at the center of the region
		pr, ok := provinces[v.ID]
		if !ok {
			// region has no tile on the map!
			fmt.Printf("error: region %q: tile %q: does not exist\n", v.Name, v.ID)
			panic("region: missing tile")
		}
		// todo: what is inside supposed to represent?
		const inside = true
		if pr.Terrain == domain.TerrOcean {
			waterCount++
			floodWaterInside(v.Name, provinces, v.ID, inside)
		} else {
			landCount++
			floodLandInside(v.Name, provinces, v.ID, inside)
		}
	}

	// sort the names to get deterministic output
	terrainNames := make([]string, 0, len(terrainHistogram))
	for name := range terrainHistogram {
		terrainNames = append(terrainNames, name)
	}
	sort.Strings(terrainNames)
	// display the histogram
	if *showHistogram {
		for _, name := range terrainNames {
			tileData, _ := mapKeys.Terrains[name]
			fmt.Printf("%-60s %6d maps to %q\n", name, terrainHistogram[name], tileData.Kind)
		}
	}
	fmt.Printf("\t%8d terrain tiles defined\n", len(terrainNames))

	// convert the provinces to artifact provinces
	for _, pr := range provinces {
		province := Province{
			Q:       pr.ID.Q,
			R:       pr.ID.R,
			Terrain: pr.Terrain.String(),
		}
		artifact.Provinces = append(artifact.Provinces, province)
	}
	sort.Slice(artifact.Provinces, func(i, j int) bool {
		if artifact.Provinces[i].Q < artifact.Provinces[j].Q {
			return true
		}
		if artifact.Provinces[i].Q == artifact.Provinces[j].Q {
			return artifact.Provinces[i].R < artifact.Provinces[j].R
		}
		return false
	})

	if buf, err := json.MarshalIndent(artifact, "", "  "); err != nil {
		return err
	} else if *out == "" {
		fmt.Printf("%s\n", buf)
	} else if err = os.WriteFile(*out, buf, 0644); err != nil {
		return err
	}

	return fmt.Errorf("Worldographer .wxx import is not yet implemented")
}

// coordPair is the flag.Value behind the pin flags --x-y and --q-r. It parses
// an "A,B" pair of integers and records whether the flag was set, so run can
// enforce that the two pin flags are supplied together. The pair is coordinate
// system agnostic: run reads it as a Worldographer offset for --x-y and as an
// axial coordinate for --q-r.
type coordPair struct {
	a, b int
	set  bool
}

func (c *coordPair) Set(s string) error {
	a, b, err := parsePair(s)
	if err != nil {
		return err
	}
	c.a, c.b, c.set = a, b, true
	return nil
}

func (c *coordPair) String() string {
	if c == nil || !c.set {
		return ""
	}
	return fmt.Sprintf("%d,%d", c.a, c.b)
}

// parsePair parses an "A,B" pair of integers, tolerating surrounding
// whitespace around each value.
func parsePair(s string) (int, int, error) {
	a, b, ok := strings.Cut(s, ",")
	if !ok {
		return 0, 0, fmt.Errorf("want two integers as A,B, got %q", s)
	}
	ai, err := strconv.Atoi(strings.TrimSpace(a))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid first value in %q", s)
	}
	bi, err := strconv.Atoi(strings.TrimSpace(b))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid second value in %q", s)
	}
	return ai, bi, nil
}

// usage writes the command banner followed by the ff-rendered flag list, so
// the flag help stays in sync with the definitions in run automatically.
func usage(w io.Writer, fs *ff.FlagSet) {
	_, _ = fmt.Fprintln(w, `woly — build the opyl map artifact from a Worldographer source

Usage:
  woly --source MAP.wxx --key KEY.json [--out ARTIFACT.json] [--x-y X,Y --q-r Q,R]`)
	_, _ = fmt.Fprintf(w, "\n%s\n", ffhelp.Flags(fs))
	_, _ = fmt.Fprintln(w, `woly emits a complete, deterministic artifact from the source; it never
reads or extends a prior one. See docs/content/reference/model/map-artifact.md
and docs/adr ADR 0004.`)
}

// floodWaterInside until we find a different terrain OR we hit a region boundary
func floodWaterInside(name string, provinces map[hex.Axial]*InputProvince, coords hex.Axial, inside bool) {
	// visit all the neighbors of the current hex
}

// floodLandInside until we find a different terrain OR we hit a region boundary
func floodLandInside(name string, provinces map[hex.Axial]*InputProvince, coords hex.Axial, inside bool) {
	// visit all the neighbors of the current hex
}
