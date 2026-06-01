// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package prng_test

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"

	"github.com/mdhender/opyl/internal/infra/prng"
)

// Basic seeded use. Construct a PRNG from a 128-bit seed and roll some dice.
func ExampleNewFromSeed() {
	p := prng.NewFromSeed(42, 43)
	for i := 0; i < 5; i++ {
		fmt.Println(p.Roll(1, 6))
	}
	// Output:
	// 5
	// 2
	// 5
	// 6
	// 5
}

// RollDice sums n dice of s sides.
func ExamplePRNG_RollDice() {
	p := prng.NewFromSeed(42, 43)
	for i := 0; i < 3; i++ {
		fmt.Println(p.RollDice(3, 6)) // 3d6
	}
	// Output:
	// 12
	// 17
	// 15
}

// PRNG embeds *rand.Rand, so anywhere a *rand.Rand is required, pass p.Rand.
// (The embedded field is also what powers helpers like Shuffle, Perm, etc.)
func ExamplePRNG_Rand() {
	p := prng.NewFromSeed(1, 2)

	// p.Rand is a *rand.Rand. Pass it to anything that wants one.
	var r *rand.Rand = p.Rand

	xs := []string{"a", "b", "c", "d", "e"}
	r.Shuffle(len(xs), func(i, j int) { xs[i], xs[j] = xs[j], xs[i] })
	fmt.Println(xs)
	// Output:
	// [b e c a d]
}

// Split derives independent substreams from a master generator. This is the
// recommended way to isolate per-stage or per-entity randomness so that
// consuming one stream does not perturb another.
func ExamplePRNG_Split() {
	master := prng.NewFromSeed(100, 200)
	a := master.Split()
	b := master.Split()
	for i := 0; i < 3; i++ {
		fmt.Println(a.Roll(1, 1000), b.Roll(1, 1000))
	}
	// Output:
	// 706 775
	// 917 39
	// 1 310
}

// MarshalBinary / UnmarshalBinary round-trips the full generator state.
// A zero-value PRNG can be unmarshaled into directly.
func ExamplePRNG_MarshalBinary() {
	p := prng.NewFromSeed(7, 11)
	for i := 0; i < 5; i++ {
		p.Uint64() // burn some state
	}

	data, _ := p.MarshalBinary()

	var q prng.PRNG
	_ = q.UnmarshalBinary(data)

	// p and q now produce identical sequences.
	for i := 0; i < 3; i++ {
		fmt.Println(p.Uint64() == q.Uint64())
	}
	// Output:
	// true
	// true
	// true
}

// PRNG also implements json.Marshaler / json.Unmarshaler (base64 inside a
// JSON string) and driver.Valuer / sql.Scanner (raw bytes), so the same
// state can be persisted to a JSON document or a SQL BLOB column.
func ExamplePRNG_MarshalJSON() {
	p := prng.NewFromSeed(13, 17)
	js, _ := json.Marshal(p)
	fmt.Println(string(js))
	// Output:
	// "cGNnOgAAAAAAAAANAAAAAAAAABE="
}
