# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Read first

**[AGENTS.md](AGENTS.md) is the authoritative source for this project's rules** — layer table, ports, project-specific overrides, open architectural decisions, and documentation conventions. Read it before writing code or docs. This file is a quick orientation; AGENTS.md wins on any conflict.

Before writing code, load the **[`applying-sousa`](.agents/skills/applying-sousa/SKILL.md)** skill. Before writing docs, load the **[`diataxis`](.agents/skills/diataxis/SKILL.md)** skill. opyl follows both strictly — it is greenfield with no legacy exemption.

## What this is

opyl is a turn-based play-by-email game engine in Go (`module github.com/mdhender/opyl`, go 1.23). It ingests player order files, resolves turns deterministically, renders per-player reports (text + PDF), and dispatches them by email. **No HTTP server, no UI, no auth** — operators run CLI subcommands by hand or cron. The codebase is currently mostly stubs (`doc.go` per infra package; `main.go` prints usage); most architectural decisions in AGENTS.md's "Open architectural decisions" table are not yet settled.

## Commands

```sh
# Go code (run from repo root before declaring work done)
go build ./...
go vet ./...
go test ./...
go test ./internal/app/ -run TestProcessTurn   # single package / test

# Docs (docs/ is its own Hugo module, independent of the Go module)
(cd docs && hugo --quiet)        # production build
(cd docs && hugo server -D)      # local preview at http://localhost:1313

# SOUSA import conformance — both must print nothing
go list -deps ./internal/domain/... | grep mdhender/opyl/internal/ | grep -v /domain
go list -deps ./internal/app/... | grep -E 'mdhender/opyl/internal/(infra|delivery)'
```

## Architecture (the non-obvious parts)

Dependencies flow **inward only**: Domain ← Application ← Infra/Delivery ← Runtime. Inner layers never import outer ones; Infra and Delivery are peers and never import each other — add a port instead.

- `internal/domain/` — pure game types, invariants, deterministic transforms. No I/O, no `time.Now`, no randomness.
- `internal/cerr/` — sentinel errors with business meaning (`ErrInvalidOrders`, `ErrTurnAlreadyProcessed`, …).
- `internal/app/` — use cases (`services.go`) and the **ports** they declare (`ports.go`): `OrderSource`, `GameStateStore`, `ReportRenderer`, `ReportDispatcher`, `TurnLedger`, `Clock`. App never imports a concrete infra package.
- `internal/infra/{orderfile,store,render/text,render/pdf,mail}/` — adapters implementing the ports. Each external library lives behind exactly one port and is swappable without touching game rules.
- `internal/delivery/cli/` — thin subcommand handlers.
- `cmd/opyl/main.go` — composition root: the only place that knows every layer; wires adapters into app services and dispatches.

Pipeline: `ingest → process → render → dispatch`, each an independent use case invocable as its own CLI subcommand, plus a `pipeline` that runs all four. State is persisted as per-turn snapshots.

### Rules that shape every change

- **Order files are untrusted.** `internal/infra/orderfile/` is the validation boundary; it rejects malformed input with `cerr.ErrInvalidOrders`. Only typed `domain.OrderBundle` values reach `app` — domain never sees raw bytes.
- **Idempotency is an Application concern, not infra's.** Every state-mutating use case must be safe to rerun for the same `(gameID, turn)` via `TurnLedger` + an input hash. Keep infra adapters dumb.
- **Render and dispatch are two separate ports**, never combined — this keeps dry-runs (render without send) free.
- **New external capability → add/refine a port in `app/ports.go` first**, then implement the adapter in `internal/infra/<x>/`, then wire it in `cmd/opyl/main.go`. If a CLI handler seems to need a parser or renderer directly, it needs a use case instead.
- When code is user-visible, update `docs/` in the same change, placing it by Diataxis type (reference / how-to / tutorial / explanation) — see AGENTS.md.
