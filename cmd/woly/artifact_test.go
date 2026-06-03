// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package main

import (
	"strings"
	"testing"
)

// TestMarshalArtifactCanonicalOrder verifies that MarshalArtifact imposes the
// artifact's deterministic ordering regardless of the order fields were built:
// regions by id, provinces by (q, r), routes by direction, sub-locations by id.
func TestMarshalArtifactCanonicalOrder(t *testing.T) {
	civ := 3
	a := &Artifact{
		SchemaVersion: CurrentSchemaVersion,
		Origin:        Origin{WXY: OffsetXY{X: 5, Y: 8}, QR: AxialQR{Q: 0, R: 0}},
		NextEntityID:  6001,
		Regions: []Region{
			{ID: "tollus", Name: "Tollus", Kind: "normal"},
			{ID: "andelay", Name: "Andelay", Kind: "normal"},
		},
		Provinces: []Province{
			{
				Q: 8, R: -5, Terrain: "plains", Region: "tollus", CivSeed: &civ,
				Routes: []Route{
					{Dir: "SW", To: AxialQR{Q: 7, R: -4}},
					{Dir: "N", To: AxialQR{Q: 8, R: -6}},
					{Dir: "SE", To: AxialQR{Q: 9, R: -5}},
				},
				Sublocations: []Sublocation{
					{ID: 3102, Type: "inn", Name: "Hooting Owl Inn"},
					{ID: 2845, Type: "city", Name: "Carim"},
				},
			},
			{Q: 8, R: -7, Terrain: "ocean", Region: "andelay"},
			{Q: 3, R: 1, Terrain: "forest", Region: "andelay"},
		},
	}

	out, err := MarshalArtifact(a)
	if err != nil {
		t.Fatalf("MarshalArtifact: %v", err)
	}

	if got, want := a.Regions[0].ID, "andelay"; got != want {
		t.Errorf("regions[0].id = %q, want %q", got, want)
	}
	// provinces sort by q asc, then r asc.
	wantQR := [][2]int{{3, 1}, {8, -7}, {8, -5}}
	for i, qr := range wantQR {
		if a.Provinces[i].Q != qr[0] || a.Provinces[i].R != qr[1] {
			t.Errorf("provinces[%d] = (%d,%d), want (%d,%d)", i, a.Provinces[i].Q, a.Provinces[i].R, qr[0], qr[1])
		}
	}
	// the (8,-5) province is last; its routes sort N, SE, SW.
	p := a.Provinces[2]
	wantDirs := []string{"N", "SE", "SW"}
	for i, d := range wantDirs {
		if p.Routes[i].Dir != d {
			t.Errorf("routes[%d].dir = %q, want %q", i, p.Routes[i].Dir, d)
		}
	}
	// sub-locations sort by id.
	if p.Sublocations[0].ID != 2845 || p.Sublocations[1].ID != 3102 {
		t.Errorf("sublocations ids = %d,%d, want 2845,3102", p.Sublocations[0].ID, p.Sublocations[1].ID)
	}

	s := string(out)
	// civSeed present as an integer on the authored province, null elsewhere.
	if !strings.Contains(s, `"civSeed": 3`) {
		t.Errorf("expected civSeed: 3 in output:\n%s", s)
	}
	if !strings.Contains(s, `"civSeed": null`) {
		t.Errorf("expected a civSeed: null in output:\n%s", s)
	}
	// empty collections render as arrays, never null.
	if strings.Contains(s, "null,") && strings.Contains(s, `"routes": null`) {
		t.Errorf("collections must serialize as [], not null:\n%s", s)
	}
	if !strings.Contains(s, `"routes": []`) || !strings.Contains(s, `"sublocations": []`) {
		t.Errorf("expected empty routes/sublocations as []:\n%s", s)
	}
}

// TestMarshalArtifactDeterministic verifies byte-identical output across runs
// and across differently-ordered but equal inputs.
func TestMarshalArtifactDeterministic(t *testing.T) {
	build := func(regionsReversed bool) *Artifact {
		regions := []Region{
			{ID: "a", Name: "A", Kind: "normal"},
			{ID: "b", Name: "B", Kind: "normal"},
		}
		if regionsReversed {
			regions[0], regions[1] = regions[1], regions[0]
		}
		return &Artifact{
			SchemaVersion: CurrentSchemaVersion,
			NextEntityID:  100,
			Regions:       regions,
			Provinces: []Province{
				{Q: 0, R: 0, Terrain: "plains", Region: "a"},
			},
		}
	}

	first, err := MarshalArtifact(build(false))
	if err != nil {
		t.Fatalf("MarshalArtifact: %v", err)
	}
	second, err := MarshalArtifact(build(true))
	if err != nil {
		t.Fatalf("MarshalArtifact: %v", err)
	}
	if string(first) != string(second) {
		t.Errorf("output not deterministic across input orderings:\n--- first ---\n%s\n--- second ---\n%s", first, second)
	}
}
