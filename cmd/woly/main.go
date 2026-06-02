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
// See docs/content/reference/model/map-artifact.md for the artifact shape
// and the offset→axial conversion woly performs.
package main

import (
	"encoding/json"
	"flag"
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
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "woly:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("woly", flag.ContinueOnError)
	fs.Usage = func() { usage(fs.Output()) }
	var (
		source = fs.String("source", "", "path to the Worldographer .wxx source (required)")
		key    = fs.String("key", "", "path to the opyl-key.json file (required)")
		out    = fs.String("out", "", "path to write the map artifact JSON (default: stdout)")
		xy     = fs.String("x-y", "", "Worldographer hex to pin, as X,Y (e.g. 5,8)")
		qr     = fs.String("q-r", "", "axial coordinate to pin the hex to, as Q,R (e.g. 0,0)")
	)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *source == "" {
		fs.Usage()
		return fmt.Errorf("missing required -source")
	}
	if *key == "" {
		fs.Usage()
		return fmt.Errorf("missing required -key")
	}

	// The pin re-centres the world in axial space: convert both the pinned
	// Worldographer hex and the target coordinate to axial, then translate
	// every hex by their difference. Raw offset subtraction is wrong because
	// r depends nonlinearly on the column (map-artifact.md).
	delta := domain.Coord{Q: 0, R: 0}
	if (*xy == "") != (*qr == "") {
		return fmt.Errorf("-x-y and -q-r must be given together")
	}
	if *xy != "" {
		x, y, err := parsePair(*xy)
		if err != nil {
			return fmt.Errorf("-x-y: %w", err)
		}
		q, r, err := parsePair(*qr)
		if err != nil {
			return fmt.Errorf("-q-r: %w", err)
		}
		pinHex := domain.Offset{Col: x, Row: y}
		pinTo := domain.Coord{Q: q, R: r}
		delta = pinTo.Sub(pinHex.ToAxial())
	}

	dest := "stdout"
	if *out != "" {
		dest = *out
	}
	_, _ = fmt.Fprintf(os.Stderr, "woly: source %s . %s; axial pin offset %s\n", *source, dest, delta)

	// load the key file
	type tile struct {
		Code  string
		Kind  string
		Count int
	}
	wxxKeys := struct {
		TerrainNames []string // wxx terrain names, sorted
		Tiles        map[string]*tile
	}{
		Tiles: map[string]*tile{},
	}
	if data, err := os.ReadFile(*key); err != nil {
		fmt.Printf("\t%v\n", err)
		return err
	} else if err = json.Unmarshal(data, &wxxKeys); err != nil {
		fmt.Printf("\t%v\n", err)
		return err
	}

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
	indexToLabel := map[int]string{}
	unknownTiles := 0
	for _, terrain := range w.TerrainMap.List {
		if _, ok := wxxKeys.Tiles[terrain.Label]; !ok {
			fmt.Printf("\tterrain %-55q not defined in opyl\n", terrain.Label)
			unknownTiles++
			continue
		}
		wxxKeys.TerrainNames = append(wxxKeys.TerrainNames, terrain.Label)
		indexToLabel[terrain.Index] = terrain.Label
	}
	if unknownTiles != 0 {
		return fmt.Errorf(".wxx file contains unknown terrain")
	}
	sort.Strings(wxxKeys.TerrainNames)

	rnd := prng.NewFromSeed(42, 43)

	// iterate through all the Worldographer terrain tiles
	for r, row := range w.Tiles.Tiles {
		if row == nil {
			continue
		}
		for c, col := range row {
			if col == nil {
				continue
			}
			index, label := col.Terrain, indexToLabel[col.Terrain]
			fmt.Printf("(%4d,%4d) => (%4d,%4d) %2d %q\n", r, c, col.Row, col.Column, index, label)

			keyTile := wxxKeys.Tiles[label]
			keyTile.Count++

			t := domain.Tile{Glyph: keyTile.Code}
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
				fmt.Println(`n = create_a_city(row, col, "Drassa", true`)
				fmt.Println(`subloc[n].IsSafeHaven = true`)
				fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "2":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`n = create_a_city(row, col, "Rimmon", true)`)
				fmt.Println(`subloc[n].IsSafeHaven = true`)
				fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "3":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`n = create_a_city(row, col, "Harn", true)`)
				fmt.Println(`subloc[n].IsSafeHaven = true`)
				fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "4":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`n = create_a_city(row, col, "Imperial City", true)`)
				fmt.Println(`subloc[n].IsSafeHaven = true`)
				fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "5":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`n = create_a_city(row, col, "Port Aurnos", true)`)
				fmt.Println(`subloc[n].IsSafeHaven = true`)
				fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "6":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`n = create_a_city(row, col, "Greyfell", true)`)
				fmt.Println(`subloc[n].IsSafeHaven = true`)
				fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "7":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`n = create_a_city(row, col, "Yellowleaf", true)`)
				fmt.Println(`subloc[n].IsSafeHaven = true`)
				fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
			case "8":
				t.Terrain = domain.TerrForest
				t.Color = 19
				fmt.Println(`n = create_a_city(row, col, "Golden City", true)`)
				fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)

				// starting city with a random name
			case "9", "0":
				t.Terrain = domain.TerrForest
				t.Color = 19
				t.IsSafeHaven = true
				fmt.Println(`n = create_a_city(row, col, NULL, true)`)
				fmt.Println(`subloc[n].IsSafeHaven = true`)
				fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)

			case "*":
				t.Terrain = domain.TerrLand
				fmt.Println(`create_a_city(row, col, NULL, true)`)
				break

			case "%":
				t.Terrain = domain.TerrLand
				fmt.Println(`create_a_city(row, col, NULL, true)`)
				break

			default:
				panic(fmt.Sprintf("unknown terrain %q", t.Glyph))
			}
		}
	}
	for _, terrainName := range wxxKeys.TerrainNames {
		tileKey, _ := wxxKeys.Tiles[terrainName]
		fmt.Printf("\tterrain %6d %-55q maps to %q\n", tileKey.Count, terrainName, tileKey.Kind)
	}

	return fmt.Errorf("Worldographer .wxx import is not yet implemented")
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

func usage(w io.Writer) {
	_, _ = fmt.Fprintln(w, `woly — build the opyl map artifact from a Worldographer source

Usage:
  woly -source MAP.wxx [-out ARTIFACT.json] [-x-y X,Y -q-r Q,R]

Flags:
  -source   path to the Worldographer .wxx source (required)
  -out      path to write the map artifact JSON (default: stdout)
  -x-y      Worldographer hex to pin, as X,Y (e.g. 5,8)
  -q-r      axial coordinate to pin the hex to, as Q,R (e.g. 0,0)

woly emits a complete, deterministic artifact from the source; it never
reads or extends a prior one. See docs/content/reference/model/map-artifact.md
and docs/adr ADR 0004.`)
}
