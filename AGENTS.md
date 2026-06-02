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
| Documentation  | `docs/content/`                               | Hugo + Hextra site organised by Diataxis (see below)       |
| Arch. decisions| `docs/adr/`                                   | ADRs + open-decisions register; outside Diataxis & the site |

## Ports declared by `internal/app`

These are the seams between layers. Application owns them; infra implements them. **Never import a concrete infra package from app.**

- `OrderSource` — read player orders for a turn
- `GameStateStore` — load/save authoritative game state
- `ReportRenderer` — turn data → bytes + MIME type
- `ReportDispatcher` — send attachment to recipient
- `TurnLedger` — record processed turns for idempotency
- `Clock` — abstract time for determinism
- `RNG` — abstract randomness for deterministic dice / stochastic decisions

Two more ports are **planned but not yet declared**, surfaced by the design work (see the open-decisions register in [`docs/adr/`](docs/adr/README.md)); add them here when their adapters are built:

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

Architecture decisions — the choices that shape how the engine is built, the records of the settled ones, and the register of those still open — live in **[`docs/adr/`](docs/adr/README.md)**, not in this file. AGENTS.md holds the rules that are **always true** (the layer table, ports, SOUSA discipline, the invariants above); `docs/adr/` holds the choices that **could have gone another way**. See the routing rule under [Documentation](#documentation) for the full split.

Currently open (decide before substantial implementation): **state storage**, **PDF library**, **mail transport**, **CLI framework**, **map-artifact format**, **report-store format**, and the **JSON results schema**. Decided: **order file format** (custom line-oriented DSL, ADR 0001), **concurrency model** (serial-per-game, ADR 0002), **randomness source** (`RNG` port, ADR 0003). The live register — options, statuses, and the constraints the design work has pinned on still-open rows — is in [`docs/adr/README.md`](docs/adr/README.md).

Whichever choices land, each should affect **only** the relevant `internal/infra/<adapter>/` package. If a decision starts requiring changes outside its infra package, that is a signal the port boundary is wrong — stop and fix the port first.

## Documentation

Engine documentation lives in [`docs/content/`](docs/content/) as a Hugo site using the [Hextra](https://imfing.github.io/hextra/) theme (loaded as a Hugo module), organised by the [Diataxis](https://diataxis.fr) framework. Before writing or editing any engine documentation, load the [`diataxis`](.agents/skills/diataxis/SKILL.md) skill and follow its compass. But Diataxis governs only the **engine-as-a-product** docs — the routing rule below places the game-world and the project-governance writing that sits outside it.

### The routing rule — every doc has one home

Every piece of writing in this repo has **exactly one home**. Decide it with two questions, in order:

1. **Whose need does it serve?** The *game world* (players/GMs), the *engine as a product* (operators, integrators, and the coding agents that build against it), or the *project itself* (how we build it — contributors and agents)?
2. **If it serves the engine-as-a-product, run the Diataxis compass** — action vs. cognition × acquisition vs. application — to pick among the four types.

| Home | What belongs here (and its voice) | Never goes here |
| --- | --- | --- |
| `docs/content/rules/` | The **game-world rulebook** — the fiction and rules opyl simulates. Player & GM facing. Outside Diataxis. *Narrative/rules voice.* | Engine facts, architecture, code or agent guidance. |
| `GAME-DESIGN.md` | Resolved **game-design** decisions converging the draft rulebook into concrete mechanics (values, phase order, attributes). A workbench for devs/testers — **not a build source**. *"We decided the game does X."* | Engine architecture, ports, infra choices, ADRs. "Architectural implications" → `docs/adr/`, `reference/`, or `explanation/` per the three-way split below. |
| `docs/content/reference/` | Austere **description** of the engine machinery a reader consults while working — CLI, error codes, the ports catalog, the domain model. **The sole authority engine code & tests build from.** *"X is. X does."* | "You must…", "we decided…", rationale, open questions, trade-offs. |
| `docs/content/explanation/` | The **why**: design rationale, trade-offs, how the pieces connect. Answers *"Can you tell me about…?"* Architecture *rationale* lives here. *Discussion.* | Step-by-step procedure, austere fact catalogs, binding rules. |
| `docs/content/how-to/` | Goal-oriented **recipes** for an operator or developer completing a real task. *"If you want X, do Y."* | Background teaching, machinery tours. |
| `docs/content/tutorials/` | Guided, perfectly-reliable **first-run learning** for newcomers. *"We will…"* | Options, alternatives, explanation, edge cases. |
| `AGENTS.md` | **Standing rules that are always true** regardless of any choice — the layer table, ports list, SOUSA/Diataxis discipline, the untrusted-input and idempotency invariants, validation commands, this routing rule, the change procedure. Governance; outside Diataxis. | Decisions among alternatives & their records → `docs/adr/`. Game mechanics → `GAME-DESIGN.md`. |
| `docs/adr/` | **Architecture decisions** — settled ADRs (choice + rationale of record), the open-decisions register, binding build-time constraints. Governance; outside Diataxis **and** outside the Hugo site. | Game mechanics, product-user docs, standing always-true discipline (that is AGENTS.md's). |

### "Architecture" is not one thing — triage it three ways

When a game-design decision has an engine consequence, that consequence does **not** belong in one place. It splits across the Diataxis edge, and letting all three sit together as an "architectural implications" blob is the mixing-types mistake the compass forbids — applied to docs that straddle the framework's boundary:

- **Descriptive** — *"the `OrderSource` port reads orders for a turn"* → **`reference/`** (it mirrors the product structure; consulted at work).
- **Rationale** — *"why idempotency lives in app, and the trade-off"* → **`explanation/`** (*"Can you tell me about…?"*).
- **Decision / constraint** — *"SQLite vs. per-turn JSON — open"*; *"never combine render + dispatch"* → **`docs/adr/`**, unless it is a standing always-true rule, in which case **AGENTS.md**.

The bright line between the two governance homes: **AGENTS.md holds rules you follow (always true); `docs/adr/` holds choices you made (could have gone another way), and the ones still open.**

### The rulebook sits outside Diataxis

`docs/content/rules/` holds the **player-facing rulebook** — game-world rules (the subject opyl simulates), not documentation of the engine. It is intentionally **outside** the Diataxis taxonomy: Diataxis organises docs *about the tool*, and the rules describe the game itself. The rulebook is the **primary entry point for players and game-masters**, surfaced through its own card in [`docs/content/_index.md`](docs/content/_index.md), kept separate from the four-type Diataxis card block below it. Do **not** fold `rules/` into the four engine sections or classify its pages by Diataxis type.

The rulebook is **not authoritative for the engine** — engine code and tests derive **only** from the reference docs, never from `rules/` or `GAME-DESIGN.md` (see item 5 at the top of this file for the full authority pipeline). This is why a glossary of *game-world* terms belongs here in `rules/`, not in the engine `reference/` section, even though Diataxis would otherwise file a glossary under reference. By contrast, the **reference** section documents the engine for **coding agents first** — it is the sole authority they build from — but it is also where players and game-masters look for engine-facing facts they need (CLI usage, error codes), so reference pages serve that wider audience too.

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
# Domain must not import outward. cerr is part of the innermost layer
# (sentinel errors are a Domain concern), so domain may import it; the
# check excludes both domain and cerr and flags any other internal import.
go list -deps ./internal/domain/... | grep mdhender/opyl/internal/ | grep -vE '/(domain|cerr)'
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
