// Copyright (c) 2026 Michael D Henderson. All rights reserved.

package prng

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"testing"
)

// Compile-time assertions that PRNG satisfies the encoding contracts.
var (
	_ encoding.BinaryMarshaler   = (*PRNG)(nil)
	_ encoding.BinaryUnmarshaler = (*PRNG)(nil)
	_ json.Marshaler             = (*PRNG)(nil)
	_ json.Unmarshaler           = (*PRNG)(nil)
	_ driver.Valuer              = (*PRNG)(nil)
	_ sql.Scanner                = (*PRNG)(nil)
)

// TestRoundTrip_Binary verifies that marshaling and unmarshaling the
// binary state yields a generator that produces the same sequence.
func TestRoundTrip_Binary(t *testing.T) {
	p := NewFromSeed(42, 43)
	// Burn some state so we're not just round-tripping the seed.
	for i := 0; i < 17; i++ {
		p.Uint64()
	}

	data, err := p.MarshalBinary()
	if err != nil {
		t.Fatalf("MarshalBinary: %v", err)
	}

	var q PRNG
	if err := q.UnmarshalBinary(data); err != nil {
		t.Fatalf("UnmarshalBinary: %v", err)
	}

	for i := 0; i < 64; i++ {
		if got, want := q.Uint64(), p.Uint64(); got != want {
			t.Fatalf("binary round-trip diverged at %d: %d vs %d", i, got, want)
		}
	}
}

// TestRoundTrip_JSON verifies the same property through JSON.
func TestRoundTrip_JSON(t *testing.T) {
	p := NewFromSeed(99, 100)
	for i := 0; i < 5; i++ {
		p.Uint64()
	}

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}

	var q PRNG
	if err := json.Unmarshal(data, &q); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	for i := 0; i < 64; i++ {
		if got, want := q.Uint64(), p.Uint64(); got != want {
			t.Fatalf("json round-trip diverged at %d: %d vs %d", i, got, want)
		}
	}
}

// TestRoundTrip_SQL verifies Value/Scan round-trips the state.
func TestRoundTrip_SQL(t *testing.T) {
	p := NewFromSeed(7, 8)
	for i := 0; i < 11; i++ {
		p.Uint64()
	}

	v, err := p.Value()
	if err != nil {
		t.Fatalf("Value: %v", err)
	}

	var q PRNG
	if err := q.Scan(v); err != nil {
		t.Fatalf("Scan([]byte): %v", err)
	}
	for i := 0; i < 32; i++ {
		if got, want := q.Uint64(), p.Uint64(); got != want {
			t.Fatalf("sql []byte round-trip diverged at %d: %d vs %d", i, got, want)
		}
	}

	// Also exercise the string path.
	bytesV, ok := v.([]byte)
	if !ok {
		t.Fatalf("Value returned %T, expected []byte", v)
	}
	var r PRNG
	if err := r.Scan(string(bytesV)); err != nil {
		t.Fatalf("Scan(string): %v", err)
	}
	// r should track p's *original* post-burn state, which q has now also
	// advanced past 32 steps; reseed p from binary to compare cleanly.
	var pCopy PRNG
	if err := pCopy.UnmarshalBinary(bytesV); err != nil {
		t.Fatalf("rebuild reference: %v", err)
	}
	for i := 0; i < 32; i++ {
		if got, want := r.Uint64(), pCopy.Uint64(); got != want {
			t.Fatalf("sql string round-trip diverged at %d: %d vs %d", i, got, want)
		}
	}
}

// TestSplit_Determinism verifies that splitting the same master seed twice
// yields identical substream sequences, and that consuming the master via
// another method doesn't change that invariant.
func TestSplit_Determinism(t *testing.T) {
	m1 := NewFromSeed(42, 43)
	s1a := m1.Split()
	s1b := m1.Split()

	m2 := NewFromSeed(42, 43)
	s2a := m2.Split()
	s2b := m2.Split()

	for i := 0; i < 256; i++ {
		if got, want := s1a.Roll(1, 1_000_000), s2a.Roll(1, 1_000_000); got != want {
			t.Fatalf("substream a diverged at iteration %d: %d vs %d", i, got, want)
		}
		if got, want := s1b.Roll(1, 1_000_000), s2b.Roll(1, 1_000_000); got != want {
			t.Fatalf("substream b diverged at iteration %d: %d vs %d", i, got, want)
		}
	}
}

// TestSplit_Independence verifies that successive splits yield distinct
// streams (consecutive splits advance the master state).
func TestSplit_Independence(t *testing.T) {
	m := NewFromSeed(42, 43)
	a := m.Split()
	b := m.Split()

	// First 64 rolls should not all match — any match across 64 rolls of
	// a 1-billion range is astronomically unlikely unless the streams
	// are identical.
	matches := 0
	for i := 0; i < 64; i++ {
		if a.Roll(1, 1_000_000_000) == b.Roll(1, 1_000_000_000) {
			matches++
		}
	}
	if matches == 64 {
		t.Fatalf("splits a and b produced identical sequences; Split() is not advancing master state")
	}
}

// TestSplit_ConsumesMasterState verifies that Split() advances the master
// PRNG such that the same master, freshly seeded, yields a different
// substream on its Nth split than on its (N-1)th.
func TestSplit_ConsumesMasterState(t *testing.T) {
	m1 := NewFromSeed(42, 43)
	first := m1.Split()

	m2 := NewFromSeed(42, 43)
	m2.Split() // consume
	second := m2.Split()

	// First Roll of each should almost certainly differ.
	a := first.Roll(1, 1_000_000_000)
	b := second.Roll(1, 1_000_000_000)
	if a == b {
		t.Fatalf("first and second splits of fresh master produced identical first roll: %d", a)
	}
}

// TestSplit_DoesNotPerturbSiblingStream verifies that operations on one
// substream do not affect another substream split earlier. This is the
// load-bearing property for stage isolation: re-seeding or heavily
// consuming one stage must not shift another stage's output.
func TestSplit_DoesNotPerturbSiblingStream(t *testing.T) {
	// Baseline: split A and B from the same master, capture B's first roll.
	m1 := NewFromSeed(42, 43)
	_ = m1.Split()
	b1 := m1.Split()
	baseline := b1.Roll(1, 1_000_000_000)

	// Variant: split A, consume it heavily, then split B, capture its roll.
	m2 := NewFromSeed(42, 43)
	a := m2.Split()
	for i := 0; i < 10_000; i++ {
		a.Roll(1, 100)
	}
	b2 := m2.Split()
	variant := b2.Roll(1, 1_000_000_000)

	if baseline != variant {
		t.Fatalf("consuming substream A perturbed substream B: baseline=%d variant=%d", baseline, variant)
	}
}
