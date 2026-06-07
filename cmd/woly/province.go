// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"

	"github.com/mdhender/opyl/internal/domain"
	"github.com/mdhender/opyl/internal/infra/prng"
	"github.com/mdhender/ottomap"
	"github.com/mdhender/ottomap/hex"
)

type InputProvince struct {
	ID       hex.Axial // coordinates uniquely identify a province
	Glyph    string    // original Worldographer map glyph
	Region   string
	Name     string
	Terrain  domain.TerrainCode
	Color    int    // map coloring for flood fills
	SaveChar string // todo: unknown use?

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

func DecodeTile(tile ottomap.Tile, coord hex.Axial, glyph string, rnd *prng.PRNG) *InputProvince {
	t := InputProvince{
		ID:    coord,
		Glyph: glyph,
	}
	switch t.Glyph {
	case ";":
		t.Terrain = domain.TerrOcean
		t.Color = 1
		t.IsSeaLane = true
	case ",":
		t.Terrain = domain.TerrOcean
		t.Color = 1

	case ":":
		t.Terrain = domain.TerrOcean
		t.Color = 2
		t.IsSeaLane = true
	case ".":
		t.Terrain = domain.TerrOcean
		t.Color = 2

	case "~":
		t.Terrain = domain.TerrOcean
		t.Color = 3
		t.IsSeaLane = true
	case " ":
		t.Terrain = domain.TerrOcean
		t.Color = 3

	case "\"":
		t.Terrain = domain.TerrOcean
		t.Color = 4
		t.IsSeaLane = true
	case "'":
		t.Terrain = domain.TerrOcean
		t.Color = 4

	case "p":
		t.Terrain = domain.TerrPlain
		t.Color = 5
	case "P":
		t.Terrain = domain.TerrPlain
		t.Color = 6

	case "d":
		t.Terrain = domain.TerrDesert
		t.Color = 7
	case "D":
		t.Terrain = domain.TerrDesert
		t.Color = 8

	case "m":
		t.Terrain = domain.TerrMountain
		t.Color = 9
	case "M":
		t.Terrain = domain.TerrMountain
		t.Color = 10

	case "s":
		t.Terrain = domain.TerrSwamp
		t.Color = 11
	case "S":
		t.Terrain = domain.TerrSwamp
		t.Color = 12

	case "f":
		t.Terrain = domain.TerrForest
		t.Color = 13
	case "F":
		t.Terrain = domain.TerrForest
		t.Color = 14

	case "o":
		switch rnd.Roll(1, 10) {
		case 1, 2, 3:
			t.Terrain = domain.TerrForest
		case 4, 5, 6:
			t.Terrain = domain.TerrPlain
		case 7, 8:
			t.Terrain = domain.TerrMountain
		case 9:
			t.Terrain = domain.TerrSwamp
		case 10:
			t.Terrain = domain.TerrDesert
		}
		t.Color = -1

	case "?":
		t.IsHidden = true
		t.Terrain = domain.TerrLand

		// Special stuff

	case "^":
		t.Terrain = domain.TerrMountain
		t.Color = 9 // was 15, unique
		t.UldimFlag = UldimFlag1
		t.IsRegionBoundary = true
	case "v":
		t.Terrain = domain.TerrMountain
		t.Color = 9 // was 15, unique
		t.UldimFlag = UldimFlag2
		t.IsRegionBoundary = true
	case "{":
		t.Name = "Uldim pass"
		t.Terrain = domain.TerrMountain
		t.Color = 16
		t.UldimFlag = UldimFlag3
		t.IsRegionBoundary = true
	case "}":
		t.Name = "Uldim pass"
		t.Terrain = domain.TerrMountain
		t.Color = 16
		t.UldimFlag = UldimFlag4
		t.IsRegionBoundary = true

	case "]":
		t.Name = "Summerbridge"
		t.Terrain = domain.TerrSwamp
		t.SummerbridgeFlag = SummerbridgeFlag1
		t.IsRegionBoundary = true
	case "[":
		t.Name = "Summerbridge"
		t.Terrain = domain.TerrSwamp
		t.SummerbridgeFlag = SummerbridgeFlag2
		t.IsRegionBoundary = true

	case "O":
		t.Name = "Mt. Olympus"
		t.Terrain = domain.TerrMountain
		t.Color = -1

	case "1":
		t.Terrain = domain.TerrForest
		t.Color = 19
		t.IsSafeHaven = true
		fmt.Println(`todo: implement create a city here`)
		//fmt.Println(`n = create_a_city(row, col, "Drassa", true`)
		//fmt.Println(`subloc[n].IsSafeHaven = true`)
		//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
	case "2":
		t.Terrain = domain.TerrForest
		t.Color = 19
		t.IsSafeHaven = true
		fmt.Println(`todo: implement create a city here`)
		//fmt.Println(`n = create_a_city(row, col, "Rimmon", true)`)
		//fmt.Println(`subloc[n].IsSafeHaven = true`)
		//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
	case "3":
		t.Terrain = domain.TerrForest
		t.Color = 19
		t.IsSafeHaven = true
		fmt.Println(`todo: implement create a city here`)
		//fmt.Println(`n = create_a_city(row, col, "Harn", true)`)
		//fmt.Println(`subloc[n].IsSafeHaven = true`)
		//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
	case "4":
		t.Terrain = domain.TerrForest
		t.Color = 19
		t.IsSafeHaven = true
		fmt.Println(`todo: implement create a city here`)
		//fmt.Println(`n = create_a_city(row, col, "Imperial City", true)`)
		//fmt.Println(`subloc[n].IsSafeHaven = true`)
		//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
	case "5":
		t.Terrain = domain.TerrForest
		t.Color = 19
		t.IsSafeHaven = true
		fmt.Println(`todo: implement create a city here`)
		//fmt.Println(`n = create_a_city(row, col, "Port Aurnos", true)`)
		//fmt.Println(`subloc[n].IsSafeHaven = true`)
		//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
	case "6":
		t.Terrain = domain.TerrForest
		t.Color = 19
		t.IsSafeHaven = true
		fmt.Println(`todo: implement create a city here`)
		//fmt.Println(`n = create_a_city(row, col, "Greyfell", true)`)
		//fmt.Println(`subloc[n].IsSafeHaven = true`)
		//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
	case "7":
		t.Terrain = domain.TerrForest
		t.Color = 19
		t.IsSafeHaven = true
		fmt.Println(`todo: implement create a city here`)
		//fmt.Println(`n = create_a_city(row, col, "Yellowleaf", true)`)
		//fmt.Println(`subloc[n].IsSafeHaven = true`)
		//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)
	case "8":
		t.Terrain = domain.TerrForest
		t.Color = 19
		fmt.Println(`todo: implement create a city here`)
		//fmt.Println(`n = create_a_city(row, col, "Golden City", true)`)
		//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)

		// starting city with a random name
	case "9", "0":
		t.Terrain = domain.TerrForest
		t.Color = 19
		t.IsSafeHaven = true
		fmt.Println(`todo: implement create a city here`)
		//fmt.Println(`n = create_a_city(row, col, NULL, true)`)
		//fmt.Println(`subloc[n].IsSafeHaven = true`)
		//fmt.Println(`fmt.Printf("Start city #%c %s at (%d,%d)\n", buf[col], subloc[n].name, row, col)`)

	case "*":
		t.Terrain = domain.TerrLand
		fmt.Println(`todo: implement create a city here`)
		//fmt.Println(`create_a_city(row, col, NULL, true)`)

	case "%":
		t.Terrain = domain.TerrLand
		fmt.Println(`todo: implement create a city here`)
		//fmt.Println(`create_a_city(row, col, NULL, true)`)

	default:
		panic(fmt.Sprintf("unknown terrain %q", t.Glyph))
	}
	return &t
}
