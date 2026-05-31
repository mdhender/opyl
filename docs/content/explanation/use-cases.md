---
title: How use cases work
weight: 3
prev: /explanation/idempotency
---

A "use case" in opyl is a method on `app.Services` that owns exactly one
stage of the turn pipeline. There is no separate use-case interface, no
command/handler split, no orchestrator pattern. This page explains why
the shape ended up so plain and what that shape buys us.

## One method, one stage

opyl's pipeline has four stages — ingest, process, render, dispatch —
plus a top-level `RunPipeline` that chains them. Each stage is a method
on `Services` with the same signature: `(ctx, gameID, turn) → error`.
Nothing else. No per-stage struct, no `Execute()` interface, no request
DTO.

The temptation is to wrap each stage in its own type for "purity." We
don't, because the four stages are not independent products — they share
the same ports, the same identity tuple, and the same operator. Giving
each one its own struct would mean five constructors that all take the
same six adapters, just to look more object-oriented. `Services` is the
honest representation: one bag of ports, one method per thing the
operator can ask for.

The rule that does matter is **one method per stage, never a mega
method.** `RunPipeline` exists, but it is pure composition — it calls
the four stage methods in order and does nothing else. Operators must
be able to re-enter the pipeline at any stage after fixing a problem,
and that only works if each stage is independently invocable.

## Ports, not adapters

Use case bodies talk to the outside world exclusively through the
interfaces declared in `app/ports.go`. The `Services` struct holds those
interfaces; runtime injects concrete adapters at startup. The use case
itself never sees a `*sql.DB`, a `*gofpdf.Fpdf`, or an `smtp.Client`.

This is the SOUSA rule made concrete. Its practical consequence is that
every architectural decision still open in opyl — storage backend, PDF
library, mail transport, order format — can land without touching a
single use case. If implementing one of those choices forces a change
in `services.go`, the port is wrong; fix the port, not the use case.

## Domain types at the seams

Inputs and outputs at the use-case boundary are domain types:
`GameID`, `TurnNumber`, `OrderBundle`, `PlayerReport`, `Recipient`,
`Attachment`. Raw bytes stop at infra. Sentinel errors from `cerr`
travel back the other way to express business meaning
(`ErrInvalidOrders`, `ErrTurnAlreadyProcessed`, …).

This is what makes the parser-is-the-boundary rule enforceable. By the
time `IngestOrders` sees anything, it is already a validated
`OrderBundle`. The use case never has to ask "is this input
well-formed?" — that question was answered at the `infra/orderfile/`
edge.

## Idempotency lives here

Every state-mutating use case (`ProcessTurn`, `DispatchReports`) is
expected to be safe to invoke twice for the same `(gameID, turn)`. It
achieves that by consulting `TurnLedger` with a hash of its validated
input and short-circuiting on a match. The adapter underneath the
ledger stays dumb; the policy is in the use case.

The full rationale lives in [Idempotency by design](idempotency.md).
The point here is that idempotency is not a property of any adapter —
it is a property of the use case, which is why the use case is the
right place to enforce it.

## Render and dispatch stay separate

`RenderReports` produces bytes. `DispatchReports` sends bytes. The two
use cases call two different ports (`ReportRenderer` and
`ReportDispatcher`), and the runtime wires them together. We
deliberately resist combining them into one "render-and-send" use case,
even though the operator almost always runs them in sequence.

The reason is that splitting them makes dry-runs and re-sends trivial.
An operator who wants to inspect this turn's PDFs before mailing them
just runs `render` alone. An operator whose SMTP relay failed
yesterday just reruns `dispatch` against already-rendered artifacts.
Combine the two and you lose both abilities for the sake of saving one
line of composition in `RunPipeline`.

## What this rules out

- **Use cases that span pipeline stages.** If a method does ingest *and*
  process, split it. The operator's right to re-enter mid-pipeline is
  more important than the convenience of one method call.
- **Use cases that import infra.** The compile-time check
  (`go list -deps ./internal/app/... | grep infra`) exists for this. If
  a use case "needs" a parser or PDF library, it needs a port instead.
- **Use cases that take raw bytes or framework types.** If the signature
  mentions `[]byte`, `*http.Request`, or `*sql.Tx`, the boundary moved
  to the wrong layer.

## Trade-offs

The cost of this shape is a small amount of ceremony: every new
external capability needs a port before it can be called, and every use
case takes its dependencies through a `Services` field rather than
directly. For a four-stage pipeline that is a tiny tax.

The benefit is that the use cases are the one thing in opyl that does
not have to change when any of the still-open architectural decisions
lands. They describe what the engine does; everything else describes
how.
