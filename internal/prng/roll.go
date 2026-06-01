// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package prng

// Roll returns a uniformly distributed random integer in the
// range [low, high] (inclusive on both ends).
//
// Intentionally ignores over and under flow.
func (p *PRNG) Roll(low, high int) int {
	if high < low {
		low, high = high, low
	}
	delta := high - low + 1
	return low + p.IntN(delta)
}

// RollDice returns the total from rolling `n` dice, each with `d` sides.
//
// Example: RollDice(3, 4) simulates rolling 3 d4, returning a number from 3 to 12.
//
// Intentionally ignores over and under flow. Returns 0 if n or d are invalid.
func (p *PRNG) RollDice(n, d int) int {
	if d < 1 {
		return 0
	}
	result := 0
	for ; n > 0; n-- {
		result += 1 + p.IntN(d)
	}
	return result
}

// D4 returns the total from rolling `n` d4.
func (p *PRNG) D4(n int) int {
	return p.RollDice(n, 4)
}

// D6 returns the total from rolling `n` d6.
func (p *PRNG) D6(n int) int {
	return p.RollDice(n, 6)
}

// D8 returns the total from rolling `n` d8.
func (p *PRNG) D8(n int) int {
	return p.RollDice(n, 8)
}

// D10 returns the total from rolling `n` d10.
func (p *PRNG) D10(n int) int {
	return p.RollDice(n, 10)
}

// D12 returns the total from rolling `n` d12.
func (p *PRNG) D12(n int) int {
	return p.RollDice(n, 12)
}

// D20 returns the total from rolling `n` d20.
func (p *PRNG) D20(n int) int {
	return p.RollDice(n, 20)
}

// D100 returns the total from rolling `n` d100.
func (p *PRNG) D100(n int) int {
	return p.RollDice(n, 100)
}
