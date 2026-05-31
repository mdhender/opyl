---
title: Idempotency by design
weight: 2
prev: /explanation/sousa-in-opyl
---

opyl's pipeline runs offline, on an operator's schedule. Operators rerun
things — sometimes because something failed, sometimes because they want
to regenerate a report, sometimes by mistake. This page explains why
idempotency is an explicit Application concern in opyl and not a database
detail.

## The temptation

The obvious place to make `ProcessTurn` re-runnable is the database: add
a `UNIQUE` constraint, catch the constraint violation, treat it as a
no-op. This is tempting and wrong.

It is wrong because it bakes the idempotency policy into one specific
storage backend. opyl deliberately defers the storage choice (SQLite vs.
per-turn file snapshots) behind `app.GameStateStore`. If idempotency
lives in SQLite, switching to file snapshots silently loses it.

## Where it lives instead

`internal/app` declares `TurnLedger`:

```go
type TurnLedger interface {
    AlreadyProcessed(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber, inputHash string) (bool, error)
    Record(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber, inputHash string) error
}
```

`Services.ProcessTurn` computes a hash of the validated input
(`(prior state, validated orders)`), asks the ledger whether that exact
input has already been processed, and short-circuits if so. The ledger
implementation can be a SQLite table, a sidecar JSON file, anything — it
is an infra concern.

## Why hash the input?

Hashing the `(state, orders)` tuple, not just `(gameID, turn)`, lets the
operator intentionally rerun a turn after fixing bad orders. The hash
changes, so the ledger does not short-circuit. This matters because in a
batch game it is often the operator's job to fix garbled orders and
resolve the turn again.

## What this rules out

- "Lock the turn before processing." Locks are a concurrency tool, not an
  idempotency tool. opyl processes one turn per game serially; the
  ledger handles repeated invocations, not concurrent ones.
- "Make adapters detect duplicates." Adapters should be dumb. Putting
  policy in the adapter scatters it across every storage backend and
  every renderer.

## Trade-offs

The cost is one extra port and one extra step per use case. The benefit
is that idempotency survives every architectural decision still on the
table, including storage and renderer choice. That is a cheap insurance
premium for a project whose primary failure mode is "the operator ran it
twice."
