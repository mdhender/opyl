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

	"github.com/maloquacious/wxx"
	"github.com/maloquacious/wxx/xmlio"
	"github.com/mdhender/opyl/internal/domain"
	"github.com/mdhender/opyl/internal/infra/prng"
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
		source = fs.StringLong("source", "", "path to the Worldographer .wxx source (required)")
		key    = fs.StringLong("key", "", "path to the opyl-key.json file (required)")
		out    = fs.StringLong("out", "", "path to write the map artifact JSON (default: stdout)")
		xy     coordPair
		qr     coordPair
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
		Regions:  map[string]*MapRegion{},
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
		}
		for k, v := range mapKeys.Terrains {
			v.Label = k
		}
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
	fp, err := os.Open(*source)
	if err != nil {
		fmt.Printf("\t%v\n", err)
		return err
	}
	defer fp.Close()

	fmt.Printf("info: wxx version %s\n", wxx.Version().Core())

	var decoderDiagnostics xmlio.DecoderDiagnostics
	decoder := xmlio.NewDecoder(xmlio.WithDecoderDiagnostics(&decoderDiagnostics))
	w, err := decoder.Decode(fp)
	if err != nil {
		fmt.Printf("\t%v\n", err)
		return err
	}
	fmt.Printf("\t%8s schema version %q\n", decoderDiagnostics.Schema, w.MetaData.DataVersion.String())
	fmt.Printf("\t%8d tiles high\n", w.Tiles.TilesHigh)
	fmt.Printf("\t%8d tiles wide\n", w.Tiles.TilesWide)
	fmt.Printf("\t%8d terrain tiles defined\n", len(w.TerrainMap.List))

	// verify the map the Worldographer terrain to opyl terrain
	unknownTiles := 0
	mapKeys.Indexes = make([]string, len(w.TerrainMap.List), len(w.TerrainMap.List))
	for _, terrain := range w.TerrainMap.List {
		key, ok := mapKeys.Terrains[terrain.Label]
		if !ok {
			fmt.Printf("\tterrain %-55q not defined in opyl\n", terrain.Label)
			unknownTiles++
			continue
		}
		key.Label, key.Index = terrain.Label, terrain.Index
		mapKeys.Labels = append(mapKeys.Labels, key.Label)
		mapKeys.Indexes[key.Index] = key.Label
	}
	if unknownTiles != 0 {
		return fmt.Errorf(".wxx file contains unknown terrain")
	}
	sort.Strings(mapKeys.Labels)

	// force a seed; later the GM will be able to specify it.
	rnd := prng.NewFromSeed(42, 43)

	mapTiles := map[AxialQR]*domain.Tile{}
	for _, cells := range w.Tiles.Tiles {
		if cells == nil {
			continue
		}
		for _, cell := range cells {
			if cell == nil {
				continue
			}

			label := mapKeys.Indexes[cell.Terrain]

			// convert Worldographer coordinates to axial
			xCell, yCell := cell.Column, cell.Row
			axial := (domain.Offset{Col: xCell, Row: yCell}).ToAxial().Add(delta)

			tileData := mapKeys.Terrains[label]
			tileData.Count++

			t := &domain.Tile{Q: axial.Q, R: axial.R, Glyph: tileData.Glyph}
			switch t.Glyph {
			case ";":
				t.Terrain = domain.TerrOcean
				t.Color = 1
				t.IsSeaLane = true
			case ",":
				t.Terrain = domain.TerrOcean
				t.Color = 1

			case ":":
				t.Terrain = domain.TerrOcean
				t.Color = 2
				t.IsSeaLane = true
			case ".":
				t.Terrain = domain.TerrOcean
				t.Color = 2

			case "~":
				t.Terrain = domain.TerrOcean
				t.Color = 3
				t.IsSeaLane = true
			case " ":
				t.Terrain = domain.TerrOcean
				t.Color = 3

			case "\"":
				t.Terrain = domain.TerrOcean
				t.Color = 4
				t.IsSeaLane = true
			case "'":
				t.Terrain = domain.TerrOcean
				t.Color = 4

			case "p":
				t.Terrain = domain.TerrPlain
				t.Color = 5
			case "P":
				t.Terrain = domain.TerrPlain
				t.Color = 6

			case "d":
				t.Terrain = domain.TerrDesert
				t.Color = 7
			case "D":
				t.Terrain = domain.TerrDesert
				t.Color = 8

			case "m":
				t.Terrain = domain.TerrMountain
				t.Color = 9
			case "M":
				t.Terrain = domain.TerrMountain
				t.Color = 10

			case "s":
				t.Terrain = domain.TerrSwamp
				t.Color = 11
			case "S":
				t.Terrain = domain.TerrSwamp
				t.Color = 12

			case "f":
				t.Terrain = domain.TerrForest
				t.Color = 13
			case "F":
				t.Terrain = domain.TerrForest
				t.Color = 14

			case "o":
				switch rnd.Roll(1, 10) {
				case 1, 2, 3:
					t.Terrain = domain.TerrForest
				case 4, 5, 6:
					t.Terrain = domain.TerrPlain
				case 7, 8:
					t.Terrain = domain.TerrMountain
				case 9:
					t.Terrain = domain.TerrSwamp
				case 10:
					t.Terrain = domain.TerrDesert
				}
				t.Color = -1

			case "?":
				t.IsHidden = true
				t.Terrain = domain.TerrLand

				// Special stuff

			case "^":
				t.Terrain = domain.TerrMountain
				t.Color = 9 // was 15, unique
				t.UldimFlag = domain.UldimFlag1
				t.IsRegionBoundary = true
			case "v":
				t.Terrain = domain.TerrMountain
				t.Color = 9 // was 15, unique
				t.UldimFlag = domain.UldimFlag2
				t.IsRegionBoundary = true
			case "{":
				t.Name = "Uldim pass"
				t.Terrain = domain.TerrMountain
				t.Color = 16
				t.UldimFlag = domain.UldimFlag3
				t.IsRegionBoundary = true
			case "}":
				t.Name = "Uldim pass"
				t.Terrain = domain.TerrMountain
				t.Color = 16
				t.UldimFlag = domain.UldimFlag4
				t.IsRegionBoundary = true

			case "]":
				t.Name = "Summerbridge"
				t.Terrain = domain.TerrSwamp
				t.SummerbridgeFlag = domain.SummerbridgeFlag1
				t.IsRegionBoundary = true
			case "[":
				t.Name = "Summerbridge"
				t.Terrain = domain.TerrSwamp
				t.SummerbridgeFlag = domain.SummerbridgeFlag2
				t.IsRegionBoundary = true

			case "O":
				t.Name = "Mt. Olympus"
				t.Terrain = domain.TerrMountain
				t.Color = -1

			case "1":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`todo: implement create a city here`)
				//fmt.Println(`n = create_a_city(row, col, "Drassa", true`)
				//fmt.Println(`subloc[n].IsSafeHaven = true`)
				//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "2":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`todo: implement create a city here`)
				//fmt.Println(`n = create_a_city(row, col, "Rimmon", true)`)
				//fmt.Println(`subloc[n].IsSafeHaven = true`)
				//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "3":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`todo: implement create a city here`)
				//fmt.Println(`n = create_a_city(row, col, "Harn", true)`)
				//fmt.Println(`subloc[n].IsSafeHaven = true`)
				//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "4":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`todo: implement create a city here`)
				//fmt.Println(`n = create_a_city(row, col, "Imperial City", true)`)
				//fmt.Println(`subloc[n].IsSafeHaven = true`)
				//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "5":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`todo: implement create a city here`)
				//fmt.Println(`n = create_a_city(row, col, "Port Aurnos", true)`)
				//fmt.Println(`subloc[n].IsSafeHaven = true`)
				//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "6":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`todo: implement create a city here`)
				//fmt.Println(`n = create_a_city(row, col, "Greyfell", true)`)
				//fmt.Println(`subloc[n].IsSafeHaven = true`)
				//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "7":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`todo: implement create a city here`)
				//fmt.Println(`n = create_a_city(row, col, "Yellowleaf", true)`)
				//fmt.Println(`subloc[n].IsSafeHaven = true`)
				//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "8":
				t.Terrain = domain.TerrForest
				t.Color = 19
				fmt.Println(`todo: implement create a city here`)
				//fmt.Println(`n = create_a_city(row, col, "Golden City", true)`)
				//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)

				// starting city with a random name
			case "9", "0":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`todo: implement create a city here`)
				//fmt.Println(`n = create_a_city(row, col, NULL, true)`)
				//fmt.Println(`subloc[n].IsSafeHaven = true`)
				//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)

			case "*":
				t.Terrain = domain.TerrLand
				fmt.Println(`todo: implement create a city here`)
				//fmt.Println(`create_a_city(row, col, NULL, true)`)

			case "%":
				t.Terrain = domain.TerrLand
				fmt.Println(`todo: implement create a city here`)
				//fmt.Println(`create_a_city(row, col, NULL, true)`)

			default:
				panic(fmt.Sprintf("unknown terrain %q", t.Glyph))
			}

			mapTiles[AxialQR{Q: t.Q, R: t.R}] = t
		}
	}

	for _, label := range mapKeys.Labels {
		tileData, _ := mapKeys.Terrains[label]
		fmt.Printf("\tterrain %6d %-55q maps to %q\n", tileData.Count, label, tileData.Kind)
	}

	// convert the map tile data to artifact provinces
	for _, mapTile := range mapTiles {
		province := Province{
			Q:       mapTile.Q,
			R:       mapTile.R,
			Terrain: mapTile.Terrain.String(),
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
