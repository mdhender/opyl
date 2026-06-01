# Migration plan — prompts for the GAME-DESIGN architecture extraction

This file holds **session prompts**, not docs. Each fenced block below is **self-contained**:
paste one into a fresh session to run that step of the migration that moves GAME-DESIGN.md's
architecture content (the §X.9 "Architectural implications" blocks and §13) out to their
correct homes, leaving GAME-DESIGN as pure game-design. It is the follow-up to commit
`0d7e28f` (which added `docs/adr/` and the doc routing rule). **Delete this file in the §13
closeout step, once the migration is done.**

## Progress

- **§2 (Map) — done** (commit `18086f7`). This was the pattern-validation step. The proven
  shape, reused by every block below:
  - Both §2.9 items were already represented in `docs/adr/` (the "Map artifact format"
    register row + planned `MapSource` port; ADR 0003 for the `RNG` port), so nothing was
    added there — **verify-don't-duplicate** held.
  - §2.9 was **replaced by a pointer** that keeps its `### 2.9` heading/anchor (other sections
    link to it) and routes each consequence to its home with a link.
  - The chapter's stale "belong in AGENTS.md's open-decisions table" cross-ref was repointed
    at `docs/adr/`.
- **§3–§12 — todo.** One chapter per session (batch small/adjacent chapters — see the
  template). Architecture blocks are at: §3.8, §4.9, §5.9, §6.8, §7.9, §8.10, §9.8, §10.8,
  §11.9, §12.9.
- **§13 — todo, last.** The reconciliation register can only be retired once every §3–§12
  verdict is confirmed present in `docs/adr/`. Use the closeout block.

---

## Reusable per-section prompt (§3–§12)

Replace **`§N`** in the first line with the chapter(s) for this session — a single chapter
(e.g. `§3`) or a small batch of adjacent ones (e.g. `§4–§5`) when the blocks are thin. Then
paste the whole block into a fresh session.

```
We're migrating GAME-DESIGN.md's architecture content out of the design doc, per the doc
routing rule in commit 0d7e28f. This session does §3 only.

Before starting, read: AGENTS.md (esp. "Documentation → the routing rule" and the "'Architecture'
is not one thing — triage it three ways" subsection), item 5 at the top of AGENTS.md (the
authority pipeline: rulebook → GAME-DESIGN → reference/, authority narrowing at each step), and
docs/adr/README.md. Load the `applying-sousa` and `diataxis` skills.

## Background

GAME-DESIGN.md carries two kinds of content: game-design decisions (its real job) and
engine-architecture residue that doesn't belong there — the §X.9 "Architectural implications"
subsection closing most chapters, and all of §13. The architecture table and its ADRs have
moved to docs/adr/. The goal is to leave GAME-DESIGN.md as PURE game-design, with the
architecture content triaged to its correct homes. §2 is already done (commit 18086f7) and
established the pattern; this session applies it to §N.

## The triage rule (three-way split)

Route each §X.9 item by the routing rule in AGENTS.md:
- Descriptive ("the OrderSource port reads orders for a turn") → docs/content/reference/
- Rationale / trade-off ("why idempotency lives in app") → docs/content/explanation/
- Decision / constraint / open question ("SQLite vs per-turn JSON — open") → docs/adr/
  (or AGENTS.md if it's a standing always-true rule, not a choice that could have gone
  another way)

Constraints:
- reference/ is the SOLE authority the engine builds from. Promote a fact there only if it's
  a DECIDED (✅) game fact — never an open (❓) one. Don't invent rules the draft doesn't
  establish; flag contradictions instead of resolving them silently.
- Don't delete §X.9 content — relocate it, and leave a one-line pointer/link from GAME-DESIGN
  so traceability survives.
- Preserve the design-decision content in the chapter (the §X.1–§X.7 body) untouched. Only
  the §X.9 architecture residue moves.
- Verify, don't duplicate: check the target already holds the fact before copying anything.
  Much may already be distilled.
- Don't touch §13 yet — it's retired in the final closeout step.

## The proven pattern (from §2)

1. Read the chapter's architecture block plus any reference/ and explanation/ pages that
   already cover the chapter (check docs/content/reference/ and docs/content/explanation/ for
   pages matching the chapter's topic). For ports, check docs/content/reference/ports.md; for
   open decisions, check docs/adr/README.md's register and ADRs.
2. Triage each bullet three ways. For decision/open items, CONFIRM they're already in
   docs/adr/ (register row or ADR); add only what's genuinely missing. For descriptive facts,
   confirm they're in reference/ — promote a missing DECIDED fact, skip open ones. For
   rationale, confirm it's in explanation/ — promote if missing.
3. Replace the §X.9 block with a short pointer that KEEPS the heading/anchor (e.g. keep
   `### X.9 Architectural implications`; some chapters use `.8` or `.10`) and routes each
   consequence to its home with a link. Keep the anchor if any other section links to it.
