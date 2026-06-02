# Architecture Decision Records — opyl

This directory is opyl's home for **architecture decisions**: the choices that shape how the
engine is built, the rationale of record for the settled ones, and the register of those still
open.

## What this is — and why it lives here, not in `docs/content/`

`docs/adr/` sits **outside the Diataxis taxonomy** and **outside the Hugo site** (`docs/content/`).
Diataxis organises documentation by a _product user's_ needs — learning, working, looking up, and
understanding the engine. Architecture decisions serve a different reader: the **contributor or
coding agent building the engine**, who needs to know what was chosen, why, and what is still
undecided. That is governance, not product documentation, so it lives beside
[`AGENTS.md`](../../AGENTS.md) rather than on the published site.

This is the same carve-out that puts the player-facing rulebook outside Diataxis (it documents the
_game world_, not the tool) — applied for the parallel reason: ADRs document _how the tool is
built and constrained_, not how it is used.

## The line between this and its neighbours

- **vs [`AGENTS.md`](../../AGENTS.md)** — AGENTS.md holds **standing rules that are always true**
  regardless of any choice (SOUSA discipline, the layer table, the ports list, the untrusted-input
  and idempotency invariants, validation commands). This directory holds **decisions that could
  have gone another way** — and the record of the ones that did. The test: _"Is this a rule I
  follow, or a choice I made?"_
- **vs [`GAME-DESIGN.md`](../../GAME-DESIGN.md)** — that file resolves **game mechanics** (values,
  phase order, entity attributes). This directory resolves **engine architecture** (storage,
  transport, layout). A decision belongs here the moment it stops being about the game and starts
  being about the program.
