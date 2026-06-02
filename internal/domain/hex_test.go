package domain

import (
	"errors"
	"testing"
)

func TestCoordString(t *testing.T) {
	tests := []struct {
		c    Coord
		want string
	}{
		{Coord{8, -5}, "[8,-5]"},
		{Coord{0, 0}, "[0,0]"},
		{Coord{-3, 12}, "[-3,12]"},
	}
	for _, tt := range tests {
		if got := tt.c.String(); got != tt.want {
			t.Errorf("Coord%v.String() = %q, want %q", tt.c, got, tt.want)
		}
	}
}

func TestParseCoord(t *testing.T) {
	ok := []struct {
		in   string
		want Coord
	}{
		{"[8,-5]", Coord{8, -5}},
		{"[0,0]", Coord{0, 0}},
		{"[-3,12]", Coord{-3, 12}},
		{"[ 8 , -5 ]", Coord{8, -5}}, // lenient: interior whitespace
		{"  [0,0]  ", Coord{0, 0}},   // lenient: outer whitespace
	}
	for _, tt := range ok {
		got, err := ParseCoord(tt.in)
		if err != nil {
			t.Errorf("ParseCoord(%q) returned error: %v", tt.in, err)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseCoord(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}

	bad := []string{"", "8,-5", "[8,-5", "8,-5]", "[8]", "[8,5,2]", "[a,b]", "[8 5,2]"}
	for _, in := range bad {
		if _, err := ParseCoord(in); !errors.Is(err, ErrInvalidCoord) {
			t.Errorf("ParseCoord(%q) error = %v, want ErrInvalidCoord", in, err)
		}
	}
}

func TestParseCoordRoundTrip(t *testing.T) {
	for _, c := range []Coord{{8, -5}, {0, 0}, {-3, 12}, {100, -100}} {
		got, err := ParseCoord(c.String())
		if err != nil {
			t.Fatalf("ParseCoord(%q): %v", c.String(), err)
		}
		if got != c {
			t.Errorf("round trip %v -> %q -> %v", c, c.String(), got)
		}
	}
}

func TestNeighborAndVectors(t *testing.T) {
	// Direction vectors per reference/model/map.md, applied from origin.
	origin := Coord{0, 0}
	want := map[Direction]Coord{
		North:     {0, -1},
		Northeast: {1, -1},
		Southeast: {1, 0},
		South:     {0, 1},
		Southwest: {-1, 1},
		Northwest: {-1, 0},
	}
	for dir, w := range want {
		if got := origin.Neighbor(dir); got != w {
			t.Errorf("origin.Neighbor(%s) = %v, want %v", dir, got, w)
		}
	}

	// Every neighbour is exactly one step away, and the opposite step
	// returns home.
	c := Coord{4, -2}
	for _, dir := range Directions {
		n := c.Neighbor(dir)
		if d := c.Distance(n); d != 1 {
			t.Errorf("Distance(%v, neighbour %s %v) = %d, want 1", c, dir, n, d)
		}
		if back := n.Sub(c).Add(c); back != n {
			t.Errorf("Sub/Add inconsistent for %s", dir)
		}
	}
}

func TestDistance(t *testing.T) {
	tests := []struct {
		a, b Coord
		want int
	}{
		{Coord{0, 0}, Coord{0, 0}, 0},
		{Coord{0, 0}, Coord{0, -1}, 1},  // one step N
		{Coord{0, 0}, Coord{3, 0}, 3},   // three steps SE
		{Coord{0, 0}, Coord{-2, -1}, 3}, // around the corner
		{Coord{1, -1}, Coord{1, -1}, 0},
	}
	for _, tt := range tests {
		if got := tt.a.Distance(tt.b); got != tt.want {
			t.Errorf("Distance(%v, %v) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
		if got := tt.b.Distance(tt.a); got != tt.want {
			t.Errorf("Distance(%v, %v) = %d, want %d (symmetry)", tt.b, tt.a, got, tt.want)
		}
	}
}

func TestParseDirection(t *testing.T) {
	for _, dir := range Directions {
		got, err := ParseDirection(dir.String())
		if err != nil {
			t.Fatalf("ParseDirection(%q): %v", dir.String(), err)
		}
		if got != dir {
			t.Errorf("ParseDirection(%q) = %s, want %s", dir.String(), got, dir)
		}
	}
	if got, err := ParseDirection(" se "); err != nil || got != Southeast {
		t.Errorf("ParseDirection(\" se \") = %s, %v; want SE, nil", got, err)
	}
	if _, err := ParseDirection("E"); !errors.Is(err, ErrInvalidDirection) {
		t.Errorf("ParseDirection(\"E\") error = %v, want ErrInvalidDirection", err)
	}
}

// TestCoordIsMapKey documents that Coord is comparable and usable as a
// map key directly. Two Coord values with the same Q and R hash and
// compare equal, so a second insert overwrites the first entry.
func TestCoordIsMapKey(t *testing.T) {
	m := map[Coord]string{}
	m[Coord{Q: 0, R: 0}] = "origin"
	m[Coord{Q: 1, R: -1}] = "NE"
	m[Coord{Q: 0, R: 0}] = "origin again" // same key: overwrites

	if len(m) != 2 {
		t.Fatalf("len(m) = %d, want 2", len(m))
	}
	if got := m[Coord{Q: 0, R: 0}]; got != "origin again" {
		t.Errorf("m[origin] = %q, want %q", got, "origin again")
	}
	if got := m[Coord{Q: 1, R: -1}]; got != "NE" {
		t.Errorf("m[NE] = %q, want %q", got, "NE")
	}
	if _, ok := m[Coord{Q: 99, R: 99}]; ok {
		t.Errorf("unexpected hit for absent key")
	}
}

func TestOffsetAxialRoundTrip(t *testing.T) {
	// The provenance example from reference/model/map-artifact.md: the
	// origin block pins Worldographer (5, 8) — but that is a re-centred
	// pin, not the raw conversion. Here we exercise the raw odd-q formula
	// and its inverse over a spread of columns, including odd ones.
	for _, o := range []Offset{{0, 0}, {5, 8}, {1, 0}, {2, 0}, {-3, 4}, {-4, 4}, {7, -2}} {
		c := o.ToAxial()
		if back := c.ToOffset(); back != o {
			t.Errorf("Offset%v -> Coord%v -> Offset%v, want round trip", o, c, back)
		}
	}

	// Spot-check the formula directly: odd column 1 shifts the axial r.
	if got := (Offset{Col: 1, Row: 0}).ToAxial(); got != (Coord{Q: 1, R: 0}) {
		t.Errorf("Offset{1,0}.ToAxial() = %v, want [1,0]", got)
	}
	if got := (Offset{Col: 2, Row: 0}).ToAxial(); got != (Coord{Q: 2, R: -1}) {
		t.Errorf("Offset{2,0}.ToAxial() = %v, want [2,-1]", got)
	}
}
