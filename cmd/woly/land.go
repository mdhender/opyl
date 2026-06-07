// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"

	"github.com/mdhender/opyl/internal/domain"
	"github.com/mdhender/opyl/internal/infra/prng"
	"github.com/mdhender/ottomap/hex"
)

type InputLand struct {
	ID    hex.Axial // coords uniquely identify a land
	Name  string
	Glyph string
}

// name groups of provinces
func setProvinceNames(provinces map[hex.Axial]*InputProvince, lands []*InputLand, rnd *prng.PRNG) {
	landCount, provinceCount := 0, 0
	for _, land := range lands {
		landCount++
		tile, ok := provinces[land.ID]
		if !ok {
			fmt.Printf("setProvinceNames: %q: (%d,%d) missing\n", land.Name, land.ID.Q, land.ID.R)
			continue
		}
		if tile.SaveChar != land.Glyph {
			fmt.Printf("setProvinceNames: %q: (%d,%d): want %q, got %q\n", land.Name, land.ID.Q, land.ID.R, land.Glyph, tile.SaveChar)
		}
		if tile.Name != "" {
			fmt.Printf("setProvinceNames: %q: (%d,%d): name collision %q\n", land.Name, land.ID.Q, land.ID.R, tile.Name)
			panic("setProvinceNames: name collision")
		}
		provinceCount += floodLandClumps(provinces, tile.ID, land.Name)
	}
	fmt.Printf("set_province_clumps: named %d areas, %d provinces\n", landCount, provinceCount)

}

// floodLandClumps until we hit water, a different glyph, or a different color
func floodLandClumps(provinces map[hex.Axial]*InputProvince, coords hex.Axial, name string) int {
	tile, ok := provinces[coords]
	if !ok {
		fmt.Printf("floodLandClumps: (%d,%d) missing\n", coords.Q, coords.R)
		return 0
	}
	tile.Name = name
	provincesNamed := 1

	// visit all the neighbors of the tile
	for _, neighborCoords := range coords.Neighbors() {
		if neighbor, ok := provinces[coords]; !ok {
			continue
		} else if neighbor.Terrain == domain.TerrOcean || neighbor.Color == -1 || neighbor.Color != tile.Color {
			continue
		} else if neighbor.Name == tile.Name {
			continue // already been here
		} else if neighbor.Name != "" {
			fmt.Printf("floodLandClumps: %q: (%d,%d): name collision %q\n", tile.Name, tile.ID.Q, tile.ID.R, neighbor.Name)
			panic("floodLandClumps: name collision")
		}
		// recursively flood fill this neighbor's neighbors
		provincesNamed += floodLandClumps(provinces, neighborCoords, name)
	}
	return provincesNamed
}
