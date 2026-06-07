// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package mapgen

import "fmt"

// OffsetXY is a Worldographer offset hex (column, row).
// We assume that the map has been created in COLUMN orientation whichs gives
// an odd-q vertical layout when converting to Cube/Axial coordinates
type OffsetXY struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// String implements the stringer interface
func (c OffsetXY) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

// ToAxial converts a Worldographer odd-q offset coordinate to its axial
// Coord: q = col, r = row − (col − (col & 1)) / 2
// (reference/model/map-artifact.md).
func (c OffsetXY) ToAxial() AxialQR {
	return OffsetToAxial(c.X, c.Y)
}

// OffsetToAxial converts an odd-q offset hex (column, row) to axial (q, r).
// Subtracting (x & 1) makes the term even before halving, so the division is
// exact and behaves correctly for negative columns.
func OffsetToAxial(x, y int) AxialQR {
	return AxialQR{
		Q: x,
		R: y - (x-(x&1))/2,
	}
}

// AxialQR is an axial coordinate.
type AxialQR struct {
	Q int `json:"q"`
	R int `json:"r"`
}

// String implements the stringer interface
func (c AxialQR) String() string {
	return fmt.Sprintf("%d,%d", c.Q, c.R)
}

// Sub returns the translation that carries o to c (c − o).
func (c AxialQR) Sub(o AxialQR) AxialQR {
	return AxialQR{Q: c.Q - o.Q, R: c.R - o.R}
}

// AxialToOffset converts an axial hex (q, r) back to an odd-q offset hex
// (column, row). It is the inverse of OffsetToAxial.
func AxialToOffset(q, r int) OffsetXY {
	return OffsetXY{
		X: q,
		Y: r + (q-(q&1))/2,
	}
}
