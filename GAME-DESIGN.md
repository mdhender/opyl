# GAME-DESIGN.md — opyl

The record of **resolved game-design decisions** layered on top of the draft rulebook
([`docs/content/rules/_index.md`](docs/content/rules/_index.md)). The rulebook is the
authoritative draft of the rules; this file captures the decisions we make as we converge
that draft — where it is silent, contradicts itself, or needs a concrete value chosen.

This file is design, not published documentation. As sections stabilize they feed the
Diataxis docs in `docs/content/`:

- Mechanical facts (order syntax, phase order, entity attributes) → **reference** pages
- Rationale and trade-offs ("why X resolves before Y") → **explanation** pages
- It also resolves the open decisions tracked in [AGENTS.md](AGENTS.md).

Status legend: ✅ decided · 🟡 partially specified · ❓ open question, undecided.

## 1. Concept & victory

opyl is an **open-ended fantasy game** derived from Skrenta's Olympia. ✅

## 2. The map 🟡

The world is a square grid of **provinces** grouped into named **regions**. Provinces may
contain **inner locations** (cities, inns, ports, …). This section records the decisions
that turn the [Geography & Movement](docs/content/rules/geography.md) rulebook draft into a
buildable model. Spatial flavor (terrain *yields*, special realms) is deferred where noted.

### 2.1 Representation & source ✅

- The map is a **fixed, authored artifact** loaded by the engine — a static province graph,
  not procedurally generated. Rationale: deterministic, human-inspectable, git-diffable, and
  consistent with the domain rule "no randomness, no I/O." The world is authored once.
- Loading the artifact is an **infra concern** (an adapter); `domain` holds the in-memory
  province-graph type and treats it as immutable input to turn resolution. See
  [Architectural implications](#29-architectural-implications) — this likely wants a new
  `MapSource` port; the on-disk **format is undecided** (carry to AGENTS.md's open-decisions
  table).
- Map **dimensions are a property of the authored map**, not fixed by the rules.

### 2.2 Coordinates & addressing ✅

- A province's identity is **numeric, one-based `(row, col)`**. The bracketed code (`[a1]`,
  `[aa1]`) is **cosmetic display formatting**, not the identity.
- The coordinate space runs from `(1, 1)` in the **top-left (NW) corner** to `(n, n)`; rows
  increase **south**, columns increase **east**. Map dimensions are the GM's choice (§2.1).
- **Display code = row letters + column number**, "compressed" — no fixed width, no leading
  zeros. The top-left province `(1, 1)` renders as **`a1`**, *not* the rulebook's `aa00`
  (which is dropped).
  - **Row** is a **bijective base-22** numeral over the alphabet
    `a b c d e f g h k m n p q r s t u v w x y z` (a–z minus `i`, `j`, `l`, `o`, which read as
    `1`/`0`). So `a`=1 … `z`=22, then `aa`=23, `ab`=24, …; **two-letter rows begin at row 23**
    (bijective numbering has no zero digit). No hard cap — the code simply grows a letter
    (`aaa`=507) for very large maps. This alphabet **supersedes** the rulebook's
    `abcdfghjkmnpqrstvwxz` list, which excluded `e` yet used `[ae48]` as its main example.
  - **Column** is a plain decimal ordinal (`1`, `2`, …), written without leading zeros.
- **GM convention:** the world origin is the GM's choice, but maps are typically laid out with
  **`aa1` at the upper-left = `(row 23, col 1)`**, leaving rows 1–22 (`a`–`z`) as northern
  margin. The engine stores only `(row, col)` and is indifferent to the convention.
- Map edges are impassable. Sub-location codes are arbitrary and carry **no** coordinate
  meaning.

> **Rulebook items still to reconcile:** the rulebook's coordinate prose (`aa00` origin,
> two-letters-plus-two-digits, the `abcdfghjkmnpqrstvwxz` row sequence) and its ASCII grid all
> predate §2.2 and must be regenerated against this one-based, compressed scheme — 🟡 TODO.
> Already applied this pass: the `sail south` typo fix in `geography.md` (was `sail e`).

### 2.3 Terrain ✅ / 🟡

- **Six terrain types**: plains, forest, swamp, mountain, desert, ocean. ✅
- Terrain affects **movement cost** (§2.4). Other terrain **effects** (resource yield,
  defense, sighting) are **deferred to the economy/combat passes** — 🟡.

