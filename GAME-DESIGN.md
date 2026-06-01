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
buildable model. Spatial flavor (terrain *yields*, special realms) is deferred where noted. §2
is **geometry only**: a province's non-spatial attributes — tax base, ownership, buildings,
garrisons, rank — are political/economic and live in §5 (Provinces & territory) and §6 (Economy).

> **Distilled →** the decided (✅) facts of this section are published as the
> [Map reference](docs/content/reference/model/map.md) — the sole source the engine builds
> from. Promoted so far: representation (§2.1), coordinates (§2.2), terrain types (§2.3), the
> decided parts of movement (§2.4), inner locations & ports (§2.5), holes & hidden routes
> (§2.6), the civilization formula (§2.7), and the **authored seed data** (§2.1/§2.8 — region
> membership, initial-settlement locations, and safe-haven cities). **Not yet distilled** (revisit on
> each pass): the §2.4 variance model and `FLY` rules (❓), and §2.8's **richer region attributes &
> special realms** (deferred). When any of these are decided here, promote them into the Map reference
> before coding.

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
- Beyond geometry, the authored artifact also carries **seed data** the engine loads as immutable
  input but whose mechanics live elsewhere: **region membership** (§2.8), the **locations of the
  initial settlements** where new factions begin (game setup; cross-ref §3.6), and the list of
  **cities that are safe havens** (§2.8).

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
    (`aaa`=507) for very large maps.
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
  defense, sighting) are **deferred to §6/§8 (economy/combat)** — 🟡.

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
  Orders section (§10). ❓

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

### 2.8 Regions & special realms 🟡

- **Regions** are **named collections of provinces, authored as part of the map** (§2.1): the map
  author supplies the province→region mapping, and **every province belongs to exactly one region**.
  Membership is therefore a loaded, immutable map fact — enough for the §5 political rules that read it
  (garrison binding "same region", king-hood over a whole region). ✅ Whether a region carries **further
  attributes** beyond its name and membership is still deferred. ❓
- **Hades, Faery, the Cloudlands** are lore-specified with partial mechanics (Faery Hunt
  combat ratings, flight-only Cloudlands). Treated as **later content**, not part of the core
  map pass.
