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

- opyl is an **open-ended fantasy game** derived from Skrenta's Olympia. ✅
- **No victory conditions, no winner.** Characters have no fixed goals; players pursue whatever
  ends they choose, and the game **never declares a winner** ([playing.md](docs/content/rules/playing.md)).
  The engine therefore implements **no scoring, no win-check, and no end-of-game** logic — turns
  run indefinitely until the GM stops the game. ✅

## 2. The map 🟡

The world is a square grid of **provinces** grouped into named **regions**. Provinces may
contain **inner locations** (cities, inns, ports, …). This section records the decisions
that turn the [Geography & Movement](docs/content/rules/geography.md) rulebook draft into a
buildable model. Spatial flavor (terrain *yields*, special realms) is deferred where noted.

> **Distilled →** the decided (✅) facts of this section are published as the
> [Map reference](docs/content/reference/model/map.md) — the sole source the engine builds
> from. Promoted so far: representation (§2.1), coordinates (§2.2), terrain types (§2.3), the
> decided parts of movement (§2.4), inner locations & ports (§2.5), holes & hidden routes
> (§2.6), and the civilization formula (§2.7). **Not yet distilled** (revisit on each pass):
> the §2.4 variance model and `FLY` rules (❓), and §2.8 regions/special realms (deferred).
> When any of these are decided here, promote them into the Map reference before coding.

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

> **Rulebook reconciled ✅:** the coordinate prose, ASCII grid, row sequence, and edge example
> in `geography.md`, plus the *Map coordinates* glossary entry, have been regenerated against
> this one-based, compressed scheme — `a1` origin, variable width, no leading zeros, and the
> `abcdefghkmnpqrstuvwxyz` row letters. The earlier `sail south` typo fix in `geography.md`
> (was `sail e`) landed in the same pass.

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

## 3. Factions and Nobles 🟡

The entity model the rest of the engine bottoms out on. §2 gave us the **places**; §3
defines the **entities that occupy them** — who accepts orders, what they own, and how they
are born and die. This section fixes the *shape* of the model (which entities exist, what
attributes each carries, how they nest) and **defers the mechanics** that act on that shape to
the passes that consume them — orders (§4), turn resolution (§5), and the not-yet-opened
skills/items/combat passes.

Primary source: [Playing the Game](docs/content/rules/playing.md) ("Definition of Terms"),
with [Loyalty, Stacking & Upkeep](docs/content/rules/logistics.md) and
[Health & Death](docs/content/rules/health-death.md) supplying noble/faction attributes. The
glossary entries for *Faction*, *Noble*, *Men*, *Stack*, *Unit*, and *Player entity* anchor the
vocabulary.

### 3.1 Faction & player entity ✅ / 🟡

- A **faction** is the set of **all nobles controlled by one player** — the player's whole
  position in the game. ✅ One player controls exactly one faction; the rulebook's anti-cheating
  rule (one faction per player, unique email) makes this a **1:1 player↔faction** invariant. ✅
- A faction is its **own aggregate**, distinct from any single noble. It is the home of
  **faction-level state** that no individual noble owns:
  - the **Noble-Point pool** (§3.7),
  - the roster of member nobles,
  - the addressable **player entity** — an invisible placeholder that takes a few faction-scoped
    orders (press, rumor, diplomatic forwarding) and **none** of a character's usual commands. It
    is the target of `#forwardto` mail and the byline for `PRESS`/`RUMOR`. 🟡 (its exact order set
    is an Orders-pass concern.)
- A faction **holds no items, men, or location of its own** — only nobles do (§3.4). Territory a
  faction "controls" is **derived** from where its nobles and their buildings sit (a Provinces-pass
  concern), not stored on the faction.

### 3.2 Noble identity & the entity-number space 🟡

- **Noble** and **Character** are synonyms; the domain's canonical term is **noble**. ✅ A noble is
  the only **unit that accepts orders** — every player order is addressed to a noble. ✅
- A noble's **identity is its entity number** — an opaque integer shown in brackets after its name
  (`Osswid the Destroyer [5499]`). Unlike a **province**, whose identity is its `(row, col)` and
  whose bracketed code is merely cosmetic (§2.2), **for a noble the number *is* the identity**.
