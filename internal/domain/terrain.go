// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package domain

import "fmt"

type TerrainCode int

const (
	TerrBlank       TerrainCode = iota
	TerrLand                    = 1
	TerrOcean                   = 2
	TerrForest                  = 3
	TerrSwamp                   = 4
	TerrMountain                = 5
	TerrPlain                   = 6
	TerrDesert                  = 7
	TerrWater                   = 8
	TerrIsland                  = 9
	TerrStoneCir                = 10 // circle of stones
	TerrGrove                   = 11 // mallorn grove
	TerrBog                     = 12
	TerrCave                    = 13
	TerrCity                    = 14
	TerrGuild                   = 15
	TerrGrave                   = 16
	TerrRuins                   = 17
	TerrBattlefield             = 18
	TerrEnchFor                 = 19 // enchanted forest
	TerrRockyHill               = 20
	TerrTreeCir                 = 21
	TerrPits                    = 22
	TerrPasture                 = 23
	TerrOasis                   = 24
	TerrYewGrove                = 25
	TerrSandPit                 = 26
	TerrSacGrove                = 27 // sacred grove
	TerrPopField                = 28 // poppy field
	TerrTemple                  = 29
	TerrLair                    = 30 // dragon lair
)

func (tc TerrainCode) String() string {
	switch tc {
	case TerrBlank:
		return ""
	case TerrLand:
		return "land"
	case TerrOcean:
		return "ocean"
	case TerrForest:
		return "forest"
	case TerrSwamp:
		return "swamp"
	case TerrMountain:
		return "mountain"
	case TerrPlain:
		return "plain"
	case TerrDesert:
		return "desert"
	case TerrWater:
		return "water"
	case TerrIsland:
		return "island"
	case TerrStoneCir:
		return "stonecir"
	case TerrGrove:
		return "grove"
	case TerrBog:
		return "bog"
	case TerrCave:
		return "cave"
	case TerrCity:
		return "city"
	case TerrGuild:
		return "guild"
	case TerrGrave:
		return "grave"
	case TerrRuins:
		return "ruins"
	case TerrBattlefield:
		return "battlefield"
	case TerrEnchFor:
		return "enchfor"
	case TerrRockyHill:
		return "rockyhill"
	case TerrTreeCir:
		return "treecir"
	case TerrPits:
		return "pits"
	case TerrPasture:
		return "pasture"
	case TerrOasis:
		return "oasis"
	case TerrYewGrove:
		return "yewgrove"
	case TerrSandPit:
		return "sandpit"
	case TerrSacGrove:
		return "sacgrove"
	case TerrPopField:
		return "popfield"
	case TerrTemple:
		return "temple"
	case TerrLair:
		return "lair"
	}
	return fmt.Sprintf("terrain(%d)", tc)
}
