// Copyright (c) 2026 Michael D Henderson. All rights reserved.

// Package prng implements a PRNG stream that can save and restore state.
// It does this by embedding a PCG. That allows PRNG to use the
// math/rand/v2 *Rand convenience methods while still exposing the PCG
// state for marshaling.
package prng

import (
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand/v2"
)

// PRNG is a deterministic random stream backed by a PCG generator.
// It embeds *rand.Rand so the full math/rand/v2 convenience API is
// available, and exposes the underlying PCG state through the standard
// encoding interfaces (BinaryMarshaler/Unmarshaler, json.Marshaler/
// Unmarshaler, sql/driver.Valuer, sql.Scanner).
//
// The zero value is not useful; construct a PRNG with NewFromSeed, or
// allocate a zero value and immediately unmarshal state into it.
type PRNG struct {
	*rand.Rand
	pcg *rand.PCG
}

// NewFromSeed returns a PRNG seeded from the given 128-bit PCG seed.
// The returned *PRNG can be used directly via its embedded *rand.Rand
// (accessed as p.Rand) anywhere a *rand.Rand is required.
func NewFromSeed(s1, s2 uint64) *PRNG {
	pcg := rand.NewPCG(s1, s2)
	return &PRNG{
		Rand: rand.New(pcg),
		pcg:  pcg,
	}
}

// Split returns a new PRNG whose stream is independent of p's.
// Each call advances p's state, so successive splits yield distinct
// substreams.
func (p *PRNG) Split() *PRNG {
	return NewFromSeed(p.Uint64(), p.Uint64())
}

// ensurePCG initializes the underlying PCG and Rand if the PRNG was
// allocated as a zero value (for example, before UnmarshalBinary).
func (p *PRNG) ensurePCG() {
	if p.pcg == nil {
		p.pcg = rand.NewPCG(0, 0)
		p.Rand = rand.New(p.pcg)
	}
}

// MarshalBinary implements encoding.BinaryMarshaler by delegating to
// the underlying PCG.
func (p *PRNG) MarshalBinary() ([]byte, error) {
	p.ensurePCG()
	return p.pcg.MarshalBinary()
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler by delegating
// to the underlying PCG.
func (p *PRNG) UnmarshalBinary(data []byte) error {
	p.ensurePCG()
	return p.pcg.UnmarshalBinary(data)
}

// MarshalJSON encodes the PRNG state as a base64-encoded JSON string.
func (p *PRNG) MarshalJSON() ([]byte, error) {
	b, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return json.Marshal(base64.StdEncoding.EncodeToString(b))
}

// UnmarshalJSON decodes a base64-encoded JSON string back into PRNG state.
func (p *PRNG) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return fmt.Errorf("prng: decode base64: %w", err)
	}
	return p.UnmarshalBinary(b)
}

// Value implements driver.Valuer. The PRNG state is stored as a BLOB
// (the raw binary marshaling of the underlying PCG).
func (p *PRNG) Value() (driver.Value, error) {
	return p.MarshalBinary()
}

// Scan implements sql.Scanner. It accepts the raw binary form produced
// by Value (either as []byte or as string).
func (p *PRNG) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		// Copy because some drivers reuse the underlying buffer.
		buf := make([]byte, len(v))
		copy(buf, v)
		return p.UnmarshalBinary(buf)
	case string:
		return p.UnmarshalBinary([]byte(v))
	case nil:
		return fmt.Errorf("prng: cannot scan nil into *PRNG")
	default:
		return fmt.Errorf("prng: cannot scan %T into *PRNG", src)
	}
}
