# `internal/infra/prng`

A deterministic random-number stream backed by [PCG][pcg] from `math/rand/v2`,
with state that can be saved and restored through every standard Go encoding
contract.

This is the infrastructure adapter that satisfies the `app.RNG` port.
Application use cases depend on the narrow `app.RNG` interface, never on this
concrete type. Runtime constructs a `*prng.PRNG` with a configured seed and
injects it into `app.Services.RNG`.

## What it is

`prng.PRNG` is a thin wrapper around `*rand.PCG`. It embeds `*rand.Rand` so
the full `math/rand/v2` convenience API (`IntN`, `Float64`, `Shuffle`,
`Perm`, …) is available directly. The package adds:

- **Dice helpers** — `Roll(low, high)`, `RollDice(n, sides)`, and shortcuts
  `D4` / `D6` / `D8` / `D10` / `D12` / `D20` / `D100`.
- **Independent substreams** — `Split()` returns a fresh PRNG seeded from the
  master, so per-stage or per-entity randomness can be isolated. See the
  determinism / independence / non-perturbation tests for the guarantees.
- **State persistence** — implements
  `encoding.BinaryMarshaler` / `BinaryUnmarshaler`,
  `json.Marshaler` / `json.Unmarshaler`,
  `driver.Valuer`, and `sql.Scanner`.
  A zero-value `PRNG` can be unmarshaled into directly.

## Passing it as `*rand.Rand`

`*PRNG` is **not** itself a `*rand.Rand` — that's a concrete pointer type and
Go does not let embedding satisfy it. Whenever an API requires a `*rand.Rand`,
pass the embedded field:

```go
p := prng.NewFromSeed(1, 2)

// ✓ pass p.Rand
shuffleThings(p.Rand, items)

// ✗ does not compile
shuffleThings(p, items)
```

The dice helpers and the `math/rand/v2` methods are still callable directly on
`p` because they come from the embedded `*rand.Rand`.

## Examples

Runnable examples live in
[`example_test.go`](example_test.go) and are rendered by `go doc` /
pkg.go.dev. View them with:

```sh
go doc -all github.com/mdhender/opyl/internal/infra/prng
```

Covered:

| Example | What it shows |
| --- | --- |
| `ExampleNewFromSeed` | Construct and roll |
| `ExamplePRNG_RollDice` | Sum n dice of s sides |
| `ExamplePRNG_Rand` | Passing as `*rand.Rand` via `p.Rand` |
| `ExamplePRNG_Split` | Independent substreams |
| `ExamplePRNG_MarshalBinary` | Binary round-trip of state |
| `ExamplePRNG_MarshalJSON` | JSON-encoded state |

[pcg]: https://www.pcg-random.org/
