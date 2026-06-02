// This file holds opyl's hexagonal map coordinate types: the axial
// (q, r) coordinate that is a province's identity, the six flat-top
// travel directions, and the pure transforms over them — neighbour,
// distance, display-code parsing, and Worldographer offset conversion.
//
// Every fact here derives from the engine's sole authority, the reference
// docs: docs/content/reference/model/map.md (identity, the implied cube
// axis, the direction-vector table, the display grammar) and
// map-artifact.md (the Worldographer offset→axial conversion).
//
// Domain stays self-contained: this file imports only the standard
// library, so the SOUSA import-conformance check (no internal deps under
// internal/domain) continues to pass. The shared cerr sentinels are a
// sibling inner package domain must not import; parse failures here are
// reported with domain-local sentinels, and the orderfile boundary wraps
// them into cerr.ErrInvalidOrders for untrusted input.

package domain

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mdhender/opyl/internal/cerr"
)

// Coord is an axial hex coordinate (Red Blob Games flat-top axial
// convention) and a province's identity. The implied third cube axis is
// S = −Q − R and is never stored. The origin (0, 0) sits near the centre
// of the authored world; coordinates run negative and positive in every
// direction, so the world extends outward without renumbering.
type Coord struct {
	Q, R int
}

// S returns the implied third cube axis, S = −Q − R, so that Q+R+S == 0.
func (c Coord) S() int { return -c.Q - c.R }

// Add returns the coordinate reached by translating c by d.
func (c Coord) Add(d Coord) Coord {
	return Coord{Q: c.Q + d.Q, R: c.R + d.R}
}

// Sub returns the translation that carries o to c (c − o).
func (c Coord) Sub(o Coord) Coord {
	return Coord{Q: c.Q - o.Q, R: c.R - o.R}
}

// Neighbor returns the coordinate one step from c across the edge in
// direction dir.
func (c Coord) Neighbor(dir Direction) Coord {
	return c.Add(dir.Vector())
}

// Distance returns the number of single-edge steps between c and o — the
// hex (cube) distance, computed on the implied cube coordinates.
func (c Coord) Distance(o Coord) int {
	return (abs(c.Q-o.Q) + abs(c.R-o.R) + abs(c.S()-o.S())) / 2
}

// String returns the canonical display code for c: square brackets, a
// comma, no spaces, no leading '+', and −0 normalised to 0 — for example
// "[8,-5]", "[0,0]". This is the strict form the engine always emits.
func (c Coord) String() string {
	return fmt.Sprintf("[%d,%d]", c.Q, c.R)
}

// ErrInvalidCoord is returned by ParseCoord when its input is not a
// well-formed display code. Domain keeps its own sentinel because the
// shared cerr package is an inner sibling that domain must not import; the
// orderfile boundary wraps this into cerr.ErrInvalidOrders.
const ErrInvalidCoord = cerr.Error("invalid coordinate")

// ParseCoord parses a bracketed display code such as "[8,-5]" into a
// Coord. It accepts the lenient form — interior whitespace is tolerated,
// so "[ 8 , -5 ]" reads the same as "[8,-5]" — and normalises to the
// canonical Coord. Anything that is not a well-formed "[q,r]" is rejected
// with ErrInvalidCoord; malformed input is never guessed at.
func ParseCoord(s string) (Coord, error) {
	fail := func() (Coord, error) {
		return Coord{}, fmt.Errorf("%w: %q", ErrInvalidCoord, s)
	}
	body, ok := strings.CutPrefix(strings.TrimSpace(s), "[")
	if !ok {
		return fail()
	}
	body, ok = strings.CutSuffix(body, "]")
	if !ok {
		return fail()
	}
	qs, rs, ok := strings.Cut(body, ",")
	if !ok {
		return fail()
	}
	q, err := strconv.Atoi(strings.TrimSpace(qs))
	if err != nil {
		return fail()
	}
	r, err := strconv.Atoi(strings.TrimSpace(rs))
	if err != nil {
		return fail()
	}
	return Coord{Q: q, R: r}, nil
}

// Direction is one of the six flat-top hex travel directions. There is no
// due east or west: those point at a vertex, not an edge. The constants
// are declared in canonical order — N, NE, SE, S, SW, NW — which is the
// order routes are sorted and emitted in.
type Direction int

const (
	North Direction = iota
	Northeast
	Southeast
	South
	Southwest
	Northwest
)

// Directions lists the six directions in canonical order: N, NE, SE, S,
// SW, NW. Range over it to iterate a hex's edges in emit order.
var Directions = [6]Direction{North, Northeast, Southeast, South, Southwest, Northwest}

// directionVectors maps each Direction to its axial Δ(Q, R). The vectors
// are screen-true for a flat-top hex — North is straight up on the GM's
// map — and are taken verbatim from reference/model/map.md.
var directionVectors = [...]Coord{
	North:     {Q: 0, R: -1},
	Northeast: {Q: +1, R: -1},
	Southeast: {Q: +1, R: 0},
	South:     {Q: 0, R: +1},
	Southwest: {Q: -1, R: +1},
	Northwest: {Q: -1, R: 0},
}

var directionAbbrev = [...]string{
	North:     "N",
	Northeast: "NE",
	Southeast: "SE",
	South:     "S",
	Southwest: "SW",
	Northwest: "NW",
}

// Vector returns the axial Δ(Q, R) added to a coordinate to step one hex
// in direction d.
func (d Direction) Vector() Coord {
	return directionVectors[d]
}

// String returns the canonical abbreviation for d: N, NE, SE, S, SW, NW.
func (d Direction) String() string {
	if d < North || d > Northwest {
		return fmt.Sprintf("Direction(%d)", int(d))
	}
	return directionAbbrev[d]
}

// ErrInvalidDirection is returned by ParseDirection for input that is not
// one of the six canonical abbreviations.
const ErrInvalidDirection = cerr.Error("invalid direction")

// ParseDirection parses a direction abbreviation (case-insensitive, outer
// whitespace tolerated: N, NE, SE, S, SW, NW) into a Direction. Unknown
// input is rejected with ErrInvalidDirection.
func ParseDirection(s string) (Direction, error) {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "N":
		return North, nil
	case "NE":
		return Northeast, nil
	case "SE":
		return Southeast, nil
	case "S":
		return South, nil
	case "SW":
		return Southwest, nil
	case "NW":
		return Northwest, nil
	default:
		return 0, fmt.Errorf("%w: %q", ErrInvalidDirection, s)
	}
}

// Offset is a Worldographer odd-q (vertical, flat-top) offset coordinate:
// a (Col, Row) pair in which odd columns are shifted down. It is map-source
// provenance only — the engine reasons in axial Coord, never in Offset —
// and exists so cmd/woly's coordinate math has a single home.
type Offset struct {
	Col, Row int
}

// ToAxial converts a Worldographer odd-q offset coordinate to its axial
// Coord: q = col, r = row − (col − (col & 1)) / 2
// (reference/model/map-artifact.md).
func (o Offset) ToAxial() Coord {
	return Coord{
		Q: o.Col,
		R: o.Row - (o.Col-(o.Col&1))/2,
	}
}

// ToOffset converts an axial Coord back to its Worldographer odd-q offset
// coordinate — the inverse of Offset.ToAxial: col = q,
// row = r + (q − (q & 1)) / 2.
func (c Coord) ToOffset() Offset {
	return Offset{
		Col: c.Q,
		Row: c.R + (c.Q-(c.Q&1))/2,
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
