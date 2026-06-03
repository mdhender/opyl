// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package main

// This file defines the in-memory shape of map data imported from
// Worldographer and the associated configuration files.

// TODO: give these better names

type MapKey struct {
	Terrains map[string]*TerrainData `json:"terrains"`
	Regions  map[string]*MapRegion   `json:"regions"`

	Labels  []string `json:"labels"`  // sorted tile labels
	Indexes []string `json:"indexes"` // sorted tile index
}

type TerrainData struct {
	Glyph string `json:"glyph,omitempty"`
	Kind  string `json:"kind,omitempty"`
	Count int    `json:"count,omitempty"`
	Label string `json:"label,omitempty"` // .wxx tile label
	Index int    `json:"index,omitempty"` // .wxx tile index
}

type MapRegion struct {
	Name   string   `json:"name,omitempty"`
	Coords OffsetXY `json:"coords"`
}
