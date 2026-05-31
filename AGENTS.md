# AGENTS.md — opyl

opyl is a greenfield Go project: a turn-based play-by-email-style game engine. It ingests player order files, resolves turns deterministically, renders per-player reports (text and PDF), and dispatches them by email. There is no HTTP server, no SPA frontend, and no interactive UI — operator commands are CLI subcommands run manually or by cron.

## Read this before writing code or docs

1. Load the [`applying-sousa`](.agents/skills/applying-sousa/SKILL.md) skill. **opyl follows SOUSA strictly from day one.** It is greenfield; there is no legacy exemption.
2. For any documentation work, load the [`diataxis`](.agents/skills/diataxis/SKILL.md) skill. **opyl documentation follows Diataxis.** See the [Documentation](#documentation) section below for the concrete rules.
3. Read this whole file. It records project-specific decisions and constraints that override or extend the unified skills.
4. When in doubt about placement, prefer the innermost layer that can own the behavior cleanly.
5. For the **game rules** (not the architecture), the authoritative draft is the rulebook at [`docs/content/rules/_index.md`](docs/content/rules/_index.md). It is a **first draft and may contain inconsistencies** — treat it as the design we are converging on, not a settled spec. When it contradicts itself or leaves something open, flag it rather than silently picking an interpretation; do not implement a rule the draft does not clearly establish. Resolved design decisions layered on top of the rulebook are recorded in [`GAME-DESIGN.md`](GAME-DESIGN.md), which in turn feeds the published Diataxis reference pages in `docs/`. The flow is: **rulebook draft → `GAME-DESIGN.md` decisions → reference docs.**

## Project shape

```diagram
╭────────────╮   ╭─────────────╮   ╭─────────────────────╮   ╭─────────╮
│  Domain    │◀──│ Application │◀──│ Infrastructure /    │◀──│ Runtime │
│ (pure game │   │ (use cases, │   │ Delivery            │   │ (cmd/   │
│  state)    │   │  ports)     │   │ (files, render,     │   │  opyl)  │
│            │   │             │   │  mail, CLI)         │   │         │
╰────────────╯   ╰─────────────╯   ╰─────────────────────╯   ╰─────────╯
```

| Layer          | Path                                          | Owns                                                       |
| -------------- | --------------------------------------------- | ---------------------------------------------------------- |
| Domain         | `internal/domain/`                            | Game types, invariants, deterministic transforms           |
| Sentinel errors| `internal/cerr/`                              | `ErrGameNotFound`, `ErrTurnAlreadyProcessed`, etc.         |
| Application    | `internal/app/`                               | Use cases (`IngestOrders`, `ProcessTurn`, …) and ports     |
| Infra (orders) | `internal/infra/orderfile/`                   | Flat-file parsing → `domain.OrderBundle`                   |
| Infra (state)  | `internal/infra/store/`                       | Persistence of game state and turn ledger                  |
| Infra (text)   | `internal/infra/render/text/`                 | `domain.PlayerReport` → text bytes                         |
| Infra (PDF)    | `internal/infra/render/pdf/`                  | `domain.PlayerReport` → PDF bytes                          |
| Infra (mail)   | `internal/infra/mail/`                        | Deliver `domain.Attachment` to a `domain.Recipient`        |
| Delivery       | `internal/delivery/cli/`                      | Thin CLI subcommand handlers                               |
| Runtime        | `cmd/opyl/`                                   | Composition root: parse config, wire adapters, dispatch    |
| Documentation  | `docs/`                                       | Hugo + Hextra site organised by Diataxis (see below)       |

## Ports declared by `internal/app`

These are the seams between layers. Application owns them; infra implements them. **Never import a concrete infra package from app.**

- `OrderSource` — read player orders for a turn
- `GameStateStore` — load/save authoritative game state
- `ReportRenderer` — turn data → bytes + MIME type
- `ReportDispatcher` — send attachment to recipient
- `TurnLedger` — record processed turns for idempotency
- `Clock` — abstract time for determinism

If you need a new external capability, add a small port here first, then implement it in `internal/infra/<adapter>/`.

## Project-specific rules (override or extend the unified skill)

### Input files are untrusted

Player order files arrive over the wire (email, scp, dropbox). The `internal/infra/orderfile/` parser is the boundary: it must reject malformed input and surface `cerr.ErrInvalidOrders`. Validated `domain.OrderBundle` values are what reach `app`. Domain never sees raw bytes.

### Idempotency is an Application concern

Operators rerun things. Every state-mutating use case (`ProcessTurn`, `DispatchReports`) must be safe to invoke twice for the same `(gameID, turn)`. Use `TurnLedger` plus an input hash to short-circuit reruns. Do not push idempotency down into infra adapters — they should be dumb.

### Pipeline stages are independent use cases

`RunPipeline` composes `IngestOrders → ProcessTurn → RenderReports → DispatchReports`. Each stage is its own use case and can be invoked alone via its CLI subcommand. Do not build a single mega-use-case.

### Rendering and dispatch are two separate ports

Never combine "render this report and email it" into one adapter. `ReportRenderer` produces bytes; `ReportDispatcher` sends bytes. The use case wires them. This makes dry-runs (render without sending) trivial.

### Operator trust model (in place of auth)

There is no end-user-facing surface, so there are no JWT, session, or authz concerns. The CLI operator is implicitly trusted. Player identity matters only as routing metadata (`domain.PlayerID` → email), not as a security principal. Drop the auth chapter of the unified skill; it does not apply here.

## Open architectural decisions

These should be decided explicitly before substantial implementation begins. Add a short ADR-style note here when each is settled.

| Decision               | Options under consideration                                                                                                      |
| ---------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| State storage          | SQLite (queryable, concurrent) **vs.** directory of versioned JSON/YAML per turn (human-inspectable, git-diffable audit trail)   |
| PDF library            | `gofpdf` / `signintech/gopdf` (pure Go) **vs.** `typst` CLI (rich layout, external binary) **vs.** `chromedp` (HTML → PDF, heavy)|
| Order file format      | Custom line-oriented DSL **vs.** YAML **vs.** structured email body                                                              |
| Mail transport         | Direct SMTP **vs.** SES / SendGrid API **vs.** "drop EML files in `/outbox` for an external mailer"                              |
| CLI framework          | stdlib `flag` (matches Diacous) **vs.** `cobra` (matches GemGem)                                                                 |
| Concurrency model      | Turns processed serially per game; multiple games in parallel — confirm before adding goroutines                                 |

Whichever choices land, they should affect **only** the relevant `internal/infra/<adapter>/` package. If a decision starts requiring changes outside its infra package, that is a signal the port boundary is wrong — stop and fix the port first.

## Documentation

opyl documentation lives in [`docs/`](docs/) as a Hugo site using the [Hextra](https://imfing.github.io/hextra/) theme (loaded as a Hugo module). The **engine documentation** is organised by the [Diataxis](https://diataxis.fr) framework. Before writing or editing any engine documentation, load the [`diataxis`](.agents/skills/diataxis/SKILL.md) skill and follow its compass.

### The rulebook sits outside Diataxis

`docs/content/rules/` holds the **player-facing rulebook** — game-world rules (the subject opyl simulates), not documentation of the engine. It is intentionally **outside** the Diataxis taxonomy: Diataxis organises docs *about the tool*, and the rules describe the game itself. The rulebook is the **primary entry point for players**, surfaced through its own card in [`docs/content/_index.md`](docs/content/_index.md), kept separate from the four-type Diataxis card block below it. Do **not** fold `rules/` into the four sections below or classify its pages by Diataxis type. (For the rulebook's status as the authoritative *design* draft, see item 5 at the top of this file.)

### The four sections — what goes where

These four sections cover **engine documentation only**. The rulebook (above) is separate.


| Section | Path | Purpose | Voice |
| --- | --- | --- | --- |
| Tutorials | `docs/content/tutorials/` | Learning by guided practice. New operators. | "We will…" |
| How-to guides | `docs/content/how-to/` | Goal-oriented recipes. Competent reader, real task. | "If you want X, do Y." |
| Reference | `docs/content/reference/` | Austere, complete description of the machinery. | "X is. X does." |
| Explanation | `docs/content/explanation/` | Context, design decisions, trade-offs, opinion. | Discussion. |

### Hard rules for documentation

1. **Pick one type per page.** If a page is teaching *and* describing *and* explaining, split it.
2. **Use the compass.** Action vs. cognition × acquisition vs. application. The classification is in each section's `_index.md`.
3. **Reference describes, never discusses.** No "we recommend", no "you should". Move opinions to explanation.
4. **How-to addresses goals, not machinery.** "How to switch the PDF renderer", not "How to use `infra/render/pdf`".
5. **Tutorials must be perfectly reliable.** Every command works exactly as written. No "you might need to…".
6. **Explanation makes connections.** It is the only place where rationale, history, and trade-offs belong.
7. **Cross-link generously between types.** A how-to that needs background should link to explanation, not inline it.
8. **One improvement at a time.** Do not plan grand restructurings. Classify → assess → make one change → ship.

### Common mistakes (and the fix)

| Mistake | Fix |
| --- | --- |
| Tutorial explains too much | Move explanation to `explanation/`, link to it |
| How-to teaches background | Strip to action steps only |
| Reference includes opinions | Move discussion to `explanation/` |
| Explanation gives step-by-step | Move procedure to `how-to/` |
| Mixing types in one page | Split into separate pages by type |

For the full theoretical grounding and decision logic, read the references under [`.agents/skills/diataxis/references/`](.agents/skills/diataxis/references/).

### Working on the docs

```sh
cd docs
hugo mod get -u             # update Hextra and dependencies
hugo server -D              # local preview at http://localhost:1313
hugo --quiet                # production build to public/ (gitignored)
```

The `docs/` directory is its own Hugo module (`docs/go.mod`). It is independent of the Go application module.

## Validation commands

Run these from the `opyl/` directory before declaring work done:

```sh
# Go code
go build ./...
go vet ./...
go test ./...

# Documentation
(cd docs && hugo --quiet)
```

For SOUSA conformance, also spot-check imports:

```sh
# Domain must not import outward.
go list -deps ./internal/domain/... | grep mdhender/opyl/internal/ | grep -v /domain
# (should print nothing)

# App must not import infra, delivery, or runtime.
go list -deps ./internal/app/... | grep -E 'mdhender/opyl/internal/(infra|delivery)' || true
# (should print nothing)
```

## When making a change

### Code

1. Identify the layer that owns the behavior (use the table above).
2. Place new code in the innermost layer that can own it cleanly.
3. If you need an external capability, add or refine a port in `internal/app/ports.go` first.
4. Implement the adapter in the appropriate `internal/infra/<x>/` package.
5. Wire it in `cmd/opyl/main.go`.
6. Add the CLI surface in `internal/delivery/cli/` if there is a new subcommand.
7. Update or add documentation (see below).
8. Run the validation commands above.

### Documentation

When the code change is user-visible, update docs in the **same** change. Use the Diataxis compass to decide where the change lands:

| Kind of change | Update |
| --- | --- |
| New CLI flag, port, error code | `docs/content/reference/` |
| New operator task or developer recipe | `docs/content/how-to/` |
| Newcomer-facing learning experience | `docs/content/tutorials/` |
| New design decision or trade-off worth explaining | `docs/content/explanation/` |

Do not duplicate content across sections; cross-link instead.

## When you are unsure

- Choose the layer that minimizes outward dependencies.
- Keep logic pure first, then add adapters.
- Two ports beats one over-broad port — split before combining.
- If a CLI handler "needs" a file parser or PDF renderer directly, stop. It needs a use case instead.
