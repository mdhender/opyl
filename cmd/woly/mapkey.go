// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package main

import "github.com/mdhender/ottomap/hex"

// This file defines the in-memory shape of map data imported from
// Worldographer and the associated configuration files.

// TODO: give these better names

type MapKey struct {
	Terrains map[string]*TerrainData `json:"terrains"`
	Regions  map[string]*RegionData  `json:"regions"`

	RegionList  []*RegionData  // sorted region names
	TerrainList []*TerrainData // sorted terrain names
}

type TerrainData struct {
	Glyph string `json:"glyph,omitempty"`
	Kind  string `json:"kind,omitempty"`
	Count int    `json:"count,omitempty"`
	Name  string `json:"-"` // .wxx terrain name
	Index int    `json:"-"` // .wxx terrain index
}

type RegionData struct {
	ID     hex.Axial // coords are unique identifier for regions
	Name   string    `json:"name,omitempty"`
	Coords OffsetXY  `json:"coords"`
}
