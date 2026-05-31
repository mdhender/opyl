# opyl

A greenfield, batch-style turn-based game engine in Go.

opyl reads player order files, resolves turns deterministically, renders per-player reports (text and PDF), and dispatches them by email. There is no HTTP server and no UI — operator commands are CLI subcommands run by hand or by cron.

opyl is the fifth project to adopt **SOUSA** (strict Onion/Clean layering with one-way inward dependencies) and the first non-web SOUSA project. It is greenfield; SOUSA applies from the first commit.

## Pipeline

```diagram
   player order files                                                rendered reports
   (CSV / YAML / DSL)                                                (text or PDF)
          │                                                                  │
          ▼                                                                  ▼
  ╭────────────────╮     ╭────────────────╮     ╭────────────────╮     ╭─────────────╮
  │   ingest       │────▶│   process      │────▶│   render       │────▶│   dispatch  │
  │ (parse + val.) │     │ (resolve turn) │     │ (text or PDF)  │     │ (email)     │
  ╰────────────────╯     ╰────────────────╯     ╰────────────────╯     ╰─────────────╯
                                  │
                                  ▼
                         ╭────────────────╮
                         │  state store   │
                         │ (per-turn      │
                         │  snapshots)    │
                         ╰────────────────╯
```

Each stage is an independent CLI subcommand calling exactly one application use case. An operator can rerun any stage after fixing a problem.

## Layout

```
opyl/
├── AGENTS.md                          ← project-specific rules for coding agents
├── .agents/skills/applying-sousa/     ← unified SOUSA skill (loaded by agents)
├── cmd/opyl/                          ← Runtime: composition root, CLI dispatch
├── internal/
│   ├── domain/                        ← pure game types and invariants
│   ├── cerr/                          ← sentinel errors with business meaning
│   ├── app/                           ← use cases and the ports they declare
│   │   ├── ports.go                   ← OrderSource, GameStateStore, …
│   │   └── services.go                ← IngestOrders, ProcessTurn, …
│   ├── infra/
│   │   ├── orderfile/                 ← flat-file order parsing
│   │   ├── store/                     ← game state + turn ledger persistence
│   │   ├── render/text/               ← text report renderer
│   │   ├── render/pdf/                ← PDF report renderer
│   │   └── mail/                      ← report dispatcher
│   └── delivery/cli/                  ← thin CLI subcommand handlers
├── docs/                              ← Hugo + Hextra site, organised by Diataxis
│   └── content/{tutorials,how-to,reference,explanation}/
└── go.mod
```

## Build

```sh
# Code
go build ./...
go vet ./...
go test ./...

# Docs
(cd docs && hugo --quiet)             # production build
(cd docs && hugo server -D)           # local preview at http://localhost:1313
```

Running the binary today prints a usage message; subcommands are stubs. See [AGENTS.md](AGENTS.md) for the open architectural decisions still to be made (storage backend, PDF library, order file format, mail transport, CLI framework).

## Why SOUSA here

opyl has unusually heavy outer-layer concerns for its size — flat-file parsing, PDF generation, SMTP, persistent state — and unusually little outer surface area (a handful of CLI subcommands). That is exactly the shape where SOUSA's port discipline pays off most: each external concern is a tiny port behind which a heavy library can live and be swapped without touching game rules.

Notable opyl-specific adaptations of the unified SOUSA skill:

- **No auth layer.** The only user is a trusted operator.
- **Idempotency / replay-safety is an Application concern,** not a database concern. A `TurnLedger` port plus input hashing makes use cases safe to rerun.
- **Rendering and dispatch are two ports,** never combined. Dry-runs are free.
- **Input files are untrusted.** The `infra/orderfile` parser is the boundary; only typed orders reach `app`.
- **Storage choice is deferred** behind `app.GameStateStore`. SQLite and per-turn JSON snapshots are both viable; the decision affects only `internal/infra/store/`.

The complete rule set lives in [.agents/skills/applying-sousa/SKILL.md](.agents/skills/applying-sousa/SKILL.md). Project-specific overrides and open decisions live in [AGENTS.md](AGENTS.md).