### 2.4 Movement & travel time 🟡

- Movement is in the **four orthogonal directions**; diagonals cost two `MOVE`s. Land travel
  auto-selects the **fastest available mode** (horseback when the whole party can be mounted;
  rough terrain may negate the horse benefit). Ocean travel requires a ship. ✅
- Each route carries a **nominal cost in days** (authored per route). ✅
- Actual cost = nominal, modified by transport mode, terrain, and **weather**. Per the
  travel-time decision, this variance is **deterministic given a per-turn seed recorded in
  the `TurnLedger`** (see §2.7). ✅ *mechanism* — the concrete **variance model**
  (distributions, per-mode/terrain modifiers, wind for ocean) is ❓ and designed later.
- `FLY` exists in the command set but its movement rules are unspecified here — cross-ref the
  Orders pass. ❓

### 2.5 Inner locations, ports & visibility ✅

- Provinces may hold **sub-locations** (cities, inns, ports, islands), entered from the
  surrounding province; `MOVE IN` enters the first listed when unambiguous.
- **Visibility**: occupants of a sub-location see the surrounding province; outsiders cannot
  see *into* a sub-location without entering. Co-located characters interact without travel.
- **Ocean ports / port cities**: a ship docks into an adjoining land province (1 day); it
  cannot dock against mountains (ocean↔mountain routes are `impassable`). A **port city**
  gates ocean access — the surrounding province cannot reach the ocean directly.

### 2.6 Holes & hidden routes ✅

- **Holes**: a province may have no route in a given direction (impassable / undiscovered).
- **Hidden routes**: discoverable via `EXPLORE`; once found, usable by the **whole owning
  faction** but by **no other faction** (even one that knows the destination code).
  Stack-mates crossing a hidden route learn it (prisoners excepted).

### 2.7 Civilization level ✅

- Every province has a civ level; **0 = wilderness**, no upper cap.
- Per turn: `civ(p) = max( buildings(p), floor( maxNeighborCiv / 2 ) )`, where
  - `buildings(p)` sums the contribution table (Safe Haven 2; Castle 1.5 + improvement/4;
    City / Tower / Temple / Inn / Mine each 1), counting only the **first of each type**,
    fractions dropped after summing;
  - `maxNeighborCiv` is the max civ among the **four orthogonal neighbors** (off-map and hole
    neighbors count as 0), **read from the previous turn's values** — a **single pass, no
    fixpoint**. Civilization therefore spreads **one hop per turn**.
- "Surrounding provinces" is pinned to the 4 orthogonal neighbors (consistent with movement);
  this was unspecified in the rulebook. 🟡 confirm the gradual one-hop spread is desired.
- Initialization: turn-zero civ comes from the authored map; absent an authored value, the
  first computation uses `buildings(p)` only.

### 2.8 Regions & special realms ❓ (deferred)

- **Regions** are named province groupings; whether a region is a mechanical entity with
  attributes (vs. a label) is undecided — deferred.
- **Hades, Faery, the Cloudlands** are lore-specified with partial mechanics (Faery Hunt
  combat ratings, flight-only Cloudlands). Treated as **later content**, not part of the core
  map pass.
- **Safe haven** placement/count and "no combat or magic" enforcement: noted, designed with
  the combat/realm passes.

### 2.9 Architectural implications

These follow from §2 and belong in AGENTS.md's "Open architectural decisions" table (not yet
added — offer pending):

- **Map artifact format + loader.** A `MapSource` port read at composition time; the on-disk
  format (JSON/YAML/custom) is undecided. The province graph is immutable input to the domain.
- **Turn seed source.** `ProcessTurn` takes a recorded seed as **input** (from the
  `TurnLedger`); the domain derives all travel-time variance from it via a seeded PRNG.
  Randomness stays a pure function of recorded state — the domain still imports no entropy
  source. This is the same discipline the `Clock` port applies to time.

## 3. Factions and Nobles ❓

| Concept             | Specified                                              | Open |
| ------------------- | ------------------------------------------------------ | ---- |
| **Faction**         | All units controlled by a player; the player's faction.| ❓   |
| **Noble Character** | The core entity and the unit that accepts orders.      | ❓   |

## 4. Orders ❓

## 5. Turn resolution ❓

## 6. Turn reports ❓

## 7. Open decisions carried from AGENTS.md ❓
