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
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mdhender/opyl/internal/domain"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "woly:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("woly", flag.ContinueOnError)
	fs.Usage = func() { usage(fs.Output()) }
	var (
		source = fs.String("source", "", "path to the Worldographer .wxx source (required)")
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
	fmt.Fprintf(os.Stderr, "woly: source %s -> %s; axial pin offset %s\n", *source, dest, delta)

	// TODO: parse the .wxx source behind an internal/infra adapter, apply
	// `delta` to every province's axial coordinate, mint entity numbers for
	// sub-locations, and emit the deterministic artifact to dest.
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
	fmt.Fprintln(w, `woly — build the opyl map artifact from a Worldographer source

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
