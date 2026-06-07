// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package domain

// types in the file are works in progress as we port the C types into
// more idiomatic Go types
//
// The breakdown looks something like
//   Game_t
//     Map_t
//       Province_t -- captures everything in the hex?
//         Coords_t    -- coordinates on the map (axial)
//         Location_t  -- location, or sub location
//     Region_t     -- a collection of provinces. Regions count towards game goals
//

// Location defines the tile's contents, which is almost everything except the terrain.
type Location struct {
	ID     int   // unique entity number
	Coords Coord // axial coordinates
	Level  int   // level of the map

	Name   string  // name of the location
	Region *Region // region the location is assigned to
	Roads  []*Road // roads leaving the location

	// values that are only used when importing from a map
	City             int
	Color            int // used for flood filling
	MajorCity        int
	Mark             int
	RegionBoundary   int
	SafeHaven        int
	SaveChar         rune
	SeaLane          int
	SummerbridgeFlag int
	UldimFlag        int
}

type Region struct {
	ID     int   // unique entity number
	Coords Coord // axial coordinates for the origin of the region
	Level  int   // level of the map

	Name string // name of the region
}

// Gate is a one-way connection from a location to a destination location entity.
type Gate struct {
	Name  string
	ToLoc *Location
	Key   string
}

// Road is a one-way connection (secret pass, channel, etc.) from a location to a destination location entity.
type Road struct {
	Name   string
	ToLoc  *Location
	Hidden int
}