- **Safe havens** are **authored**: the map author supplies the list of **cities that are safe
  havens** (a designation on a city, feeding §2.7's civ contribution of 2). ✅ placement. Their count,
  the "no combat or magic" enforcement, and any special-realm behavior are deferred to the
  combat/realm passes. ❓

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
the passes that consume them — orders (§10), turn resolution (§11), and the items, skills,
and combat sections (§4, §7, §8).

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
  and must be sanitized before reaching any report — see §10.

### 3.3 Noble attributes ✅ (slots) / 🟡 (mechanics)

The decision here is the **attribute set** — which slots a noble carries. Whether each slot's
*mechanics* are settled is marked per row; an unsettled mechanic still reserves its slot now so §10/§11
have a stable target.

| Attribute            | Slot | Mechanics settled?                                                              |
| -------------------- | ---- | ------------------------------------------------------------------------------- |
| Entity number, name  | ✅   | ✅ identity (§3.2)                                                               |
| Location             | ✅   | ✅ province/sub-location placement (§2.5)                                        |
| Stack position       | ✅   | 🟡 grouping model below; movement/combat effects deferred to §8/§11             |
| Loyalty bond         | ✅   | 🟡 kind + rating decided (§3.5); decay/desertion resolution deferred to §11     |
| Health + illness flag| ✅   | 🟡 1–100 + sick flag decided; weekly update/wound math deferred (§8/§11)        |
| Inventory: items     | ✅   | 🟡 held items incl. gold `[1]`; per-item rules deferred to §4 (items)           |
| Inventory: men       | ✅   | 🟡 typed counts (§3.4); training/upkeep/combat deferred (§6/§8)                 |
| Skills + experience  | ✅   | ❓ slot reserved; skill model deferred to §7 (skills, magic & religion)         |
| Aura (current/max)   | ✅   | 🟡 present on **every** noble, `0` for non-mages; spend/replenish & max-growth deferred to §7 |
| Combat attitude, rank| 🟡   | ❓ attitude/behind deferred to §8 (combat); rank to §5 (territory)              |
| Player-character flag| ✅   | ✅ marks the faction's first noble (§3.6); no other special behavior            |

- **Health is noble-only.** Men have no health rating (alive or dead); some NPCs read `n/a` and need
  a hit of ≥ 50 to be killed. The slot lives on nobles only. ✅
- **Aura is tracked on every noble** — current and maximum, defaulting to `0` and rising as a noble
  learns spells. Carrying it **universally** (not only on mages) keeps the noble type uniform and lets
  any noble take up magic without a shape change; the cost is one cheap integer pair per noble. ✅
  There is **no separate "piety" rating**: priesthood is simply knowing Religion `[750]` (temple
  offerings + prayers-as-skills, §7). Piety belongs to Scott Turner's *Olympia: The Age of
  Gods*, **not** this rulebook. ✅
- **Stack grouping (model):** stacking is a **tree** — each noble may be stacked *under* exactly one
  parent noble, forming a stack whose top-most member is the **leader**. Only one level is shown in
  reports, but the engine stores the **full parent chain** because break-up follows it (a noble
  follows the parent it was stacked under). Stacking is **orthogonal to faction ownership and
  loyalty**: stack-mates may belong to different factions (gated by `ADMIT`). ✅ shape; movement/combat
  consequences are §8/§11.

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
  ✅ (the capacity/upkeep math itself is §6.)
- "Men" includes beast-fighters (e.g. dragons, via Beastmastery) but **not** work-animals (horses,
  oxen) that have no combat value — both are possessions, distinguished later by Economy/Combat
  (§6/§8). 🟡

### 3.5 Loyalty bonds ✅ (model) / 🟡 (resolution)

- A noble carries **exactly one active loyalty bond to its faction**, of kind **contract**, **oath**,
  or **fear**, plus an integer **rating** (`contract-500`, `oath-2`, `fear-50`). Only one kind is
  active at a time. ✅ This is a per-noble attribute (the noble HONORs/oaths *itself*); the "lord" the
  bond names is the **owning faction**.
- **Decided starting values:** the player character begins **oath-2**; newly hired/FORMed nobles begin
  **contract-500**. ✅
- **Deferred to §11 (turn resolution):** monthly **decay** (contract `max(50, 10%)`; fear `1–2`; oath
  none), **desertion** at contract-0/fear-0 (50%/mo), and **bribe/oath defection** resistance
  (oath-1 ignores bribes; oath-2 immovable). The *values* are recorded here; the *when/how* is a
  resolution-phase concern. 🟡

### 3.6 Noble lifecycle: birth & death ✅ / 🟡

- **Birth.** A faction's first noble is the **player character (PC)**, present at game start at
  oath-2. Further nobles are created with **`FORM`**, which **spends Noble Points** (§3.7). Nothing is
  special about the PC beyond being first; if it dies, play continues with the faction's other nobles.
  ✅ shape; `FORM` cost/syntax is §10.
- **Death.** Health reaching **0** (or a killing blow) ends a noble. On death the noble **becomes a
  `Body` item** dropped into its province, recoverable with `EXPLORE` (an executioner receives the
  body directly). ✅ The body **decomposes 1.5 game years after death** — **12 turns**, since a year is
  8 months/turns (die turn 20 → decompose end of turn 32). ✅
- **NP return on dissolution.** Noble Points invested in a noble **return to the original owner** when
  its body **decomposes** (or on desertion — but a contract/fear renouncer's NPs are withheld until it
  next swears to a faction or dies). ✅ values; the return is applied during resolution (§11).
- **Resurrection / `LAY TO REST`** (priest skills that hasten or reverse a spirit's passing): **later
  content**, deferred to §7 (skills, magic & religion). ❓

### 3.7 Noble Points (NP) ✅ / 🟡

- NP is a **faction-level resource** — a single pool on the faction/player entity, **not** a per-noble
  balance. ✅ A noble *consumes* NP (FORM, advanced skills) or has NP *locked into* it (an oath bond);
  the locked NP is faction property held in escrow, returned on dissolution (§3.6).
- **Decided facts:** players start with a set amount; **late joiners get catch-up NP** so all players
  hold roughly equal NP; **+1 NP every turn that is a multiple of 8** (turns 8, 16, 24, …). NP buys
  nobles (`FORM`), some advanced skills, and oath loyalty. ✅
- The **starting amount**, **catch-up formula**, and **per-skill NP costs** are 🟡 — recorded as
  present, valued during the Orders/Skills passes (§10/§7).

### 3.8 Architectural implications

These follow from §3 and belong with §2.9 in AGENTS.md's "Open architectural decisions" table:

- **Faction & Noble are core `domain` aggregates.** They are pure types with invariants (one active
  loyalty bond; health in `[0,100]`; stack parent within the same location). An `OrderBundle` (§10)
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
> `reference/model/` page. Promote the noble/faction attribute model into a reference page once §10
> (orders) confirms the attributes orders actually read and mutate — drafting that page now would
> freeze slots the orders pass may still reshape.

## 4. Items & possessions 🟡

The item entity model — what an item *is*, and the shared table of item types (gold `[1]`,
weapons, armor, scrolls, raw materials, and `Body` items). §3 reserved each noble's item and men
slots (§3.3–§3.4) and made a noble a `Body` item on death (§3.6); this section defines the items
themselves: weight/carry ratings, stacking of identical items, and creation via `MAKE`. Primary
sources: [logistics.md](docs/content/rules/logistics.md) (carrying capacity, making weapons &
armor), [tables.md](docs/content/rules/tables.md) (item table); trade-good aspects are shared
with §6.

### 4.1 What an item is: the fungible / unique split ✅

This resolves the apparent tension between §3.2 (which lists items in the entity-number namespace)
and §3.4 (which calls items *possessions, not entities*). Items take **two representations**:

- **Fungible items** — identified solely by a **type code** from the shared table (gold `[1]`,
  iron `[79]`, longsword `[74]`). A holder carries a **quantity** of a type; identical units are
  indistinguishable and **combine into one count** (three longbows are `+3` on the longbow row, not
  three entities). This is the **same mechanism** as typed-count men (§3.4) — men are simply rows in
  the same table that happen to fight.
- **Unique items** — each a distinct entity with its own **minted entity number** (a scroll
  `[yq12]`, a magical weapon, a relic, a `Body`). They carry **per-instance state** and never combine.

- **Reconciliation (decided):** the number on a fungible item (`gold [1]`) is its **type** code,
  shared by every holder — *not* a per-instance identity. So fungible items are "possessions modeled
  as typed quantities," exactly like men (§3.4); only **unique** items occupy the entity-number
  namespace as individual entities, subject to the deterministic minting of §3.8.
- A holder is always a **noble** (§3.4) — a faction holds no items. Carrying capacity and upkeep are
  computed for the **stack as a whole** (§3.4), so which stack-mate holds a given item is irrelevant
  to logistics.

### 4.2 The shared item-type table ✅ (shape) / 🟡 (numbering & contents)

- A single **authored table** keys every item *kind* by code and supplies its static ratings —
  weight, the three carry capacities, and (for men/beasts) combat values. This is the master table
  [tables.md](docs/content/rules/tables.md) calls "Weights"; **men (codes `[10]`–`[34]`) are rows in
  it** (§3.4), unifying men and items under one schema.
- Kinds present (illustrative, per tables.md): currency (gold `[1]`); men, beasts & mounts
  (`[10]`–`[34]`, `[51]`–`[55]`, `[76]`, `[271]`–`[295]`); siege engines (`[60]`–`[62]`); raw
  materials & trade goods (`[63]`–`[102]`, `[261]`); weapons & armor (`[72]`–`[75]`, `[85]`); blank
  scroll `[84]`; magic items (mithril `[82]`, gate crystal `[83]`, crystal orb `[290]`); relics
  (Imperial Throne `[401]`, Crown of Prosperity `[402]`).
- Like the map (§2.1), the table is a **fixed authored artifact** — immutable input to resolution,
  never mutated by it. The exact roster and the on-disk numbering scheme inherit §3.2's open question
  (decimal vs. base-N alphanumeric). 🟡

### 4.3 Weight & carrying capacity ✅ (ratings) / cross-ref §2.4, §6

- Every item and man carries a **weight** and three carry capacities — **land/walk, ride, fly** —
  read from the table. Capacity **excludes the item's own weight**; a `-1` capacity means "carries its
  own weight but nothing more" (wild horses, oxen, most beasts).
- A stack **rides** only if total ride capacity covers the weight of all non-riders, else it
  **walks**; it **flies** only if fly capacity covers all non-flyers. Land travel auto-selects the
  fastest available mode (§2.4).
- **Overload is deterministic:** ≥150% of walking capacity ⇒ +50% travel time; >200% ⇒ cannot move.
  ✅ The interpolation between 100% and 200% is 🟡 — designed with §2.4's variance model.
- Weight is the universal logistics unit: it drives carry capacity here and ferry fees (gold per 100
  weight, §9) — and nothing else. **Upkeep is paid in gold and applies to men only** (§6), never by
  item weight.

### 4.4 Gold `[1]` ✅

- Gold is the currency: **type `[1]`, weight 0**, fungible (§4.1). It is the medium for upkeep (§6),
  `HONOR`/bribes (§3.5), market trades (§6), and ferry fees (§9).
- **Gold left loose in a province does not carry across the turn boundary** — unheld gold is not
  banked; gold held by a noble persists normally. (A resolution detail flagged here for §11.)

### 4.5 Unique items & per-instance state 🟡

Unique items each carry state beyond their type code:

- **Scrolls / books** — the **skill they teach**; a noble may `STUDY` from one (§7). A blank scroll
  `[84]` becomes a unique taught-skill scroll when scribed (Record skill on scroll `[692]`).
- **Magical weapons / armor / bows, auraculum, palantir** — bonuses plus a **creator** identity
  (Artifact construction `[880]` has reveal/cloak-creator spells). Mechanics deferred to §7.
- **Relics** — unique quest artifacts with bespoke effects (the Crown of Prosperity raises its
  province's civ by +2 each turn it ends there) and a **return timer** (the Crown returns 12–24 turns
  after appearing; the Imperial Throne never returns). The randomized return window is a §11
  resolution concern; the relic's per-instance carry-state is fixed here. 🟡
- **Body** — see §4.6.

### 4.6 Bodies as items ✅ (cross-ref §3.6)

- A noble's death is a **type transition** to a unique `Body` item dropped in its province (§3.6),
  recoverable with `EXPLORE` (an executioner receives it directly). The body carries the **dead
  noble's identity, death turn, and invested NP** so the engine applies the **12-turn decomposition**
  and **NP return** deterministically (§3.6, §3.8). Resurrection / `LAY TO REST` reads the body —
  deferred to §7.

### 4.7 Creation, transfer & consumption 🟡

- Items enter play only through **skill-driven production**, never randomly: `MAKE` (weapons/armor
  from raw materials — one input unit → one item per day, requires Weaponsmithing `[617]`), gathering
  (`COLLECT`/`HARVEST`/`MINE`/`QUARRY`/`FISH`/`CATCH`), `BREED` (beasts, §3.4), scroll scribing,
  potion brewing (Alchemy), artifact forging (`[880]`), and turn-lead-into-gold (`[697]`). §4 fixes
  the shape — production **consumes typed inputs and yields typed outputs**; per-skill rates and
  requirements are §6/§7.
- **Transfer:** `GIVE` moves items between co-located nobles anywhere (unlike market `BUY`/`SELL`,
  which match only in a city — §6); its low order-priority lets a same-day `GIVE` settle before
  `MOVE` (§10).
- **Consumption / loss:** `DROP` releases items (and men) from a noble; spell-required and
  training/`MAKE` input items are **consumed on use**; gold is spent. The fungible-count model makes
  each of these simple arithmetic on a holder's row.
- Any creation/destruction that **mints or frees a unique entity number** must be a **pure function of
  recorded state** (§3.8) — no `rand`/`time`.

### 4.8 Trade goods — shared with §6 🟡

- A **tradegood** is an ordinary item with a **market role**: found for sale or in demand via the
  Trade `[730]` sub-skills and moved between city markets for profit (§6). It is structurally nothing
  new — no distinct entity, no extra slot — so §4 records only that "tradegood" is a *role over the
  item table*, leaving matching, pricing, and the `BUY`/`SELL` economy to §6.

### 4.9 Architectural implications

These follow from §4 and join §2.9 / §3.8 in AGENTS.md's "Open architectural decisions" table:

- **The item-type table is authored reference data**, loaded immutably like the map (§2.1) — the same
  artifact-and-loader concern as the `MapSource` port, or a sibling to it. The domain holds it as a
  static lookup that resolution reads but never mutates.
- **Fungible items are typed counts, not entities.** With men-as-counts (§3.8) this keeps the
  entity-number table small — only nobles, **unique** items, skills, and sub-locations are minted. The
  men/item distinction is a *combat/identity* split over one shared schema, not two data models.
- **Unique-item minting reuses the §3.8 discipline** — a scribed scroll or forged artifact advances
  the same deterministic counter the domain uses for `FORM`.
- **Bodies and relics carry timers in the snapshot.** Decomposition (12 turns) and relic return
  (12–24 turns) must be reconstructible from recorded death/appearance turns; any randomized window
  derives from the §2.9 turn seed, never live entropy.

> **Not yet distilled.** Like §3, §4's decided facts wait on the orders pass (§10) before promotion to
> a `reference/model/` page — the item slots that orders read and mutate (`MAKE`, `GIVE`, `DROP`,
> `STUDY`) may still reshape per-instance state.

## 5. Provinces & territory control 🟡

The province as a **political/economic entity**, layered on §2's spatial graph: tax base,
ownership and control, the buildings it holds, castles and the garrisons bound to them, noble
**rank** (lord → king, by provinces controlled), pledge chains and shared **rulers**, decrees, and
**relics**. §2 stays geometry; this section owns everything non-spatial about a province. Primary
sources: [provinces.md](docs/content/rules/provinces.md),
[buildings-economy.md](docs/content/rules/buildings-economy.md). Cross-refs: §2.7 (civ level), §6.

### 5.1 The province as a political entity ✅ (slots)

§2 fixed a province's **geometry** (coordinates, terrain, routes); §5 adds the **mutable political /
economic facet** layered on top. The attribute set a province carries beyond §2:

- **Civilization level** (§2.7) and the **tax base** derived from it (§5.2);
- the **buildings** it holds — at most one **castle** and one **mine**, plus cities, inns, towers,
  temples (§5.3);
- at most one **garrison** (§5.5);
- the **controlling castle** and the **pledge-chain rulers** that share control (§5.4, §5.6);
- **depression state** — pillage and opium each lower *future* tax base on their own timers (§5.2).

This is the *slot set*; per-row mechanics are settled in the subsections below or deferred to §6
(economy flow) and §11 (resolution timing). Geometry stays immutable input (§2.1); the political
state is what resolution mutates.

### 5.2 Tax base ✅ (value) / cross-ref §6 (flow)

- A province's monthly **tax base** is a function of its civ level: **`50 + 50·civ`** gold
  (wilderness `50`, civ-1 `100`, … civ-7 `400`). A **city adds a flat +100**, folded into the
  province tax base **at end of turn** (diminished if the city was pillaged that month).
- **Tax base does not accumulate** — gold uncollected at turn's end is lost (consistent with §4.4's
  loose-gold rule).
- **Pillaging** seizes the whole tax base and lowers *future* revenue: each pillaging costs **4 months**
  of recovery (five consecutive months ⇒ 20 months to recover). **Opium** consumption in a market
  likewise depresses the province's tax base, scaling with volume.
- §5 owns the **value** (the civ→gold mapping, city bonus, depression). The **collection &
  distribution flow** — garrison maintenance first, then the castle's half — is §6 / §11.

### 5.3 Buildings ✅ (catalog & placement) / cross-ref §6, §2.7

- Five buildings exist: **castle, tower, inn, temple, mine**. Each contributes to civ level (§2.7) and
  has a role — castle (taxation & ownership, §5.4), tower (`RESEARCH` lab; up to **six per castle**),
  inn (monthly income + a sheltered noble's weekly illness resistance), temple (Religion study +
  offerings, §7), mine (resource extraction, §6).
- **Placement & cardinality** (decided): **one castle per province** (built in the outer province or
  its city, never another sub-location); **one mine per mountain province or rocky hill**; buildings
  may **not** nest inside buildings, **except** up to six towers inside a castle. Cities are
  authored sub-locations (§2.5), not built.
- **Ownership of a building** is positional: the **first character inside** owns it; leaving forfeits
  ownership to the next, and `ADMIT` gates entry (default: refuse other factions). Same rule governs
  ships (§9). ✅
- **Construction mechanics** (worker-days of effort, materials, the `BUILD` order, mine depth &
  collapse) are §6; this subsection fixes only **which buildings exist, where they may sit, and who
  owns them**.

### 5.4 Castles & land ownership ✅

- A **castle is the foundation of land ownership**: it auto-collects all gold in its province, and its
  owner receives **half of the remaining tax base** each month (after any garrison's maintenance is
  paid). One castle per province (§5.3).
- A castle **alone does not rule** — a **garrison must be stationed outside it** in the province to
  hold it against pillaging (§5.5).
- **Castle improvement** runs levels **1–6** (the `IMPROVE` order; each level costs stone +
  worker-days per the castle-improvement table — amounts are §6). Improvement level drives three
  things: the **civ contribution** `1.5 + level/4` (§2.7), the **garrison capacity / rank band** it can
  anchor (§5.6), and **protection** (a castle shelters the first **500** men in combat — §8).

### 5.5 Garrisons ✅

- A **garrison** is installed with **`GARRISON CASTLE`**, requires **≥10 soldiers**, and is **bound to a
  castle in the same region**. **One garrison per province** — so it is referable by the bare keyword
  `GARRISON` without its entity number.
- **Contiguous spread:** a garrison may be placed only in a province **adjoining one already garrisoned
  to the same castle** (or the castle's own province). A castle's garrison network therefore grows as a
  connected blob within its region.
- **Tax role:** a garrison pays its men's maintenance **from the province tax base**, then forwards
  **½ of the remainder** to its castle. A garrison that falls **below 10 fighting men** pays its own
  upkeep but **cannot** forward tax, guard against pillaging, or obey decrees (§5.7).
- **Limited reporting:** garrisons report only resource-depletion activity and **large/unusual parties**
  (any stack of ≥5 units, any party of ≥20 men, most monsters) — not full location reports, and never
  activity in hidden locations (§2.6).
- **Entity model ✅:** a garrison is **not a noble**. It is a **distinct station entity** in the
  entity-number space (§3.2): a unit with an entity number (`Garrison [780]`) holding typed soldiers,
  "on guard," bound to a castle, that takes **no player orders** — it acts only on its rulers' decrees
  (§5.7). Modeling it as its own kind keeps the noble type uniform (no order-taking exceptions) and
  spares the garrison every noble-only attribute (loyalty bonds, aura, noble health, NP).

### 5.6 Rank, rulers & pledge chains ✅ (bands) / reconciled

- A noble's **rank** is a function of the **number of provinces controlled**:

  | provinces | rank     |
  | --------- | -------- |
  | 1–5       | lord     |
  | 6–12      | knight   |
  | 13–24     | baron    |
  | 25–37     | count    |
  | 38–50     | earl     |
  | 51–63     | marquess |
  | 64+       | duke     |
  | whole region (≥15 provinces) | king |

- **King** requires controlling **every** province of a region that has **≥15 provinces**.
- **Castle improvement gates territorial reach:** the same bands cap how many provinces a single castle
  may anchor garrisons for — improvement level 0 ⇒ up to 5 (lord) … level 6 ⇒ 64+ (duke). Attaining a
  rank thus needs **both** the province count **and** a castle improved enough to hold them.
- **`PLEDGE`** lets a noble swear lands to another, granting the target **status** (the target's
  controlled-province count grows by the pledged provinces) and **shared control**. The pledger's own
  rank becomes **`min(original rank, one rank below the target)`**. **Income is unaffected** — it still
  flows to the castles; the target gains no extra gold.
- **Shared rulers:** every noble in a pledge chain **shares control** of the garrisoned provinces
  (rename province/sub-locations, take items from garrisons, issue decrees, receive garrison reports),
  so a province may have **many rulers**; visitors see the **top-most** ruler. A castle keeps receiving
  its garrison income even when its owner is pledged away. ✅

### 5.7 Decrees ✅

- A **ruler** (anyone in a province's pledge chain) issues **decrees** that all **functional garrisons**
  (≥10 men) in controlled provinces obey: **`DECREE WATCH WHO`** (surface a named unit in garrison
  reports — useful for spotting otherwise-unnoticed travelers) and **attack-on-sight / hostile** decrees
  against specified units.

### 5.8 Relics & province effects ✅ / cross-ref §4.5, §7

§4.5 fixed relics as **unique items** with per-instance state and return timers; §5 records their
**territorial effects**:

- **Crown of Prosperity `[402]`** — each turn it ends in a province, that province gains a **+2 civ
  level** modifier (a per-turn province effect). Returns 12–24 turns after appearing.
- **Imperial Throne `[401]`** — the **Emperor of Olympia** title capstone: whoever **rebuilds the
  castle on Mt. Olympus and sits on the throne** is titled Emperor. The Throne **never returns** to the
  netherworld (unlike all other relics).
- **Skull of Bastrestric `[403]`** — an aura-burst item (`USE 403`: +50–75 current aura, capped at 5×
  max; **25% chance it kills the mage**, instantly kills non-mages); vanishes on use or in 10–20 turns.
  An **aura/combat** mechanic — deferred to §7.
- Randomized return windows derive from the **§2.9 turn seed**, never live entropy (§4.9).

### 5.9 Architectural implications

These follow from §5 and join §2.9 / §3.8 / §4.9 in AGENTS.md's "Open architectural decisions" table:

- **Province gains a mutable political state** distinct from its immutable §2 geometry: tax base,
  buildings, garrison, controlling castle, ruler chain, civ level, and depression timers — all carried
  in the per-turn snapshot and rewritten by resolution. The authored map (§2.1) seeds the initial
  buildings and civ; geometry never changes.
- **Control and rank are derived, not stored.** A pledge chain is a forest of edges over nobles;
  controlled-province sets, rank, rulers, and king-hood are computed each turn as a **pure function** of
  castle ownership + garrison bindings + pledge edges — consistent with §3.1 ("territory a faction
  controls is derived").
- **Garrisons are a distinct station entity**, not a noble variant (§5.5) — a separate kind in the
  entity-number space (§3.2) that carries soldiers but none of a noble's order-taking, loyalty, aura,
  or NP machinery.
- **Region membership is authored map data.** Every province is assigned to **exactly one** region by
  the map author (§2.8), loaded immutably with the map (§2.1). Garrison binding ("same region") and
  king-hood ("every province in a region, ≥15 provinces") read this membership set directly — no
  separate region-ownership state is stored. Richer region attributes stay §2.8-deferred.
- **Depression & timer state in the snapshot.** Pillage recovery (4 months each), opium demand, and mine
  collapse (vanishes after 8 months) join §4's decomposition/return timers as recorded, deterministic
  countdowns.

> **Not yet distilled.** Like §3 and §4, §5's decided facts (the rank bands, tax-base table, building
> catalog) wait on the orders pass (§10) before promotion to a `reference/model/` page — the orders that
> read and mutate them (`GARRISON`, `PLEDGE`, `DECREE`, `BUILD`, `IMPROVE`, `PILLAGE`) may still reshape
> the slots.

## 6. Economy ❓

The monthly flow of money and materials: tax base → income, building **construction** (effort in
worker-days), **markets** (`BUY`/`SELL` matched in a shared city), **trade goods**, the
**maintenance/upkeep** of men, and **training** peasants into other kinds of men. Primary sources:
[buildings-economy.md](docs/content/rules/buildings-economy.md),
[markets.md](docs/content/rules/markets.md), [provinces.md](docs/content/rules/provinces.md) (tax
base), [logistics.md](docs/content/rules/logistics.md) (upkeep, training). Ships' economic role
cross-refs §9.

## 7. Skills, magic & religion ❓

The skill model: the category/sub-skill tree, `STUDY`/`RESEARCH`, **experience** levels
(apprentice → grand master), the six **schools of magic** and the **aura** they spend (§3.3
reserved the skills and aura slots), and **religion** — priests, prayers, and temple offerings,
with **no separate piety rating** (§3.3). Primary sources:
[skills-magic.md](docs/content/rules/skills-magic.md),
[tables.md](docs/content/rules/tables.md).

## 8. Combat ❓

Battle resolution: `ATTACK`/`DEFEND`, the combat **attitudes** (`HOSTILE`/`DEFEND`/`NEUTRAL`/
`DEFAULT`), `BEHIND` positioning, the **break point**, stack-leader targeting, garrisons and
**sieges**, **prisoners**, and the **wound** generation that feeds noble health (§3.3, §3.6).
Primary source: [combat.md](docs/content/rules/combat.md). Cross-refs: §5 (garrisons), §9 (no
siege engines at sea).

## 9. Ships ❓

Ships as combined movement + economic entities: galley and roundship, sailing requirements and
cargo capacity, **ferries** (`FEE`/`BOARD`/`FERRY`/`UNLOAD`), and docking (cross-ref §2.5). A ship
is an ordinary entity in the entity-number space (§3.2), not a sub-location. Primary source:
[ships.md](docs/content/rules/ships.md). May fold into §6 if it stays thin.

## 10. Orders ❓

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
> strip vs. escape on the way in) is 🟡 — settled with the rest of §10.

## 11. Turn resolution ❓

## 12. Turn reports ❓

## 13. Open decisions carried from AGENTS.md ❓
