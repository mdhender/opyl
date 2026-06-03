// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package main

import (
	"encoding/json"
	"sort"
)

// This file defines the on-disk shape of the opyl map artifact and a
// deterministic marshaller for it. The artifact is the single versioned JSON
// document cmd/woly emits and the engine's planned MapSource port loads as
// immutable input (docs/adr ADR 0004).
//
// The format is documented field-for-field in
// docs/content/reference/model/map-artifact.md; these types mirror it exactly.
// The artifact carries the reduced runtime model (six string terrains, string
// directions, integer-only fields) — not woly's richer internal import types.
//
// Determinism contract (ADR 0004): every collection is a sorted array, never a
// JSON object-as-map; every numeric field is an integer; the same source
// produces byte-identical output, so the artifact is git-diffable. MarshalArtifact
// enforces the sort order and the arrays-not-null rule.

// CurrentSchemaVersion is the document shape version woly writes.
const CurrentSchemaVersion = 1

// Artifact is the top-level map document.
type Artifact struct {
	SchemaVersion int        `json:"schemaVersion"`
	Origin        Origin     `json:"origin"`
	NextEntityID  int        `json:"nextEntityId"`
	Regions       []Region   `json:"regions"`   // sorted by id ascending
	Provinces     []Province `json:"provinces"` // sorted by q ascending, then r ascending
}

// Origin records the pin woly applied: the Worldographer hex wxy mapped to
// axial qr. It is provenance only; the engine does not need it.
type Origin struct {
	WXY OffsetXY `json:"wxy"`
	QR  AxialQR  `json:"qr"`
}

// OffsetXY is a Worldographer offset hex (column, row).
type OffsetXY struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// AxialQR is an axial coordinate.
type AxialQR struct {
	Q int `json:"q"`
	R int `json:"r"`
}

// Region is a named collection of provinces.
type Region struct {
	ID   string `json:"id"`   // stable slug; referenced by Province.Region
	Name string `json:"name"` // display name
	Kind string `json:"kind"` // normal · hades · faery · cloudlands
}

// Province is the coordinate-addressed unit of the map.
type Province struct {
	Q            int           `json:"q"` // axial coordinate; the province's identity
	R            int           `json:"r"`
	Terrain      string        `json:"terrain"`      // plains · forest · swamp · mountain · desert · ocean
	Region       string        `json:"region"`       // region id; exactly one
	CivSeed      *int          `json:"civSeed"`      // authored turn-zero civ level; null means derive from buildings
	Routes       []Route       `json:"routes"`       // outgoing edges, sorted in direction order N,NE,SE,S,SW,NW
	Sublocations []Sublocation `json:"sublocations"` // authored static sub-locations, sorted by id
}

// Route is an outgoing edge from a province (or sub-location). A direction with
// no Route entry is a hole — there is no route that way.
type Route struct {
	Dir        string  `json:"dir"` // N · NE · SE · S · SW · NW
	To         AxialQR `json:"to"`  // destination province
	Days       int     `json:"days"`
	Impassable bool    `json:"impassable"` // route exists and is shown but cannot be traversed
	Hidden     bool    `json:"hidden"`     // discoverable via EXPLORE; usable only by the finding faction
	WaterName  string  `json:"waterName"`  // sea name for ocean routes; "" when none
}

// Sublocation is a static authored location nested in a province or another
// sub-location: an inn inside a city inside a province.
type Sublocation struct {
	ID                int           `json:"id"`      // entity number minted by woly
	SrcUUID           string        `json:"srcUuid"` // Worldographer source UUID; engine ignores it
	Type              string        `json:"type"`    // city · town · inn · port-city · island · tower · temple · mine · castle
	Name              string        `json:"name"`
	EntryDays         int           `json:"entryDays"`         // days to enter from the surrounding location; 0 when free
	SafeHaven         bool          `json:"safeHaven"`         // feeds civ contribution and havens rules
	InitialSettlement bool          `json:"initialSettlement"` // a location where a new faction may begin
	Routes            []Route       `json:"routes"`            // extra routes (e.g. a port-city's sea access); the OUT edge to the parent is implicit
	Sublocations      []Sublocation `json:"sublocations"`      // nested sub-locations, sorted by id
}

// dirOrder ranks the six edge directions for canonical route sorting.
var dirOrder = map[string]int{
	"N": 0, "NE": 1, "SE": 2, "S": 3, "SW": 4, "NW": 5,
}

// dirRank returns a route direction's sort rank; unknown directions sort last
// (and stably among themselves) rather than panicking, so a malformed artifact
// still serializes for inspection.
func dirRank(dir string) int {
	if r, ok := dirOrder[dir]; ok {
		return r
	}
	return len(dirOrder)
}

// MarshalArtifact canonicalizes a in place — sorting every collection into the
// order the format requires and replacing nil slices with empty arrays — then
// returns its indented JSON encoding. The result is byte-identical for equal
// inputs and safe to commit to git.
func MarshalArtifact(a *Artifact) ([]byte, error) {
	canonicalize(a)
	return json.MarshalIndent(a, "", "  ")
}

// canonicalize imposes the artifact's deterministic ordering and the
// arrays-not-null rule on a.
func canonicalize(a *Artifact) {
	if a.Regions == nil {
		a.Regions = []Region{}
	}
	sort.Slice(a.Regions, func(i, j int) bool {
		return a.Regions[i].ID < a.Regions[j].ID
	})

	if a.Provinces == nil {
		a.Provinces = []Province{}
	}
	sort.Slice(a.Provinces, func(i, j int) bool {
		if a.Provinces[i].Q != a.Provinces[j].Q {
			return a.Provinces[i].Q < a.Provinces[j].Q
		}
		return a.Provinces[i].R < a.Provinces[j].R
	})
	for i := range a.Provinces {
		a.Provinces[i].Routes = canonicalRoutes(a.Provinces[i].Routes)
		a.Provinces[i].Sublocations = canonicalSublocations(a.Provinces[i].Sublocations)
	}
}

// canonicalRoutes returns routes sorted into direction order, never nil.
func canonicalRoutes(routes []Route) []Route {
	if routes == nil {
		return []Route{}
	}
	sort.SliceStable(routes, func(i, j int) bool {
		return dirRank(routes[i].Dir) < dirRank(routes[j].Dir)
	})
	return routes
}

// canonicalSublocations returns sub-locations sorted by id, never nil, and
// recurses into each one's routes and nested sub-locations.
func canonicalSublocations(subs []Sublocation) []Sublocation {
	if subs == nil {
		return []Sublocation{}
	}
	sort.Slice(subs, func(i, j int) bool {
		return subs[i].ID < subs[j].ID
	})
	for i := range subs {
		subs[i].Routes = canonicalRoutes(subs[i].Routes)
		subs[i].Sublocations = canonicalSublocations(subs[i].Sublocations)
	}
	return subs
}