- **vs `docs/content/reference/` and `docs/content/explanation/`** — once a decision lands, its
  _descriptive_ result is promoted to `reference/` ("the `OrderSource` port reads orders for a
  turn") and its _rationale for a wide audience_ to `explanation/` ("why idempotency lives in
  app"). The **binding decision itself and its open/closed status stay here**. See the three-way
  split in AGENTS.md's routing rule.

## Contents

- **[Open-decisions register](#open-decisions-register)** — the live table of architecture choices,
  their status, and the constraints the design work has pinned on still-open rows.
- **[Decision records](#decision-records)** — one entry per settled decision: context, choice,
  consequences. As the log grows, entries graduate to numbered files (`0001-…md`) and this README
  becomes the index.

The design layer's reasoning behind these verdicts lives in [`GAME-DESIGN.md`](../../GAME-DESIGN.md)
§13 (the reconciliation register) and the §X.9 "Architectural implications" notes; this directory
is where those verdicts become binding for the build.

## Open-decisions register

Decide each explicitly before substantial implementation begins; add a decision record below when
one settles. Status is reconciled against the design work in GAME-DESIGN §13.

| Decision             | Status        | Options / resolution                                                                                                                                                                                                                                                                                   |
| -------------------- | ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| State storage        | open          | SQLite (queryable, concurrent) **vs.** directory of versioned JSON/YAML per turn (human-inspectable, git-diffable audit trail)                                                                                                                                                                         |
| PDF library          | open          | `gofpdf` / `signintech/gopdf` (pure Go) **vs.** `typst` CLI (rich layout, external binary) **vs.** `chromedp` (HTML → PDF, heavy)                                                                                                                                                                      |
| Order file format    | **decided**   | Custom line-oriented DSL (rulebook envelope) — _not_ YAML, _not_ structured email body. See ADR 0001.                                                                                                                                                                                                  |
| Mail transport       | open          | Direct SMTP **vs.** SES / SendGrid API **vs.** "drop EML files in `/outbox` for an external mailer"                                                                                                                                                                                                    |
| CLI framework        | open          | stdlib `flag` (matches Diacous) **vs.** `cobra` (matches GemGem) — design-neutral, decide at implementation time                                                                                                                                                                                       |
| Concurrency model    | **confirmed** | Turns serial per game; games in parallel; **no goroutines inside a single turn's resolution**. See ADR 0002.                                                                                                                                                                                           |
| Map artifact format  | open          | On-disk format (JSON/YAML/custom) for the authored province graph behind the planned `MapSource` port (GAME-DESIGN §2.1/§13.7)                                                                                                                                                                         |
| Report store format  | open          | Where/how rendered reports persist behind the planned `ReportStore` port; interacts with State storage (GAME-DESIGN §12.7/§13.7)                                                                                                                                                                       |
| JSON results schema  | open          | Versioned projection of `domain.PlayerReport` emailed as machine-readable results; a future SQLite export of results is deferred / out of scope, so the schema must not be over-fitted to email (GAME-DESIGN §12.6/§13.7)                                                                                                                                                                                            |
| `OrderSource` output | open          | One `[]OrderBundle` channel **vs.** the bundle **plus** a separate account-directives struct — `begin`/`unit`/`end`, account/report-format settings, and immediate-effect directives (`resend`/`lore`/`players`/`public`) are account/scan-level, not per-turn unit commands (GAME-DESIGN §10.6/§10.8) |

Whichever choices land, each should affect **only** the relevant `internal/infra/<adapter>/`
package. If a decision starts requiring changes outside its infra package, that is a signal the
port boundary is wrong — stop and fix the port first.

**Constraints the design work has pinned on still-open rows (GAME-DESIGN §13):**

- **State storage** — backend open, but the per-turn snapshot's _contents_ are pinned: it must
  round-trip RNG state, per-unit in-flight command progress, all timer/countdown state, the
  per-location arrival-order list (GAME-DESIGN §13.1), the **entity-number allocation counter** (so
  numbers minted at `FORM`/item creation are a pure function of recorded state, advanced inside
  resolution — GAME-DESIGN §3.8), and **dead-body items with their death turn** (so the 12-turn
  decomposition decay and Noble-Point return resolve deterministically — GAME-DESIGN §3.6/§3.8).
  `TurnLedger`, the report store, and `MapSource` are separate stores, not part of this one.
- **PDF library** — reports are stored and GM-regenerable (GAME-DESIGN §12.7), so deterministic,
  version-stable byte output (same snapshot + code → same bytes) favors a pure-Go library over an
  external binary or headless browser (GAME-DESIGN §13.2).
- **Mail transport** — sits behind _two_ ports: `ReportDispatcher` (outbound) and `OrderSource`
  (inbound order files). `DispatchReports` idempotency lives in app, above transport
  (GAME-DESIGN §13.4).

## Decision records

### ADR 0001 — Order file format (decided)

The order file is the rulebook's custom **line-oriented DSL** — a `begin <player> [password]` …
`unit <number>` blocks … single `end` envelope, forgiving grammar (`#` comments, quoted multi-word
args), `UNIT`-replaces-not-appends semantics, 250 orders/unit cap. It is **not** YAML and **not** a
structured email-field schema. Parsed only in `internal/infra/orderfile/`, the untrusted-input
boundary. The exact tokenizer/grammar spec (quoting edge cases, numeric-vs-entity-code argument
forms) is pinned when that adapter is built. Distinct from the Mail-transport row: the DSL is the
body, mail transport is its carrier. (GAME-DESIGN §10.1/§13.3.)

### ADR 0002 — Concurrency model (confirmed)

`ProcessTurn` is a pure sequential transform — turn N's snapshot is turn N+1's input, so turns of
one game cannot overlap; distinct games share no state and may resolve in parallel. A single turn's
resolution adds **no goroutines**; RNG substream `Split()` keeps any future within-turn fan-out
deterministic. (GAME-DESIGN §11/§13.6.)

### ADR 0003 — Randomness source (resolved)

Stochastic draws go through the **`RNG` port** (`internal/app/ports.go`), mirroring `Clock`,
realized by the **`internal/infra/prng`** PCG adapter. RNG state round-trips with the per-turn
snapshot via `GameStateStore` — randomness stays a pure function of recorded state, and no
component imports an entropy source. This closes the former _Randomness source_ open row; the port
is now listed among AGENTS.md's declared ports, and its descriptive signature lives in
[`reference/ports.md`](../content/reference/ports.md). (GAME-DESIGN §11.7/§11.9.)

**Open follow-ups under the now-decided port** (GAME-DESIGN §11.7/§11.9): the **seed-derivation
rule** and the **substream-assignment scheme** (per game / stage / player, so a battle's stream
does not shift because an unrelated unit rolled earlier) are still open; and for domain-resident
math — the combat exchange (GAME-DESIGN §8.2) — it is open whether the **use case rolls and feeds
outcomes into the domain transform** or a **narrow domain-defined interface is injected**.
Substream `Split()` stays a Runtime wiring-time operation on the concrete adapter until a use case
demonstrably needs mid-turn fan-out, at which point it is promoted to the port.
