// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package domain

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
