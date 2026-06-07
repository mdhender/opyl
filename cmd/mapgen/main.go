// Command mapgen is a self-contained Go port of the Olympia G3 map generator.
//
// Like the original C program (g3-mapgen), it reads its input files from the
// current working directory:
//
//	Map      ascii map of the world
//	Regions  region (continent/ocean) names
//	Land     named land-area clumps
//	Cities   pool of city names
//	randseed 16-byte RNG seed/state
//
// and writes these output files to the current working directory:
//
//	loc      location database (provinces, sublocations, regions)
//	gate     gate database
//	road     road database
//	randseed updated RNG state
//
// Diagnostic output is written to stderr.
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mdhender/opyl/pkg/mapgen"
	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffhelp"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "mapgen:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := ff.NewFlagSet("mapgen")
	var (
		in          = fs.StringLong("in", ".", "directory containing input files")
		inRandSeed  = fs.StringLong("random-seed", "randseed", "file to read random seed from")
		out         = fs.StringLong("out", ".", "directory to write output files")
		outRandSeed = fs.StringLong("random-seed-output", "randseed", "file to write random seed to")
		xy          coordPair
		qr          coordPair
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

	// The pin re-centres the world in axial space: convert both the pinned
	// Worldographer hex and the target coordinate to axial, then translate
	// every hex by their difference. Raw offset subtraction is wrong because
	// r depends nonlinearly on the column (map-artifact.md).
	if xy.set != qr.set {
		return fmt.Errorf("--x-y and --q-r must be given together")
	}
	pinXY := mapgen.OffsetXY{X: xy.a, Y: xy.b}
	pinQR := mapgen.AxialQR{Q: qr.a, R: qr.b}

	g := mapgen.New(mapgen.Options{
		InputDir:       *in,
		InputMap:       "olympia-g4-source-v1.wxx",
		InputRandSeed:  *inRandSeed,
		OutputDir:      *out,
		OutputMap:      "olympia-g4-map-v1.wxx",
		OutputRandSeed: *outRandSeed,
		Log:            os.Stderr,
		Origin: mapgen.Origin{
			XY:    pinXY,
			QR:    pinQR,
			Delta: pinQR.Sub(pinXY.ToAxial()),
		},
	})

	if err := g.Run(); err != nil {
		return err
	}
	return g.RunHex()
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
	_, _ = fmt.Fprintln(w, `mapgen — build the opyl map artifact from a Worldographer source

Usage:
  mapgen --source MAP.wxx --key KEY.json [--out ARTIFACT.json] [--x-y X,Y --q-r Q,R]`)
	_, _ = fmt.Fprintf(w, "\n%s\n", ffhelp.Flags(fs))
	_, _ = fmt.Fprintln(w, `mapgen emits a complete, deterministic artifact from the source; it never
reads or extends a prior one. See docs/content/reference/model/map-artifact.md
and docs/adr ADR 0004.`)
}
