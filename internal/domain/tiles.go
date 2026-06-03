// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package domain

type Tile struct {
	Q, R    int    // axial coordiates
	Glyph   string // original Worldographer map glyph
	Region  string
	Name    string
	Terrain TerrainCode
	Color   int // map coloring for flood fills

	IsHidden         bool
	IsRegionBoundary bool
	IsSafeHaven      bool
	IsSeaLane        bool

	SummerbridgeFlag SummerbridgeFlag
	UldimFlag        UldimFlag
}

type SummerbridgeFlag int

const (
	SummerbridgeFlag1 SummerbridgeFlag = iota + 1
	SummerbridgeFlag2
)

type UldimFlag int

const (
	UldimFlag1 UldimFlag = iota + 1
	UldimFlag2
	UldimFlag3
	UldimFlag4
)