4. Fix that chapter's stale "…belong in / join … AGENTS.md's 'Open architectural decisions'
   table" cross-ref to point at docs/adr/ instead.
5. Validate: `(cd docs && hugo --quiet)` must succeed; `go build ./...` (no Go changes
   expected, but confirm). If you promoted anything into docs/, also rebuild docs.
6. Report what moved where (and what was already covered, so nothing was duplicated). Then
   STOP and summarize for review.

Commit §3 on its own (directly to main) once I approve.
```

---

## §13 closeout prompt (run last)

Run this only after §3–§12 are all done and committed. Paste into a fresh session.

```
This is the final step of the GAME-DESIGN architecture migration (per commit 0d7e28f): retire
§13 and remove the migration scaffolding. §2–§12 are already done — their §X.9 blocks are now
pointers into docs/adr/.

Before starting, read: AGENTS.md ("Documentation → the routing rule" and the three-way split),
item 5 at the top of AGENTS.md, and docs/adr/README.md. Load the `applying-sousa` and
`diataxis` skills.

## What §13 is

GAME-DESIGN §13 is the reconciliation register — it mapped each design verdict back to
AGENTS.md's OLD open-decisions table (now moved to docs/adr/). It is pure architecture
apparatus and must leave GAME-DESIGN entirely, but ONLY after confirming every verdict it
records already lives in docs/adr/.

## Steps

1. Read GAME-DESIGN §13 in full (from its heading to end of file).
2. For each row/verdict in §13, CONFIRM it is already represented in docs/adr/ — the
   open-decisions register row (with any pinned constraints) or a numbered ADR. The §13
   subsections map to docs/adr/ as: §13.1 → State storage constraints; §13.2 → PDF library
   constraints; §13.3 → ADR 0001 (order format); §13.4 → Mail transport constraints; §13.6 →
   ADR 0002 (concurrency); §13.7 → the MapSource / ReportStore / JSON-results-schema register
   rows; plus the RNG verdict → ADR 0003. Verify each is present and faithful.
3. If anything in §13 is NOT yet in docs/adr/, ADD it there first (a register row, a pinned
   constraint, or a new ADR) — do not drop a verdict on the floor. Flag, don't silently
   resolve, any §13 statement that contradicts docs/adr/.
4. Once every verdict is confirmed present in docs/adr/, REMOVE §13 from GAME-DESIGN.md
   (heading and body). Check for inbound links to §13 anchors from earlier chapters and
   repoint them at docs/adr/ (grep for "§13" across the repo).
5. Update GAME-DESIGN's front-matter/intro if it still claims to "resolve the open decisions
   tracked in AGENTS.md" — that job now lives in docs/adr/. Update the docs/adr/README.md
   lines that reference "GAME-DESIGN §13" as the reasoning home if they're now stale.
6. Delete docs/adr/MIGRATION-PLAN-PROMPT.md (this file) — the migration is complete.
7. Validate: `(cd docs && hugo --quiet)` and `go build ./...` must succeed. Grep GAME-DESIGN
   for any remaining "AGENTS.md's 'Open architectural decisions' table" or "Architectural
   implications" residue — there should be none beyond the §X.9 pointer stubs.
8. Report what was confirmed already-present vs. newly added to docs/adr/, and confirm
   GAME-DESIGN is now pure game-design. Commit on its own once I approve.
```
