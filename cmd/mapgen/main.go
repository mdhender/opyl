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
	"flag"
	"fmt"
	"os"

	"github.com/mdhender/opyl/pkg/mapgen"
)

func main() {
	in := flag.String("in", ".", "directory containing input files")
	out := flag.String("out", ".", "directory to write output files")
	flag.Parse()

	g := mapgen.New(mapgen.Options{
		InputDir:  *in,
		OutputDir: *out,
		Log:       os.Stderr,
	})

	if err := g.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "mapgen: %v\n", err)
		os.Exit(1)
	}
}