- **Namespace split (decided):** addressable entities fall into **two** identity spaces —
  - **provinces** → spatial `(row, col)` (§2.2);
  - **everything else** (nobles, items, skills, sub-locations) → a single **entity-number**
    namespace.

  Well-known fixtures take low fixed numbers (gold `[1]`, Shipcraft `[600]`); dynamically created
  entities (a FORMed noble, a found scroll) are minted a fresh number. The **on-disk numbering
  scheme** (decimal vs. the rulebook's base-N alphanumeric like `[yq12]`) and the **allocation
  policy** are ❓ — deferred, but constrained: allocation must be a **pure function of recorded
  state** (see §3.8), never `rand`/`time`.
- A noble also carries a **display name** — free text, player-chosen via order (as for other named
  entities). Name is **cosmetic**; the number is identity. Player-supplied names are **untrusted input**
  and must be sanitized before reaching any report — see §4.

### 3.3 Noble attributes ✅ (slots) / 🟡 (mechanics)

The decision here is the **attribute set** — which slots a noble carries. Whether each slot's
*mechanics* are settled is marked per row; an unsettled mechanic still reserves its slot now so §4/§5
have a stable target.

| Attribute            | Slot | Mechanics settled?                                                              |
| -------------------- | ---- | ------------------------------------------------------------------------------- |
| Entity number, name  | ✅   | ✅ identity (§3.2)                                                               |
| Location             | ✅   | ✅ province/sub-location placement (§2.5)                                        |
| Stack position       | ✅   | 🟡 grouping model below; movement/combat effects deferred to §5/combat pass     |
| Loyalty bond         | ✅   | 🟡 kind + rating decided (§3.5); decay/desertion resolution deferred to §5      |
| Health + illness flag| ✅   | 🟡 1–100 + sick flag decided; weekly update/wound math deferred (§5/combat)     |
| Inventory: items     | ✅   | 🟡 held items incl. gold `[1]`; per-item rules deferred to the items pass       |
| Inventory: men       | ✅   | 🟡 typed counts (§3.4); training/upkeep/combat deferred (§5/combat)             |
| Skills + experience  | ✅   | ❓ slot reserved; skill model deferred to the skills/magic pass                 |
| Aura (current/max)   | ✅   | 🟡 present on **every** noble, `0` for non-mages; spend/replenish & max-growth deferred to magic pass |
| Combat attitude, rank| 🟡   | ❓ attitude/behind/rank deferred to combat & provinces passes                   |
| Player-character flag| ✅   | ✅ marks the faction's first noble (§3.6); no other special behavior            |

- **Health is noble-only.** Men have no health rating (alive or dead); some NPCs read `n/a` and need
  a hit of ≥ 50 to be killed. The slot lives on nobles only. ✅
- **Aura is tracked on every noble** — current and maximum, defaulting to `0` and rising as a noble
  learns spells. Carrying it **universally** (not only on mages) keeps the noble type uniform and lets
  any noble take up magic without a shape change; the cost is one cheap integer pair per noble. ✅
  There is **no separate "piety" rating**: priesthood is simply knowing Religion `[750]` (temple
  offerings + prayers-as-skills, §skills pass). Piety belongs to Scott Turner's *Olympia: The Age of
  Gods*, **not** this rulebook. ✅
- **Stack grouping (model):** stacking is a **tree** — each noble may be stacked *under* exactly one
  parent noble, forming a stack whose top-most member is the **leader**. Only one level is shown in
  reports, but the engine stores the **full parent chain** because break-up follows it (a noble
  follows the parent it was stacked under). Stacking is **orthogonal to faction ownership and
  loyalty**: stack-mates may belong to different factions (gated by `ADMIT`). ✅ shape; movement/combat
  consequences are §5/combat.

### 3.4 The faction → noble → men/items hierarchy ✅

Three nesting levels, each with a sharply different status:

- **Faction** — owns nobles and the NP pool; takes no character orders, holds no possessions (§3.1).
- **Noble** — a first-class **entity** and the sole order-taker; holds items and men.
- **Men & items** — **possessions, not entities.** Men (peasant, worker, sailor, soldier, …) have
  **no entity number**, cannot learn skills, hold items, or act independently; they are modeled as
  **typed quantities** on their holder, not as units. Items (gold, weapons, scrolls, …) likewise. ✅
- **Holder vs. stack aggregate:** an item or man belongs to a *specific* noble (which matters for
  ownership, maintenance billing, and drops), but **carrying capacity and upkeep are computed for the
  stack as a whole** — one stack-mate may hold all the gold and pay the whole stack's upkeep; another
  may hold all the horses. The distribution across same-faction stack-mates is irrelevant to capacity.
  ✅ (the capacity/upkeep math itself is §5.)
- "Men" includes beast-fighters (e.g. dragons, via Beastmastery) but **not** work-animals (horses,
  oxen) that have no combat value — both are possessions, distinguished later by the combat/economy
  passes. 🟡

### 3.5 Loyalty bonds ✅ (model) / 🟡 (resolution)

- A noble carries **exactly one active loyalty bond to its faction**, of kind **contract**, **oath**,
  or **fear**, plus an integer **rating** (`contract-500`, `oath-2`, `fear-50`). Only one kind is
  active at a time. ✅ This is a per-noble attribute (the noble HONORs/oaths *itself*); the "lord" the
  bond names is the **owning faction**.
- **Decided starting values:** the player character begins **oath-2**; newly hired/FORMed nobles begin
  **contract-500**. ✅
- **Deferred to §5 (turn resolution):** monthly **decay** (contract `max(50, 10%)`; fear `1–2`; oath
  none), **desertion** at contract-0/fear-0 (50%/mo), and **bribe/oath defection** resistance
  (oath-1 ignores bribes; oath-2 immovable). The *values* are recorded here; the *when/how* is a
  resolution-phase concern. 🟡

### 3.6 Noble lifecycle: birth & death ✅ / 🟡

- **Birth.** A faction's first noble is the **player character (PC)**, present at game start at
  oath-2. Further nobles are created with **`FORM`**, which **spends Noble Points** (§3.7). Nothing is
  special about the PC beyond being first; if it dies, play continues with the faction's other nobles.
  ✅ shape; `FORM` cost/syntax is §4.
- **Death.** Health reaching **0** (or a killing blow) ends a noble. On death the noble **becomes a
  `Body` item** dropped into its province, recoverable with `EXPLORE` (an executioner receives the
  body directly). ✅ The body **decomposes 1.5 game years after death** — **12 turns**, since a year is
  8 months/turns (die turn 20 → decompose end of turn 32). ✅
- **NP return on dissolution.** Noble Points invested in a noble **return to the original owner** when
  its body **decomposes** (or on desertion — but a contract/fear renouncer's NPs are withheld until it
  next swears to a faction or dies). ✅ values; the return is applied during resolution (§5).
- **Resurrection / `LAY TO REST`** (priest skills that hasten or reverse a spirit's passing): **later
  content**, deferred to the skills/religion pass. ❓

### 3.7 Noble Points (NP) ✅ / 🟡

- NP is a **faction-level resource** — a single pool on the faction/player entity, **not** a per-noble
  balance. ✅ A noble *consumes* NP (FORM, advanced skills) or has NP *locked into* it (an oath bond);
  the locked NP is faction property held in escrow, returned on dissolution (§3.6).
- **Decided facts:** players start with a set amount; **late joiners get catch-up NP** so all players
  hold roughly equal NP; **+1 NP every turn that is a multiple of 8** (turns 8, 16, 24, …). NP buys
  nobles (`FORM`), some advanced skills, and oath loyalty. ✅
- The **starting amount**, **catch-up formula**, and **per-skill NP costs** are 🟡 — recorded as
  present, valued during the orders/skills passes.

### 3.8 Architectural implications

These follow from §3 and belong with §2.9 in AGENTS.md's "Open architectural decisions" table:

- **Faction & Noble are core `domain` aggregates.** They are pure types with invariants (one active
  loyalty bond; health in `[0,100]`; stack parent within the same location). An `OrderBundle` (§4)
  targets nobles **by entity number**, so ingest must resolve order → noble against the current
  snapshot — a lookup the domain exposes, infra never performs.
- **Deterministic entity-ID allocation.** Minting a new entity number at `FORM` (or when an item is
  created) must be a **pure function of recorded state** — e.g. a monotonic counter persisted in the
  per-turn snapshot, advanced inside resolution — mirroring the §2.9 turn-seed discipline. The domain
  imports **no** entropy or clock source. The numbering **scheme/alphabet** stays ❓ (§3.2).
- **Men-as-possessions keep the entity table small.** Modeling men as typed counts (not units) means
  the entity-number space holds only nobles, items, skills, and sub-locations — not the thousands of
  peasants a large game spawns. This is a deliberate model choice, not an optimization.
- **Bodies are items, not nobles.** Death is a **type transition** (noble → `Body` item) with a
  decomposition timer; the snapshot must carry dead-body items and their death turn so the 12-turn
  decay and NP return are deterministic.

> **Not yet distilled.** Like §2 before its Map reference, §3's decided facts are not yet promoted to a
> `reference/model/` page. Promote the noble/faction attribute model into a reference page once §4
> (orders) confirms the attributes orders actually read and mutate — drafting that page now would
> freeze slots the orders pass may still reshape.

## 4. Orders ❓

> **Early decision (recorded ahead of the full orders pass).**
>
> **Player-supplied names are untrusted and must be sanitized.** ✅ Players name nobles — and other
> createable entities — through orders (e.g. `NAME`/`FORM`). Because order files are the engine's
> **untrusted-input boundary** (`internal/infra/orderfile/`), a name carries whatever bytes a player
> typed and must be neutralized before it is ever embedded in a report (text, PDF, or any HTML view),
> where it could otherwise inject markup/script (XSS) or terminal/PDF control sequences. Only a safe,
> typed name reaches `domain`.
>
> **Mechanism ✅:** names are **sanitized at ingest**, in the orderfile adapter — consistent with
> CLAUDE.md's rule that `orderfile/` is *the* validation boundary, so only safe, typed names cross into
> `domain` and every render target inherits a clean string. Render-time escaping, if added, is
> defense-in-depth, not the primary control. The exact allowed character set and transform (reject vs.
> strip vs. escape on the way in) is 🟡 — settled with the rest of §4.

## 5. Turn resolution ❓

## 6. Turn reports ❓

## 7. Open decisions carried from AGENTS.md ❓
