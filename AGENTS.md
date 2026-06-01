# AGENTS.md — opyl

opyl is a greenfield Go project: a turn-based play-by-email-style game engine. It ingests player order files, resolves turns deterministically, renders per-player reports (text and PDF), and dispatches them by email. There is no HTTP server, no SPA frontend, and no interactive UI — operator commands are CLI subcommands run manually or by cron.

## Read this before writing code or docs

1. Load the [`applying-sousa`](.agents/skills/applying-sousa/SKILL.md) skill. **opyl follows SOUSA strictly from day one.** It is greenfield; there is no legacy exemption.
2. For any documentation work, load the [`diataxis`](.agents/skills/diataxis/SKILL.md) skill. **opyl documentation follows Diataxis.** See the [Documentation](#documentation) section below for the concrete rules.
3. Read this whole file. It records project-specific decisions and constraints that override or extend the unified skills.

> **Skill layout.** `.agents/skills/` is the **source of truth** for both project skills. Claude Code's Skill tool only auto-discovers skills under `.claude/skills/`, so `.claude/skills/diataxis` and `.claude/skills/applying-sousa` exist as **symlinks** back into `.agents/skills/`. Edit the files under `.agents/skills/`; do not delete the `.claude/skills/` symlinks to "fix" the apparent duplication — removing them breaks `Skill(diataxis)` / `Skill(applying-sousa)` invocation.
4. When in doubt about placement, prefer the innermost layer that can own the behavior cleanly.
5. For the **game rules** (not the architecture), the authoritative draft is the rulebook at [`docs/content/rules/_index.md`](docs/content/rules/_index.md). It is a **first draft and may contain inconsistencies** — treat it as the design we are converging on, not a settled spec. When it contradicts itself or leaves something open, flag it rather than silently picking an interpretation; do not implement a rule the draft does not clearly establish. Resolved design decisions layered on top of the rulebook are recorded in [`GAME-DESIGN.md`](GAME-DESIGN.md), which is then distilled into the published Diataxis reference pages in `docs/`. The flow — and the audience at each stage — is: **rulebook draft (users & game-masters) → [`GAME-DESIGN.md`](GAME-DESIGN.md) decisions (developers & testers) → reference docs (coding agents first; usable by users & game-masters too).** Authority narrows at each step:

- **The reference docs are the *sole* authority for the engine.** Engine code and tests derive **only** from `docs/content/reference/`.
- **`GAME-DESIGN.md` is a translation workbench**, not a build source — it helps us turn rulebook prose into reference facts. Coding agents may *read* it to understand intent, but must not treat it as the source of any rule the code depends on.
- **The rulebook is not authoritative for the engine** at all. It is a player-and-game-master document; the engine must not depend on anything in it.

The practical invariant: **a decision that lives only in `GAME-DESIGN.md` (or the rulebook) and has not yet been distilled into `reference/` is not buildable.** If you need a fact and it is missing from `reference/`, promote it into `reference/` first, then code against `reference/` — never reach back into the design doc or the rulebook.

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
| Infra (rng)    | `internal/infra/prng/`                        | PCG-backed `app.RNG` adapter with persistable state        |
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
- `RNG` — abstract randomness for deterministic dice / stochastic decisions

Two more ports are **planned but not yet declared**, surfaced by the design work (GAME-DESIGN §13.7); add them here when their adapters are built:

- `MapSource` — load the authored province graph as immutable domain input (GAME-DESIGN §2.1/§2.9)
- `ReportStore` — persist / retrieve / remove rendered reports keyed `(gameID, turn, playerID, format)` (GAME-DESIGN §12.7/§12.9)

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

These should be decided explicitly before substantial implementation begins. Add a short ADR-style note here when each is settled. Status is reconciled against the design work in GAME-DESIGN §13.

| Decision               | Status     | Options / resolution                                                                                                             |
| ---------------------- | ---------- | -------------------------------------------------------------------------------------------------------------------------------- |
| State storage          | open       | SQLite (queryable, concurrent) **vs.** directory of versioned JSON/YAML per turn (human-inspectable, git-diffable audit trail)   |
| PDF library            | open       | `gofpdf` / `signintech/gopdf` (pure Go) **vs.** `typst` CLI (rich layout, external binary) **vs.** `chromedp` (HTML → PDF, heavy)|
| Order file format      | **decided**| Custom line-oriented DSL (rulebook envelope) — *not* YAML, *not* structured email body. See ADR below.                          |
| Mail transport         | open       | Direct SMTP **vs.** SES / SendGrid API **vs.** "drop EML files in `/outbox` for an external mailer"                              |
| CLI framework          | open       | stdlib `flag` (matches Diacous) **vs.** `cobra` (matches GemGem) — design-neutral, decide at implementation time                |
| Concurrency model      | **confirmed**| Turns serial per game; games in parallel; **no goroutines inside a single turn's resolution**. See ADR below.                 |
| Map artifact format    | open       | On-disk format (JSON/YAML/custom) for the authored province graph behind the planned `MapSource` port (GAME-DESIGN §2.1/§13.7) |
| Report store format    | open       | Where/how rendered reports persist behind the planned `ReportStore` port; interacts with State storage (GAME-DESIGN §12.7/§13.7)|
| JSON results schema    | open       | Versioned projection of `domain.PlayerReport` emailed as machine-readable results (GAME-DESIGN §12.6/§13.7)                     |

Whichever choices land, they should affect **only** the relevant `internal/infra/<adapter>/` package. If a decision starts requiring changes outside its infra package, that is a signal the port boundary is wrong — stop and fix the port first.

**Constraints noted on still-open rows (GAME-DESIGN §13):**

- **State storage** — backend open, but the per-turn snapshot's *contents* are pinned: it must round-trip RNG state, per-unit in-flight command progress, all timer/countdown state, and the per-location arrival-order list (GAME-DESIGN §13.1). `TurnLedger`, `ReportStore`, and `MapSource` are separate stores, not part of this one.
- **PDF library** — reports are stored and GM-regenerable (GAME-DESIGN §12.7), so deterministic, version-stable byte output (same snapshot + code → same bytes) favors a pure-Go library over an external binary or headless browser (GAME-DESIGN §13.2).
- **Mail transport** — sits behind *two* ports: `ReportDispatcher` (outbound) and `OrderSource` (inbound order files). `DispatchReports` idempotency lives in app, above transport (GAME-DESIGN §13.4).

**ADR — Order file format (decided):** the order file is the rulebook's custom **line-oriented DSL** — a `begin <player> [password]` … `unit <number>` blocks … single `end` envelope, forgiving grammar (`#` comments, quoted multi-word args), `UNIT`-replaces-not-appends semantics, 250 orders/unit cap. Parsed only in `internal/infra/orderfile/`, the untrusted-input boundary. The exact tokenizer/grammar spec is pinned when that adapter is built. (GAME-DESIGN §10.1/§13.3.)

**ADR — Concurrency model (confirmed):** `ProcessTurn` is a pure sequential transform — turn N's snapshot is turn N+1's input, so turns of one game cannot overlap; distinct games share no state and may resolve in parallel. A single turn's resolution adds **no goroutines**; RNG substream `Split()` keeps any future within-turn fan-out deterministic. (GAME-DESIGN §11/§13.6.)

**Resolved & promoted:** the former *Randomness source* row is closed — stochastic draws go through the `RNG` port (above), realized by the `internal/infra/prng` PCG adapter, with RNG state round-tripped in the snapshot via `GameStateStore`. (GAME-DESIGN §11.7/§11.9.)

## Documentation

opyl documentation lives in [`docs/`](docs/) as a Hugo site using the [Hextra](https://imfing.github.io/hextra/) theme (loaded as a Hugo module). The **engine documentation** is organised by the [Diataxis](https://diataxis.fr) framework. Before writing or editing any engine documentation, load the [`diataxis`](.agents/skills/diataxis/SKILL.md) skill and follow its compass.

### The rulebook sits outside Diataxis

`docs/content/rules/` holds the **player-facing rulebook** — game-world rules (the subject opyl simulates), not documentation of the engine. It is intentionally **outside** the Diataxis taxonomy: Diataxis organises docs *about the tool*, and the rules describe the game itself. The rulebook is the **primary entry point for players and game-masters**, surfaced through its own card in [`docs/content/_index.md`](docs/content/_index.md), kept separate from the four-type Diataxis card block below it. Do **not** fold `rules/` into the four sections below or classify its pages by Diataxis type.

The rulebook is **not authoritative for the engine** — engine code and tests derive **only** from the reference docs, never from `rules/` or `GAME-DESIGN.md` (see item 5 at the top of this file for the full authority pipeline). This is why a glossary of *game-world* terms belongs here in `rules/`, not in the engine `reference/` section, even though Diataxis would otherwise file a glossary under reference. By contrast, the **reference** section documents the engine for **coding agents first** — it is the sole authority they build from — but it is also where players and game-masters look for engine-facing facts they need (CLI usage, error codes), so reference pages serve that wider audience too.

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
