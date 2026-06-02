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
  `MapSource` port; the on-disk **format is undecided** (see [`docs/adr/`](docs/adr/README.md),
  the "Map artifact format" register row).
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
  travel-time decision, this variance is **deterministic given the recorded RNG state in the
  game-state snapshot** (drawn through the `RNG` port, §2.9/§11.7). ✅ *mechanism* — the concrete **variance model**
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

The architectural consequences of §2 have moved to their correct homes per the routing rule in
[AGENTS.md](AGENTS.md); this section remains only as a pointer (other sections link to its anchor):

- **Map artifact format + `MapSource` port** → [`docs/adr/`](docs/adr/README.md): the "Map
  artifact format" register row (on-disk format still open) and the planned `MapSource` port
  noted in AGENTS.md's Ports section. The descriptive model is in
  [reference/model/map.md](docs/content/reference/model/map.md).
- **Randomness source** (decided) → [`docs/adr/`](docs/adr/README.md) ADR 0003: the `RNG` port
  (`app/ports.go`) realized by `internal/infra/prng`, with RNG state round-tripping through
  `GameStateStore`. Cross-ref §11.7/§11.9.

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

The architectural consequences of §3 have moved to their correct homes per the routing rule in
[AGENTS.md](AGENTS.md); this section remains only as a pointer (§3.2, §4, §5, and §6 link to its
anchor):

- **Faction & Noble as core `domain` aggregates** — pure types with invariants (one active loyalty
  bond; health in `[0,100]`; stack parent within the same location), and the order → noble lookup
  ingest resolves **by entity number against the current snapshot** — a domain-exposed lookup infra
  never performs. This is the standing SOUSA/domain discipline in [AGENTS.md](AGENTS.md) (untrusted
  input stops at the `orderfile` boundary; only typed `OrderBundle` values reach `app`). The
  descriptive attribute model itself awaits its `reference/model/` page (see the deferral note below).
- **Deterministic entity-number allocation** → [`docs/adr/`](docs/adr/README.md): minting a number at
  `FORM`/item creation is a pure function of recorded state — a **monotonic counter persisted in the
  per-turn snapshot**, advanced inside resolution (now pinned in the State-storage snapshot
  constraints, alongside the §2 turn-seed discipline). The determinism rule itself is the standing one
  in [AGENTS.md](AGENTS.md) (domain imports no entropy or clock source). The on-disk/display numbering
  **scheme/alphabet** stays open (§3.2) — a reference-side representation fact to settle with the
  model, like the province display code in
  [reference/model/map.md](docs/content/reference/model/map.md), not an architecture decision.
- **Men-as-possessions** (typed counts, not entities) keep the entity-number space to nobles, items,
  skills, and sub-locations — not the thousands of peasants a large game spawns; a deliberate model
  choice, not an optimization. The descriptive fact belongs in `reference/model/` and its rationale in
  `explanation/`; both ride with the deferred model page below.
- **Bodies are items, not nobles** → [`docs/adr/`](docs/adr/README.md): death is a **type transition**
  (noble → `Body` item) on a decomposition timer, so the snapshot must carry dead-body items and their
  death turn for the 12-turn decay and NP return to be deterministic (now pinned in the State-storage
  snapshot constraints). The type-transition fact itself is descriptive model, awaiting `reference/model/`.

> **Not yet distilled.** Like §2 before its Map reference, §3's decided facts are not yet promoted to a
> `reference/model/` page. Promote the noble/faction attribute model — and its men-as-possessions
> rationale into `explanation/` — once §10 (orders) confirms the attributes orders actually read and
> mutate; drafting them now would freeze slots the orders pass may still reshape.

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

The architectural consequences of §4 have moved to their correct homes per the routing rule in
[AGENTS.md](AGENTS.md); this section remains only as a pointer (§13.1 links to its anchor):

- **The item-type table as authored reference data** → the descriptive fact (an immutable static
  lookup resolution reads but never mutates, loaded like the map of §2.1) belongs in `reference/model/`,
  awaiting the item model page (see the deferral note below). Its on-disk format and loader are the
  **same open artifact-and-loader concern** as the "Map artifact format" register row and the planned
  `MapSource` port in [`docs/adr/`](docs/adr/README.md) — a sibling to it, not yet surfaced as a
  separate decision (§13.7).
- **Fungible items are typed counts, not entities** — with men-as-counts (§3.8) this keeps the
  entity-number space to nobles, **unique** items, skills, and sub-locations; the men/item distinction
  is a *combat/identity* split over one shared schema, not two data models. The descriptive fact belongs
  in `reference/model/` and its rationale in `explanation/`; both ride with the deferred model page below,
  as men-as-possessions does (§3.8).
- **Unique-item minting reuses the §3.8 discipline** → [`docs/adr/`](docs/adr/README.md): a scribed
  scroll or forged artifact advances the same deterministic entity-number counter the domain uses for
  `FORM`, already pinned in the State-storage snapshot constraints (numbers minted at `FORM`/item
  creation are a pure function of recorded state).
- **Bodies and relics carry timers in the snapshot** → [`docs/adr/`](docs/adr/README.md): decomposition
  (12 turns) and relic return (12–24 turns) must be reconstructible from recorded death/appearance turns,
  with any randomized window derived from the turn seed (§2.9), never live entropy. Already pinned in the
  State-storage snapshot constraints — dead-body items with their death turn explicitly, relic-return
  timers under the "all timer/countdown state" the snapshot must round-trip.

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

The architectural consequences of §5 have moved to their correct homes per the routing rule in
[AGENTS.md](AGENTS.md); this section remains only as a pointer (§6.8, §7.9, and §13.1 link to its
anchor):

- **Province carries a mutable political state** distinct from its immutable §2 geometry — tax base,
  buildings, garrison, controlling castle, ruler chain, civ level, depression timers, all in the
  per-turn snapshot and rewritten by resolution, with the authored map (§2.1) seeding initial
  buildings and civ. This is descriptive model awaiting the deferred `reference/model/` page (see the
  note below); its snapshot-carried timers are already pinned under "all timer/countdown state" in
  the [`docs/adr/`](docs/adr/README.md) State-storage constraints.
- **Control and rank are derived, not stored** — controlled-province sets, rank, rulers, and king-hood
  are a **pure function** of castle ownership + garrison bindings + pledge edges, recomputed each turn,
  consistent with the §3.8 "territory a faction controls is derived" fact. Descriptive model, awaiting
  the same deferred `reference/model/` page.
- **Garrisons are a distinct station entity**, not a noble variant (§5.5) — a separate kind in the
  entity-number space (§3.2) carrying soldiers but none of a noble's order-taking, loyalty, aura, or NP
  machinery. Descriptive model, stated in the §5.5 body, awaiting the deferred `reference/model/` page.
- **Region membership is authored map data** → already distilled in
  [reference/model/map.md](docs/content/reference/model/map.md): every province is assigned to exactly
  one region, loaded immutably with the map; garrison binding ("same region") and king-hood read this
  membership set directly, with no separate region-ownership state stored.
- **Depression & timer state in the snapshot** → [`docs/adr/`](docs/adr/README.md): pillage recovery (4
  months each), opium demand, and mine collapse (8 months) are recorded, deterministic countdowns —
  already pinned under the "all timer/countdown state" the snapshot must round-trip in the State-storage
  constraints (§13.1 names them explicitly, alongside §4's decomposition/return timers).

> **Not yet distilled.** Like §3 and §4, §5's decided facts (the rank bands, tax-base table, building
> catalog) wait on the orders pass (§10) before promotion to a `reference/model/` page — the orders that
> read and mutate them (`GARRISON`, `PLEDGE`, `DECREE`, `BUILD`, `IMPROVE`, `PILLAGE`) may still reshape
> the slots.

## 6. Economy 🟡

The monthly flow of money and materials, layered on §5's tax base and §3/§4's possessions. This
section closes the deferrals routed here from elsewhere: upkeep math (§3.4, §4.3), production rates
(§4.7), the tax **collection & distribution flow** (§5.2, §5.4, §5.5), and the recruit/train
mechanics (§3.4). It owns **how gold and materials move each month** — income collection and payout,
paying and training men, constructing buildings and improving castles, extracting and making
resources, and clearing city markets. Most *values* are fixed by the rulebook tables and decided
here; *when* each step runs within a turn is a §11 (turn resolution) concern, marked where it bites.
Primary sources: [buildings-economy.md](docs/content/rules/buildings-economy.md),
[markets.md](docs/content/rules/markets.md), [provinces.md](docs/content/rules/provinces.md) (tax
base), [logistics.md](docs/content/rules/logistics.md) (upkeep, training),
[tables.md](docs/content/rules/tables.md). Cross-refs: §5 (tax-base value, garrisons, castles), §7
(skill mechanics & per-use yields), §9 (ships' economic role), §11 (resolution timing).

### 6.1 The monthly money flow ✅ (split, un-garrisoned castle) / 🟡 (timing)

§5.2 fixed the tax-base **value** (`50 + 50·civ`, +100 per city) and deferred the **flow** here. The
end-to-end path for a **garrisoned province** bound to a castle:

1. the province generates its tax base (§5.2);
2. the garrison pays its own men's **maintenance** from the tax base (§6.2);
3. the castle owner receives **⌊(tax base − garrison maintenance) / 2⌋**;
4. the **un-forwarded half stays in the province and is lost** — gold does not accumulate (§4.4, §5.2).

Worked example (from [provinces.md](docs/content/rules/provinces.md)): a civ-5 province (base 300)
with a 10-soldier garrison (upkeep 20) leaves 280; the castle owner receives 140; the other 140 is
lost.

- **Reconciliation (decided ✅):** [provinces.md](docs/content/rules/provinces.md) says both that "a
  castle automatically collects **all** of the gold in its province" and that "the owner receives
  **half** of the remaining tax base." These describe **two different things, not a conflict**:
  "automatically … all" means collection is **passive** — the castle sweeps the province's gold with
  **no order required** (there is no `COLLECT` command) — while "half" governs **how much the owner
  keeps**. So the castle collects automatically, and the owner receives half of `(base − garrison
  upkeep)`; the un-kept half is lost, consistent with the no-accumulation rule.
- **Pillaging** (§5.2) **seizes the whole tax base** for the month (a pillager must first defeat any
  guarding garrison, §8) and lowers *future* revenue (4 months' recovery per pillaging).
- **Un-garrisoned castle (settled ✅ in §11.6):** a castle with **no functional garrison** (≥10 men)
  in its own province is **undefended, and its tax goes uncollected** that month — collection is gated
  on the same garrison presence that gates protection. (Previously pinned to §11 for confirmation;
  resolved there.)
- **Timing (🟡):** the order of operations within a turn — when tax is computed, when upkeep debits,
  when the owner is paid — is a §11 concern and matters for idempotency (the `TurnLedger`).

### 6.2 Maintenance / upkeep of men ✅

This closes the upkeep deferral from §3.4 and §4.3. Men cost gold **per month**, charged to the
**holding noble** at end of month:

| man                         | gold/mo |
| --------------------------- | ------- |
| peasant `[10]`              | 1       |
| worker `[11]`, sailor `[19]`, crossbowman `[21]`, soldier `[12]` | 2 |
| archer `[13]`, pikeman `[16]`, blessed soldier `[17]`, swordsman `[20]`, pirate `[24]` | 3 |
| knight `[14]`, elite archer `[22]` | 4 |
| elite guard `[15]`          | 5       |

- **Stack-aggregate billing (§3.4):** if the holding noble lacks gold, it draws from **same-faction
  stack-mates**; one stack-mate carrying gold can pay the whole stack's upkeep. Nobles never share
  gold across factions.
- **Partial payment:** if a noble can pay only some of its men, **one-third of those unpaid leave**
  service at month's end. **Peasants do not desert but starve** if unpaid. Which men leave/starve is
  the engine's choice — and so must be a **deterministic** selection from recorded state, never
  `rand` (§6.8).
- `DROP` releases men deliberately (§4.7). **Upkeep is gold-per-man only** — item *weight* drives
  carry capacity and ferry fees (§4.3, §9) but never upkeep.

### 6.3 Recruiting & training men ✅ (table) / ❓ (recruit supply)

- **`RECRUIT`** draws **peasants `[10]`** from the surrounding province. It is a production command:
  it **fails immediately (zero time)** where there are no peasants. The **province peasant supply
  model** and any gold cost are **unspecified ❓** — recorded as a gap to settle with §11 (does
  recruiting deplete a province pool, scale with civ level, cost gold?).
- **`TRAIN`** converts one kind of man into another at **one day per man**, consuming an **input man**
  (and, for some, an **input item**), gated by a **skill** and sometimes a **location**:

  | output                  | skill  | input man    | input item        | where  |
  | ----------------------- | ------ | ------------ | ----------------- | ------ |
  | worker `[11]`           | none   | peasant      | —                 | —      |
  | soldier `[12]`          | 610    | peasant      | —                 | —      |
  | sailor `[19]`           | 601    | peasant      | —                 | —      |
  | crossbowman `[21]`      | 610    | peasant      | crossbow `[85]`   | —      |
  | archer `[13]`           | 615    | soldier      | longbow `[72]`    | —      |
  | pikeman `[16]`          | 610    | soldier      | pike `[75]`       | —      |
  | swordsman `[20]`        | 616    | soldier      | longsword `[74]`  | —      |
  | blessed soldier `[17]`  | 750    | soldier      | —                 | temple |
  | knight `[14]`           | 616    | swordsman    | warmount `[53]`   | —      |
  | elite archer `[22]`     | 615    | archer       | —                 | castle |
  | elite guard `[15]`      | 616    | knight       | plate armor `[73]`| castle |
  | pirate `[24]`           | 616    | sailor       | longsword `[74]`  | ship   |

  The peasant→soldier→swordsman→knight→elite-guard spine, plus the worker / sailor / crossbowman /
  archer / pikeman / blessed-soldier branches, matches the training tree in
  [logistics.md](docs/content/rules/logistics.md). The skill numbers are recorded here; the **skill
  mechanics** (experience gating, what each subskill of Combat/Shipcraft/Religion permits) are §7.

### 6.4 Construction & castle improvement ✅

Closes the construction-mechanics deferral from §5.3/§5.4. Building requires **Construction `[680]`**
(a mine may also use **Mining `[720]`**), at least **three workers**, and the materials below; the
builder issues a `BUILD` order at the outer level of the province (unstacked):

| building | effort (worker-days) | material   | where                      |
| -------- | -------------------- | ---------- | -------------------------- |
| inn      | 300                  | 75 wood    | province or city           |
| mine     | 500                  | 25 wood    | mountain or rocky hill     |
| temple   | 1,000                | 50 stone   | anywhere                   |
| tower    | 2,000                | 100 stone  | anywhere                   |
| castle   | 10,000               | 500 stone  | province or city (§5.3)    |

- **Materials are staged in fifths:** the builder must hold ≥1/5 of the materials to start (deducted
  immediately); the next fifth is due at 20% complete, the next at 40%, etc. Construction **halts**
  if materials run out. The building completes when the required worker-days are invested; the
  builder and workers are placed **inside** the new structure. Resume a partial build by **entering**
  it and re-issuing the `BUILD` order.
- **Placement & cardinality** are §5.3 (one castle per province; one mine per mountain/rocky-hill;
  no nesting except up to six towers in a castle). **Ownership** is positional — first character
  inside owns it, gated by `ADMIT` (§5.3).
- **Castle improvement** (`IMPROVE [days]`, runs to the next level or the day budget) climbs **levels
  1–6**, each costing stone + worker-days:

  | level | stone | worker-days |
  | ----- | ----- | ----------- |
  | 1     | 50    | 1,000       |
  | 2     | 60    | 1,250       |
  | 3     | 70    | 1,500       |
  | 4     | 80    | 1,750       |
  | 5     | 90    | 2,000       |
  | 6     | 100   | 2,500       |

  Improvement level drives the **civ contribution** `1.5 + level/4` (§2.7), the **rank/garrison reach
  band** (§5.6), and **combat protection** (a castle shelters the first 500 men — §8).

### 6.5 Resource extraction & production ✅ (shape) / 🟡 (yields, terrain)

All items enter play through **skill-driven production** (§4.7), never randomly — each command
**consumes typed inputs and yields typed outputs**, taking time:

- **Gathering**, gated by terrain and skill: Forestry `[700]` (harvest lumber `[702]`, yew `[703]`,
  mallorn `[705]`, rare foliage `[704]`, opium `[706]`); Mining `[720]` (iron `[721]`, gold `[722]`,
  mithril `[723]`); Fishing `[603]`; Stone quarrying `[682]`; Collect rare elements `[695]`. **Per-use
  yields are unspecified ❓** and deferred to §7; this section records only that gathering produces
  typed item rows. The **terrain→which-gathering-works-where** linkage (opium in swamp, mines in
  mountain, lumber in forest) partially closes §2.3's terrain-yield deferral; a fuller terrain-yield
  model stays 🟡.
- **Mines** (§5.3): a new mine starts at **depth 1**; it deepens **one level per three extraction
  uses**. The **resource mix shifts with depth** — iron near the surface, gold deeper, rare elements
  deepest. Deeper mines suffer **cave-ins more often**, raising a **damage percentage**; `REPAIR`
  arrests it, and an unrepaired mine eventually **collapses**. A collapsed mine blocks entry/use and
  **vanishes after one game year (8 months)**, after which a new mine may be built. The collapse timer
  joins §5.9's recorded countdowns.
- **`MAKE`** (Weaponsmithing `[617]`) turns **one input material into one item per day**:

  | item             | material   |
  | ---------------- | ---------- |
  | longbow `[72]`   | yew `[68]` |
  | plate armor `[73]`, longsword `[74]` | iron `[79]` |
  | pike `[75]`, crossbow `[85]` | wood `[77]` |

- **Turn lead into gold** (Alchemy `[697]`) is an alchemical gold source. Other production (scroll
  scribing, potion brewing, artifact forging, `BREED`) is catalogued in §4.7 with mechanics in §7.
- Any production that **mints a unique entity** (scribed scroll, forged artifact) advances the
  **deterministic counter** of §3.8 — never `rand`/`time` (§4.9).

### 6.6 Markets & trade goods ✅

- **`BUY`/`SELL` match only in cities** (§2.5), at the local bazaar; `GIVE` moves items between
  co-located nobles anywhere and is the off-market transfer path (§4.7). A trade executes when the
  seller's price ≤ the buyer's max, the buyer can afford ≥1 unit, the seller holds ≥1 unit, and both
  are in the **same city**. It settles **at the seller's price**.
- **Partial trades execute** (buy 25 of 100 wanted; the pending order shrinks). Unmatched orders
  become **standing orders** until executed or cleared (`buy <item> 0` clears). Orders the unit
  **cannot currently honor** (penniless buyer, empty seller) are omitted from the market report.
- **Deterministic clearing (a §11 ordering fact):** a buyer with multiple sellers takes the **lowest
  price**, ties broken by **position in the location's character list** (units resident longest sit
  toward the top and win ties); a seller with multiple buyers takes the **highest** in that list.
  **City self-trades have lowest priority** — cities defer to player units. The clearing must be a
  pure, reproducible pass over the recorded per-location character order (§6.8).
- **Privacy:** middlemen can hide a trader's identity (Conceal identity of trader `[731]`); some
  trades are omitted from the report entirely.
- **Trade goods (§4.8):** a tradegood is a *role over the item table*, not a new entity. Trade `[730]`
  sub-skills surface opportunities — **Find tradegood for sale `[732]`**, **Find market for tradegood
  `[733]`** — and profit comes from moving goods between city markets. **Cities issue their own
  `BUY`/`SELL`** on certain goods, resolved identically to player trades (but at lowest priority).

### 6.7 Opium & tax depression ✅ (mechanism) / ❓ (magnitudes)

- **Opium `[93]`** is produced in **swamp** provinces (Harvest opium `[706]`; Improve opium
  production `[707]`) and **consumed by markets in desert, plain, forest, and mountain** provinces —
  **not** swamp markets. Every non-swamp market carries **some opium demand**, hidden in the report at
  low levels.
- **Demand is a feedback loop:** satisfying demand **raises next month's** demand (peasants grow
  addicted); leaving it unmet lets demand **fall**. Demand becomes visible in the market report once
  high enough.
- **Opium consumption depresses the province's tax base**, scaling with the volume sold — the exact
  magnitude is **unspecified ❓**. This depression sits alongside **pillage recovery** (§5.2) as a
  recorded **depression/timer state** on the province (§5.9).

### 6.8 Architectural implications

The architectural consequences of §6 have moved to their correct homes per the routing rule in
[AGENTS.md](AGENTS.md); this section remains only as a pointer (§6.2, §6.6, §11, and §13 link to its
anchor, as do later chapters' §X.9 sibling lists):

- **The economy is pure resolution arithmetic over the snapshot** — income collection, upkeep,
  training, production, and market clearing are deterministic transforms of recorded state, and every
  "the computer chooses" (which unpaid men starve, which equal trade matches) resolves from recorded
  order, never `rand`. This is the standing **domain-purity** rule (the domain imports no entropy or
  clock — [AGENTS.md](AGENTS.md)) realized for stochastic draws by the `RNG` port
  ([ADR 0003](docs/adr/README.md)); its descriptive model awaits the deferred `reference/model/` page
  (see the note below).
- **Market clearing reads the per-location character-list (arrival) order** → already pinned in
  [`docs/adr/`](docs/adr/README.md): the snapshot must round-trip "the per-location arrival-order
  list" under the State-storage constraints, so the tie-break (position in the location's character
  list) and the single deterministic clearing pass over it are reconstructible. The clearing rule
  itself is descriptive mechanics (§6.6 body) awaiting the same `reference/model/` page.
- **Upkeep is stack-aggregate billing debited to a concrete holder** — capacity/upkeep are computed
  for the stack (§3.4) but gold leaves a specific noble's inventory, with the "ask same-faction
  stack-mates for gold" step a deterministic, faction-scoped gather. Descriptive model (§6.2 body),
  reusing §3.8's stack model, awaiting the deferred `reference/model/` page.
- **Production reuses the typed-count and unique-mint models** — gathering/`MAKE`/`TRAIN` mutate
  **fungible rows** (§4.1); scribing/forging mint **unique entities** via the §3.8 counter, already
  pinned as the **entity-number allocation counter** in the State-storage constraints in
  [`docs/adr/`](docs/adr/README.md). No new data model — economy is arithmetic over §3/§4's
  possessions.
- **Money flow is a §11 ordering invariant** — garrison upkeep before the castle's half before the
  owner's payout; tax computed, spent, and discarded within the same turn (no carryover). The
  sequence is what makes a re-run of `(gameID, turn)` reproduce the same balances; that rerun
  guarantee is the idempotency rationale in
  [explanation/idempotency.md](docs/content/explanation/idempotency.md) (`TurnLedger` + input hash),
  riding on `ProcessTurn` as a pure sequential transform ([ADR 0002](docs/adr/README.md)).
- **Depression and mine-collapse timers live in the snapshot** → already pinned in
  [`docs/adr/`](docs/adr/README.md): pillage recovery, opium demand, and mine collapse (8 months) are
  recorded, deterministic countdowns under the "all timer/countdown state" the snapshot must
  round-trip (§5.9 already routed these), joining §4's decomposition/return countdowns.

> **Not yet distilled.** Like §3–§5, §6's decided facts (the upkeep, training, construction, and
> improvement tables; the money-flow split; the market-clearing rule) wait on the orders pass (§10)
> before promotion to a `reference/model/` page — the orders that drive them (`RECRUIT`, `TRAIN`,
> `BUILD`, `IMPROVE`, `MAKE`, `BUY`/`SELL`, `PILLAGE`) may still reshape the slots. The recruit-supply
> model (§6.3), un-garrisoned-castle collection (§6.1), and the opium/yield magnitudes (§6.5, §6.7)
> remain ❓ and are carried forward.

## 7. Skills, magic & religion 🟡

The knowledge model every other subsystem reads from: the **category / sub-skill tree**,
`STUDY` and `RESEARCH`, per-skill **experience** (apprentice → grand master), the **schools of
magic** and the **aura** they spend, and **religion** — priests, prayers, and temple offerings.
§3.3 reserved each noble's **skills + experience** and **aura (current/max)** slots and deferred
their mechanics here; this section closes those, plus the deferrals routed in from elsewhere:
**per-skill NP costs** (§3.7), **resurrection / `LAY TO REST`** acting on `Body` items (§3.6,
§4.6), the **Skull of Bastrestric** aura burst (§5.8), and the production skills whose *yields*
were left to §7 (§4.7, §6.5). It owns **what a skill is, how it is learned, and how magic and
religion spend aura and NP**; per-use *yields* of the economic skills stay with §6, and **when**
study/casting/replenishment resolve within a turn is a §11 concern, marked where it bites.
Primary sources: [skills-magic.md](docs/content/rules/skills-magic.md),
[tables.md](docs/content/rules/tables.md) ("Skill listing and learning times"); the glossary
entries for *Aura*, *Apprentice*, *Grand master*, *Magician status*, *Prayer*, and *Priest*
anchor the vocabulary. Cross-refs: §3.3 (slots), §3.7 (NP), §4.5 (scrolls), §6 (production
yields), §5.8 (Skull), §8 (combat skills), §9 (ship skills).

### 7.1 The skill tree & knowledge states ✅ (model) / 🟡 (per-skill mechanics)

- Skills form a **two-level tree**: **category skills** (Shipcraft `[600]`, Combat `[610]`,
  Stealth `[630]`, Beastmastery `[650]`, Persuasion `[670]`, Construction `[680]`, Alchemy
  `[690]`, Forestry `[700]`, Mining `[720]`, Trade `[730]`, Religion `[750]`, the magic schools
  `[800]`–`[920]`) and **sub-skills** beneath them. **The category must be known before any of
  its sub-skills may be studied** (a hard invariant). ✅
- A skill is identified by its **entity number** — a low fixed fixture in the entity-number
  namespace (§3.2), exactly like an item *type* (§4.1). A noble does not *own* a skill entity;
  its **skills + experience** slot (§3.3) holds a **per-noble set of references** to skill numbers,
  each carrying that noble's own state. This is the same reference-to-authored-type model as
  fungible items and men (§4.1, §3.4) — skills are authored rows; knowledge is per-noble. ✅
- A noble's relationship to a skill is one of **four states**, all recorded per-noble in the
  snapshot:
  - **unknown** — not referenced;
  - **studying** — `STUDY` issued, accumulating study-days toward the skill's required total
    (the report's `7/14`);
  - **partially known** — surfaced by `RESEARCH` (the report's `0/7`); discovered but not yet
    usable, and **must still be `STUDY`-ed** to become known;
  - **known** — fully learned and usable, carrying the **experience use-count** (§7.5).
- **Most skills are invoked with `USE <skill> [args]`**; each known skill yields a **lore sheet**
  in the turn report describing its arguments and limits. The lore-sheet *content* is reporting
  (§12); §7 fixes only that knowing a skill unlocks its `USE`. The **per-skill mechanics** (what
  each `USE` consumes and yields) live with the consuming section — production with §6, combat
  with §8, ships with §9 — and are 🟡 here.

### 7.2 The authored skill-type table ✅ (shape) / 🟡 (numbering, full roster)

- Like the map (§2.1) and the item-type table (§4.2), the **skill catalog is a fixed authored
  artifact** — immutable input to resolution, never mutated by it. The master roster is
  [tables.md](docs/content/rules/tables.md)'s "Skill listing and learning times". Each skill row
  carries static attributes the engine reads:
  - **category vs. sub-skill** (and, for a sub-skill, its parent category);
  - **learning time** (authored in weeks; §7.3 converts to study-days);
  - **NP cost to begin study** (0 for most; §7.3);
  - whether it is **rated for experience** (§7.5);
  - **source constraints** — commonly known (studyable anywhere once the category is known),
    location-taught, or research-/scroll-only (§7.3).
- The roster spans the production and combat skills already referenced by earlier sections
  (Weaponsmithing `[617]`, the gathering skills `[700]`–`[723]`, the `TRAIN` skills `[610]`/`[615]`/
  `[616]`/`[601]`/`[750]` of §6.3, Beastmastery `[650]`+ for beasts §3.4, Record skill on scroll
  `[692]` §4.5) — so §7's table is the **single source those sections' skill numbers resolve
  against**. The on-disk numbering scheme inherits §3.2's open question (decimal vs. base-N
  alphanumeric), and the full sub-skill roster is illustrative, not closed. 🟡

### 7.3 STUDY ✅ (mechanics) / 🟡 (values)

- **`STUDY <skill> [fast-days]`** learns a skill over time. To *begin* study the first time, a
  **source of instruction** must be available — one of: the skill is **commonly known** (listed
  on the known category's lore sheet, then studyable anywhere); the noble is **in a city that
  teaches it**; the skill is **partially known** via prior `RESEARCH`; or the noble **holds a
  book or scroll** that teaches it (§4.5). Sub-skills *not* on the category lore sheet are
  **research-/scroll-only** and cannot be studied from the parent alone, even if their number is
  known from another player. ✅
- **Fees & gates:** beginning study of any skill costs a **flat 100 gold** (charged once, when
  `STUDY` is first issued for that skill); **advanced skills also cost Noble Points** (§3.7),
  authored per-skill in the catalog — Beastmastery `[650]`, Religion `[750]`, and the magic
  schools each require **1 NP**, Necromancy `[900]` **2 NP**; sub-skills are mostly free, a few
  heroic-combat and advanced-magic spells excepted. NP is spent from the **faction pool** (§3.7),
  gold from the **studying noble** (drawing on same-faction stack-mates, §6.2). ✅
- **Study limit:** a noble may apply at most **14 study-days per turn** to study. Learning times
  are authored in **weeks**; the engine reads them as **7 study-days per week** (so a 3-week
  category skill needs 21 days — at least two turns of study). The worked examples in
  [skills-magic.md](docs/content/rules/skills-magic.md) use simplified day counts that differ from
  the table; the **"Skill listing and learning times" table is authoritative**, the prose examples
  illustrative. 🟡 (the exact starting NP/gold balances and the final per-sub-skill NP list are
  §3.7/§10 values.)
- **Fast study:** a faction begins with **200+ "fast-study" points** (a faction-level pool, like
  NP). A fast-study point applied via `STUDY <skill> <n>` substitutes for a day of study, **does
  not count against the 14-day limit**, and lets the order complete in **0 days**. ✅ model; the
  exact starting amount is 🟡.
- **Idempotence of knowledge:** once a skill is **known**, further `STUDY` of it has no effect.
  Direct character-to-character teaching does not exist — knowledge moves between nobles **only**
  via scribed scrolls (Record skill on scroll `[692]`, §4.5). ✅

### 7.4 RESEARCH ✅

- **`RESEARCH <category>`** hunts for **hidden sub-skills** not granted by the category lore
  sheet — chiefly magic spells, but any category (Shipcraft, Combat, Construction, …) may hide
  sub-skills. **Category skills themselves cannot be discovered by research.** ✅
- **Mechanics:** research costs a **flat 25 gold** (materials) and yields a **25% chance per week**
  of discovering one new sub-skill, which lands in the noble's **partially-known** list (§7.1) and
  must then be `STUDY`-ed to become usable. ✅ The 25%/week roll is **randomness ⇒ derived from the
  recorded per-turn seed** (§2.9), never live entropy.
- **Location gate:** research for every category **except Religion** must be done in a **tower**,
  by the **tower's owner** (the first character inside, §5.3) — other occupants may not research.
  **Religion `[750]` research is done in a temple**, by the temple's owner. **Black-circle
  restriction:** a mage of **maximum aura ≥ 31** ("6th black circle and above") may research only
  in a province of **civ level ≤ 1** (§2.7). ✅

### 7.5 Experience & skill levels ✅

- Each **known** skill carries a per-noble **use-count**; one **successful use per turn** counts
  (further uses that turn do not). **Multi-turn projects** (shipbuilding, castle construction)
  count **only when the project finishes**. ✅
- Use-count maps to a **level label**, shown in the skill listing:

  | uses  | level        |
  | ----- | ------------ |
  | 0–4   | apprentice   |
  | 5–11  | journeyman   |
  | 12–20 | adept        |
  | 21–34 | master       |
  | 35+   | grand master |

- **Experience speeds work** for some skills (a master shipbuilder builds a galley faster than an
  apprentice); how much each skill benefits is a §6/§8/§9 per-skill detail. ✅ **Some skills are
  not rated for experience** (e.g. Survive fatal wound `[611]`, Fight to the death `[612]`); their
  level is **omitted** from the report and the experience use-count is not tracked — an authored
  flag on the skill row (§7.2). ✅

### 7.6 Magic: aura, schools & casting ✅ (model) / 🟡 (per-spell values)

- **Casting needs three ingredients:** *knowledge* of the spell, *possession* of any **required
  item** (usually **consumed** by the attempt, §4.7), and a *sufficient current aura* level. ✅
- **Aura** is the per-noble **(current, maximum)** pair reserved in §3.3 — `0` for non-mages.
  §7 closes its dynamics:
  - **learning a spell raises both current and maximum aura by 1** (so aura grows with the spell
    book, not with study days);
  - **current aura replenishes +2 per turn**, capped at maximum;
  - **casting a spell debits current aura** by the spell's authored cost (minor spells 1, powerful
    spells 10+);
  - other aura sources exist (Meditate `[801]`, Tap health for aura `[808]`, the Skull `[403]`,
    an **auraculum** `[881]`) — modifiers on the same pair, mechanics 🟡.
  This makes aura **pure resolution arithmetic** over the recorded pair; the **per-spell aura
  costs and required items** are authored per-skill and 🟡 here. The **Skull of Bastrestric**
  (§5.8: `USE 403` → +50–75 current aura, capped at **5× maximum**, **25% chance it kills the
  mage**) is one such modifier; its kill roll derives from the §2.9 seed.
- **Schools.** Magic is **sub-skills of a magic-school category**, each needing NP to learn
  (§7.3): Magic `[800]`, Weather magic `[820]`, Scrying `[840]`, Gatecraft `[860]`, Artifact
  construction `[880]`, Necromancy `[900]`. Only some spells are commonly known; the rest are
  `RESEARCH`- or scroll-gated (§7.4). The prose names **six schools**, but the catalog also lists
  **Advanced sorcery `[920]`** (a 7th category, no NP); the **authored table wins** (920 exists),
  with the prose/table count noted as 🟡.
- **Where each school is taught is authored city data** (Magic in most cities; Weather in the
  Cloudlands cities; Scrying in the Faery Cities; Gatecraft in all safe-haven cities; Artifact &
  Necromancy in Hades cities — each also "randomly in non-safe-haven cities"). City skill-teaching
  is **seed data on the authored map** (§2.1, like safe-haven designation §2.8); the special-realm
  cities (Cloudlands, Faery, Hades) are **§2.8-deferred content**, so the precise teaching map is
  🟡.

### 7.7 Magician status & the black circle ✅ (display) / 🟡 (above 30)

- A **magician-status label** is a **cosmetic display** computed from **maximum aura**, shown in
  the turn report (`Osswid the Brave [5639], wizard, …`):

  | maximum aura | label    |
  | ------------ | -------- |
  | 6–10         | conjurer |
  | 11–15        | mage     |
  | 16–20        | wizard   |
  | 21–30        | sorcerer |
  | 31+          | (unlabeled — "black circle" tiers) |

  Below 6 there is no label. The spell **Appear common `[803]`** suppresses the label. ✅ The label
  above 30 is unspecified in the rulebook (shown as `??`); only the **black-circle research
  restriction** (§7.4, aura ≥ 31) is a real mechanic — the display label there is 🟡.

### 7.8 Religion: priests, prayers & resurrection ✅ (model) / 🟡 (prayer mechanics)

- **Learning Religion `[750]` makes a noble a priest** — there is **no separate piety rating**
  (§3.3); priesthood *is* knowing `[750]`. It costs **1 NP + 5 weeks**, is **studied only in a
  temple** (cities never teach it), and its **research is done in a temple** (§7.4). Its
  sub-skills are **prayers**. ✅
- **Temple offerings:** a temple yields **100 gold/month to its owner if the owner is a priest**
  (§5.3, §6) — the priesthood gate on that income lives here. ✅
- **Resurrection / `LAY TO REST`** (the deferral from §3.6 / §4.6): the prayers **Lay to rest
  `[752]`** and **Resurrect dead noble `[754]`** act on a `Body` **item** (§4.6). `Resurrect`
  **reverses the death type-transition** — a `Body` becomes a living noble again — so the `Body`
  must retain enough of the dead noble's state (identity, skills, NP, aura) to reconstitute it
  (the snapshot already carries this, §3.8 / §4.6); `Lay to rest` hastens the spirit's passing,
  bearing on the **12-turn decomposition timer** (§3.6). The exact ranges and success conditions
  are 🟡. Related prayers — Receive vision `[751]`, Remove blessing from soldiers `[755]`,
  Immunity from Vision `[756]` — and the **blessed-soldier** `TRAIN` at a temple (§6.3) read this
  catalog; mechanics 🟡.

### 7.9 Architectural implications

The architectural consequences of §7 have moved to their correct homes per the routing rule in
[AGENTS.md](AGENTS.md); this section remains only as a pointer (later chapters' §X.9 sibling lists
reference its anchor, e.g. §8.9):

- **The skill catalog is authored reference data** → the descriptive fact (an immutable static
  lookup resolution reads but never mutates, loaded like the map of §2.1 and the item-type table of
  §4.2 — per skill its category, learning-days, NP cost, experience-rated flag, source constraints,
  and per-use aura/item costs) belongs in `reference/model/`, awaiting the deferred model page (see
  the note below). Its on-disk format and loader are the **same open artifact-and-loader concern**
  as the "Map artifact format" register row and the planned `MapSource` port in
  [`docs/adr/`](docs/adr/README.md) — a sibling to it (a `SkillSource`-style loader), not yet
  surfaced as a separate decision (§13.7), exactly as the item-type table is (§4.9).
- **Skill knowledge is per-noble state referencing authored types** — the same reference-to-type
  model as fungible items and men (§4.1, §3.4). The descriptive per-noble slots (each skill's state
  — studying *n*/required, partially-known *n*/required, or known — and a use-count for known
  experience-rated skills) belong in `reference/model/` with the deferred page. That skills are
  **fixtures, not minted** — the entity-number table mints only nobles, unique items, and
  sub-locations — rides on the **entity-number allocation counter** already pinned in the
  State-storage snapshot constraints in [`docs/adr/`](docs/adr/README.md) (§3.8, §4.9).
- **Study, research, aura, and experience are pure resolution arithmetic** — gold/NP/fast-study
  debits, study-day accumulation, the +2/turn aura replenish capped at maximum, the +1/+1 on
  learning a spell, and use-count increments are deterministic transforms of recorded state. This
  is the standing **domain-purity** rule (the domain imports no entropy or clock —
  [AGENTS.md](AGENTS.md)); the two randomness sources here — **`RESEARCH`'s 25%/week** discovery and
  the **Skull's 25% kill** — go through the `RNG` port ([ADR 0003](docs/adr/README.md)), derived
  from the recorded **§2.9 turn seed**, never live entropy. The descriptive arithmetic awaits the
  deferred `reference/model/` page.
- **Death is reversible by a priest** — `Resurrect` (§7.8) turns a `Body` item back into a living
  noble, so the §3.6/§4.6 death type-transition must be **losslessly reconstructible** from the
  `Body`'s recorded state. Already pinned in [`docs/adr/`](docs/adr/README.md): the State-storage
  snapshot constraints round-trip **dead-body items with their death turn**, and `Lay to rest` plus
  the 12-turn decomposition write the same body-timer under the "all timer/countdown state" the
  snapshot must carry (§5.9 already routed this).
- **Fast-study is a new faction-level pool** alongside NP (§3.7) — a per-faction counter spent by
  any member noble's `STUDY`, while gold study/research fees debit the **studying noble**
  (stack-aggregate billing, §6.2) and NP debits the **faction**. The who-gets-debited routing is
  descriptive game mechanics (§7.3/§6.2 body); the faction-scoped pool is authoritative snapshot
  state alongside NP, carried by `GameStateStore`. Both await the deferred `reference/model/` page.

> **Not yet distilled.** Like §3–§6, §7's decided facts (the knowledge-state model, the
> experience-level table, the aura dynamics, the STUDY/RESEARCH fees and gates) wait on the orders
> pass (§10) before promotion to a `reference/model/` page — the orders that drive them (`STUDY`,
> `RESEARCH`, `USE`) may still reshape the slots. The full per-sub-skill NP costs (§7.3), per-spell
> aura/item costs (§7.6), prayer mechanics (§7.8), and the special-realm teaching map (§7.6) remain
> 🟡/❓ and are carried forward.

## 8. Combat 🟡

Battle resolution layered on §3/§4's possessions and §5's territory: the **combat ratings** men,
nobles, and items carry; the **hit sequence** and the **break point** that ends a battle; the
**wound** generation that feeds noble health (§3.3, §3.6); **front/rear** positioning (`BEHIND`)
and **missile** fire; **fortifications** and **sieges**; **prisoners**, loot, and escape; and the
**permission/attitude** model (`ADMIT`, `HOSTILE`/`DEFEND`/`NEUTRAL`/`DEFAULT`) that decides who
fights whom. This section closes the deferrals routed here from elsewhere: the §3.3 **combat
attitude** and **behind** slots and the **wound math** feeding health/death (§3.3, §3.6), §2.3's
**terrain combat effects** (swamp), §5.4's **castle shelters 500 men**, and §2.8's **safe-haven
"no combat" enforcement**. It owns **how a battle resolves**; the *combat ratings themselves* are
authored item-table data (§4.2), the *heroic combat sub-skills* are §7, and **when** combat
resolves within a turn is a §11 concern, marked where it bites. Combat is the **randomness-heaviest
subsystem** — every roll below derives from the recorded per-turn seed (§2.9), never live entropy.
Primary source: [combat.md](docs/content/rules/combat.md);
[tables.md](docs/content/rules/tables.md) supplies the ratings. Cross-refs: §2.3 (swamp), §3.3/§3.6
(health, death→`Body`), §4.2/§4.5 (item table, magical items), §5.4/§5.5 (castles, garrisons), §7
(combat skills), §9 (shipboard combat, no siege at sea), §10 (order syntax), §11 (resolution
timing).

### 8.1 Combat ratings & item bonuses ✅ (model) / 🟡 (modifier values)

- Every fighter carries three integer ratings — **attack, defense, missile** — read from the
  **authored item-type table** (§4.2; the table's "Combat ratings of fighters"). Men are rows in
  that table (§3.4), so their ratings are static authored data, not per-instance state: peasant /
  worker / sailor `(1,1,0)`, soldier `(5,5,0)`, pikeman `(5,30,0)`, swordsman `(15,15,0)`, knight
  `(45,45,0)`, elite guard `(90,90,0)`, crossbowman `(1,1,25)`, archer `(5,5,50)`, elite archer
  `(10,10,75)`. ✅
- A **noble's innate rating is `(80,80,0)`** — a noble attribute (§3.3), not a table row; an unarmed,
  untrained noble is still a formidable combatant. ✅
- **Authored rating modifiers** (static, from the table notes). The **shipboard** and **swamp**
  modifiers are tracked as **two independent modifiers** — they happen to share the same −25 value for
  knights/elite guard today, but recording them separately lets play-testing tune one without
  disturbing the other:
  - a **blessed soldier** fights as a regular soldier but has a **50% chance of surviving a hit**
    (§8.3);
  - **shipboard modifier** — applied when the battle is fought on a ship (§9): a **pirate gains +10**
    to attack and defense (land `(5,5,0)` → ship `(15,15,0)`), while **knights and elite guard take
    −25** (knight `45 → 20`, elite guard `90 → 65`);
  - **swamp modifier** — applied in a swamp province (closing §2.3's terrain-combat deferral; swamp is
    the only terrain with a combat effect): **knights and elite guard take −25** to attack and defense,
    and it bites only those two unit types (no effect on pirates or anyone else). ✅
- **Engagement rule:** **peasants, workers, and sailors fight only when attacked** — i.e. only when
  their party is the *target* of an attack, never as initiators. ✅
- **Item bonuses (✅ model):** a noble automatically wields combat items it holds — **one item per
  category** (attack, defense, missile), the **largest bonus** in each category winning; no order is
  needed to wield them. These bonuses come from **unique magical items** (§4.5) and add to the
  noble's ratings for the battle. The concrete bonus *values* are per-instance item state (§4.5),
  🟡 here.
- **Heroic combat sub-skills (🟡, gated by §7):** the Combat `[610]` category's sub-skills modify a
  noble's behaviour in battle — **Survive fatal wound `[611]`** and **Fight to the death `[612]`**
  (both non-experience-rated, §7.5), **Defense `[614]`**, **Archery `[615]`**, **Swordplay `[616]`**.
  §8 reserves their effect on resolution (a chance to survive a killing blow; continuing to fight
  while wounded; rating bonuses); the exact mechanics are 🟡, owned by §7's skill model. ❓

### 8.2 Battle resolution: the hit sequence ✅ (algorithm) / 🟡 (timing)

A battle between an attacking side and a defending side resolves as a deterministic alternating
exchange (combat.md "Resolution of battle"):

1. a random man on the **attacking** side hits a random target on the **defending** side;
2. a random defender similarly hits a random attacker;
3. alternate until the **smaller** side has had as many chances to hit as it has attackers;
4. the **larger** side then takes **N** consecutive hits, where `N = (larger count − smaller count)`;
5. repeat the whole exchange until a side **breaks** (§8.3).

- **Hit chance.** An attacker with attack rating `Ar` hitting a defender with defense rating `Dr`
  succeeds with probability **`Ar / (Ar + Dr)`** (so `90` vs `45` ⇒ ⅔; `90` vs `90` ⇒ ½). The
  attacker's rating used is its **effective attack** = `max(attack, applicable missile)` (§8.4),
  plus item bonuses (§8.1) and authored modifiers (swamp/ship, §8.1). ✅
  > **Rulebook reconciled ✅:** combat.md wrote the denominator as `Ar + Br`; `Br` is a typo for the
  > defender's defense rating `Dr` (confirmed by the surrounding prose and the worked-example table).
  > Corrected in the rulebook in this pass — a typo fix, no design divergence.
- **Targeting order.** Within a side, the **stack leader** (top-most unit) is the **last** to take
  hits regardless of its `BEHIND` status, and **rear** units are reached only after the front is
  killed (§8.4). ✅
- **Timing (🟡):** *when* a battle resolves relative to movement, market clearing, and pillaging
  within a turn — and the order of multiple battles in one location — is a §11 ordering concern that
  matters for idempotency (the `TurnLedger`). The *algorithm* is fixed here; its *placement* in the
  turn is deferred.

### 8.3 The break point ✅ / reconciled

- A side **breaks** (loses) when its **combat value** — the sum of `attack + defense` over its
  members still able to fight — **falls below 50% of that side's starting combat value**. Dead men
  and **wounded nobles** (who stop fighting, §8.4) are removed from the running total. ✅
- Worked examples (combat.md): a noble + two pikemen start at `(80+80) + (5+30) + (5+30) = 230`,
  break point **115** — losing both pikemen (−70) leaves 160, so the noble **fights on**. A noble +
  two knights start at `(80+80) + (45+45) + (45+45) = 340`, break point **170** — losing both knights
  (−180) leaves 160 < 170, so the side is **declared the loser**. ✅
- **Terminology reconciled ✅:** combat.md overloads "offensive value" for two different quantities —
  the **break-point sum** (`attack + defense`, used here) and the **per-hit attacking rating**
  (`max(attack, missile)`, §8.2/§8.4). This file names them distinctly — **combat value** for the
  break-point total and **effective attack** for the per-hit rating — to keep §11's implementation
  unambiguous.

### 8.4 Wounds, death & the kill rule ✅ (closes §3.3 wound math)

This closes the wound-math deferral from §3.3/§3.6. A successful hit resolves differently against a
noble than against a man:

- **Against a noble** — the noble receives a **random wound of 1–100 health points** (combat.md;
  consistent with health-death.md's "wound randomly 1–100"). There is a **1% chance a perfectly
  healthy noble is killed outright**, and a **greater chance for an already-wounded noble** (the
  wound is rolled against current health). A wounded noble **stops fighting** for the rest of the
  battle even if the wound is minor (and is removed from the side's combat value, §8.3). The wound
  feeds the **illness check** (chance of falling sick = `100 − health`, health-death.md) and, at
  health **0** or a killing blow, the **death type-transition to a `Body` item** (§3.6, §4.6). ✅
- **Against a man (fighter)** — a hit **kills** the man, **except a blessed soldier**, who has a
  **50% chance of surviving** (§8.1). Men have no health rating — alive or dead (§3.3). ✅
- **NPCs rated `n/a`** for health require a hit of **≥ 50** to be killed (health-death.md; same rule
  for `terrorize`, aura blasts, lightning bolts). ✅
- All rolls here — wound magnitude, the 1%/scaled instant-kill, the illness check, the blessed-soldier
  50% — derive from the **§2.9 turn seed**. The *weekly* health update (sick lose 3–15, wounded
  recover, 5%/10%-in-an-inn illness-shake) is **resolution timing**, deferred to §11; §8 owns only the
  **wound generation** that feeds it.

### 8.5 Front, rear & missile fire ✅ (closes §3.3 behind slot)

This closes the §3.3 **behind** slot:

- **`BEHIND`** declares a unit's battle line — **front** or **rear**. Rear units are **not targeted
  until every unit in front of them is dead**; only **missile fighters** (non-zero missile rating)
  may attack from the rear. A fighter with **missile 0 cannot attack from the rear** and instead
  fights as if in front (using its attack rating). The **stack leader** is always the **last** to
  take hits regardless of `BEHIND` (§8.2). ✅
- **Missile vs. melee:** a **front** unit attacks with `max(attack, missile)` (its effective attack,
  §8.2); a **rear** unit attacks with its **missile** rating. So a noble `(80,80,40)` does 40 from
  the rear, 80 from the front. ✅
- **Weather cuts missile fire** (deterministic from the §2.4 weather variance, itself seeded from
  §2.9): **rain or wind** halve archer / elite-archer missile and cut crossbowman missile to **¼**;
  **fog** cuts **all** missile ratings to **¼**. ✅ The per-turn weather value is the §2.4 ❓ variance
  model; §8 records only how a given weather state modifies missile ratings.

### 8.6 Fortifications & sieges ✅ (closes §5.4 castle shelter)

- A **structure** with a **defensive rating** adds that rating to the defense of the **men who fit
  inside** it during a battle fought at its location. Capacity (closing §5.4's "castle shelters the
  first 500"): **castle 500, tower 100, galley/roundship 50, other structures 50**. ✅
- **Attacking a structure:** an attacking fighter may randomly target the **structure** instead of an
  enemy fighter; the hit resolves as fighter-vs-fighter using the attacker's attack rating against the
  structure's defense rating, and a successful hit **lowers the structure's defense by 1**. Once
  defense reaches **0**, further hits accrue **damage**; at **100% damage the building collapses**,
  ejecting its occupants. ✅
- **Siege engines** (`catapult [61]` `(25,200,25)`, `battering ram [60]` `(30,250,0)`, `siege tower
  [62]` `(30,250,0)` — §4.2) **always target the structure**, doing **5–10 damage per hit** (a plain
  fighter attacking a structure does 1). Siege engines are **not used in combat at sea** (§9). ✅ They
  are produced via skill-driven construction (Construct catapult `[613]`, siege tower `[681]`,
  battering ram `[701]`) — mechanics with §6/§4.7; their combat ratings are authored item-table data
  (§4.2). The 5–10 damage roll derives from the §2.9 seed.
- A **garrison** (§5.5) is the defending force that holds a province against **pillaging**; a pillager
  must defeat it in battle first (§6.1). Garrison combat uses these same rules. ✅

### 8.7 Prisoners, loot & escape ✅ (mechanism) / derived from §2.9 seed

- **Taking prisoners.** The winner attempts to capture defeated units; the chance a given defeated
  unit is taken is **proportional to numerical advantage** — **1:1 ⇒ 25%, 2:1 ⇒ 50%, 3:1 ⇒ 75%**,
  with a hard **floor of 25% and cap of 75%** regardless of ratio. (Capture reflects manpower to run
  down fleeing enemies, not combat skill.) Uncaptured defeated units **retreat**; those in a building
  or city may **flee into the surrounding province**. ✅
- **Loot & position.** The **stack leader of the winning force** receives **all** loot. Prisoners are
  **stripped of their belongings**, including accompanying men (workers, peasants). The victor
  **claims the defender's position in the location list** if better, or **moves into the defender's
  structure, ejecting the loser** — an `ATTACK` flag can inhibit this (§10). When a unit is taken
  prisoner (in battle or via **`SURRENDER`**), **a portion of its items is always lost**; a lost
  **unique item** (§4.1) must be re-found by `EXPLORE` of the province. ✅
- **Prisoner state.** A prisoner **reports nothing** (contributes nothing to its faction's report),
  **executes no orders** (queued orders stay pending), and is shown to others as a **stacked unit
  marked `prisoner`**. **`UNSTACK <prisoner>` frees** it; **`GIVE`** transfers a prisoner between
  co-located units (§4.7). ✅
- **Escape (weekly + event rolls, all seeded §2.9):** each game **week** (4×/turn, §8.4's weekly
  cadence), a prisoner held **outside** a building escapes with **2%** chance, **inside** a structure
  (castle/tower/inn/ship) with **1%**. Additionally **+2%** on each `GIVE` transfer and **+2%** each
  time the holder **travels more than one day** (short hops — entering/leaving a building — add no
  escape chance). A freed prisoner **flees into the surrounding location**; one freed on a **ship
  leaps overboard and swims to a nearby shore**. ✅

### 8.8 Permissions & declared attitudes ✅ (closes §3.3 attitude slot)

This closes the §3.3 **combat attitude** slot. Both a **noble** and the **faction player entity**
(§3.1) keep three attitude lists; any unit not on a list has attitude `DEFAULT`.

- **`ADMIT`** governs **stacking and entry**: by default a unit may not stack with another faction's
  character, nor enter a building or ship another player controls (§3.3, §5.3); `ADMIT` grants the
  exception. ✅
- **Four combat attitudes** toward a unit or faction: **`HOSTILE`** (attack on sight), **`DEFEND`**
  (aid the other unit if it is attacked), **`NEUTRAL`** (do nothing), **`DEFAULT`** (neutral to other
  factions; **defend own-faction units** unless either is concealing its lord). ✅
- **Resolution rules:**
  - **Unit attitude beats faction attitude** — one may declare a faction `HOSTILE` yet exclude a
    specific unit as `NEUTRAL`. ✅
  - **Concealment defeats it:** declaring an attitude toward a *player* works only while that
    player's units are **not concealing their faction** (Conceal faction `[635]`, Stealth `[630]`); a
    unit **concealing its lord** is neither attacked on sight nor expected to give itself away by a
    `DEFAULT` defense (it must be told to `DEFEND` explicitly). ✅
  - **Only the top-most unit aids in defense**, and when it joins a battle via `DEFEND` it **brings
    its whole stack**, even stack-mates that did not declare `DEFEND`. ✅
  - **Defenders help only when the protected unit is *attacked*, never when it initiates** an attack.
    Units that join via `DEFEND` are marked **`ally`** in the combat report. ✅
  - Units declared `DEFEND` toward a province's **guards** aid them when the guards are attacked,
    whether explicitly (`ATTACK`) or implicitly (`pillage 1`) — tying attitude to §5.5 garrison
    defense and §5.7 decrees (`DECREE WATCH WHO`, attack-on-sight). ✅

### 8.9 Safe havens: where combat is forbidden ✅ (closes §2.8)

- A **safe haven** (the authored city designation, §2.8) **forbids combat and magic** within it. §8
  closes the enforcement: **`ATTACK` and other hostile actions are rejected** in a safe-haven city
  (alongside the magic prohibition, §7). The set of safe-haven cities is **authored map seed data**
  (§2.1/§2.8), read immutably; §8 owns only the rule that combat orders **fail** there. ✅ The exact
  rejected-order set is finalized with §10.

### 8.10 Architectural implications

The architectural consequences of §8 have moved to their correct homes per the routing rule in
[AGENTS.md](AGENTS.md); this section remains only as a pointer (later chapters' §X.9/§X.8 sibling
lists reference its anchor, e.g. §9.8, §10.8, §11.9). Combat introduces **no new** port, open
decision, or snapshot field — it is the densest *consumer* of seams §2–§7 already established, so
every consequence routes to a home those chapters already populated:

- **Combat is the densest consumer of the §2.9 turn seed.** Random attacker/target selection, the
  `Ar/(Ar+Dr)` hit roll, the 1–100 wound and instant-kill chance, the illness check, the
  blessed-soldier 50%, siege 5–10 damage, prisoner capture, and every prisoner-escape roll are
  **pure functions of recorded state + the recorded seed**. This is the standing **domain-purity**
  rule (the domain imports no entropy or clock — [AGENTS.md](AGENTS.md)); the draws go through the
  `RNG` port ([ADR 0003](docs/adr/README.md)), seeded from the recorded turn seed whose state the
  per-turn snapshot round-trips ([`docs/adr/`](docs/adr/README.md) State-storage constraints). That
  re-running `(gameID, turn)` reproduces the same battle outcomes is what makes combat idempotent —
  the `TurnLedger` rationale already lives in [explanation/idempotency.md](docs/content/explanation/idempotency.md).
- **Combat ratings are authored reference data, not stored state.** Per-unit-type `(attack, defense,
  missile)` and the shipboard/swamp/blessed modifiers live in the §4.2 item-type table; the noble's
  innate `(80,80,0)` and any **magical-item bonuses** (§4.5) are the only combat values carried on
  entities, and resolution reads ratings without mutating them. The descriptive ratings model
  belongs in `reference/model/` with the deferred page (see the note below); its on-disk format and
  loader are the **same open authored-data artifact-and-loader concern** as the "Map artifact
  format" row and the planned `MapSource` port in [`docs/adr/`](docs/adr/README.md) (a sibling
  loader, not a separate decision — §13.7, exactly as the item-type and skill tables are routed in
  §4.9/§7.9).
- **A battle is a transform over a location's occupants.** It reads the per-location character list
  (the §6.6 ordered arrival list already pinned in the snapshot constraints in
  [`docs/adr/`](docs/adr/README.md)), the stack trees (§3.3), `BEHIND` flags, attitude lists, and
  structure defense ratings; it writes wounds, deaths (noble → `Body`, §4.6), killed-men counts,
  prisoner links, loot transfers, and structure damage. No new aggregate — combat mutates the
  §3/§4/§5 state already in the snapshot. This descriptive resolution model belongs in
  `reference/model/` with the deferred page.
- **Attitudes and `BEHIND` are recorded per-noble and per-faction.** The three attitude lists live on
  both the noble and the faction player entity (§3.1); `BEHIND` and the (now-closed) combat-attitude
  slot occupy the §3.3 reservation; concealment (§7 Stealth) gates whether an attitude can even be
  applied. These descriptive per-noble/faction slots belong in `reference/model/` with the deferred
  page, alongside the §3 attribute model.
- **Structure damage and the kill rule reuse existing timers/transitions.** A collapsed building
  ejects occupants (no new state beyond §5's damage/collapse tracking); a killed noble reuses the
  §3.6/§4.6 death transition and decomposition timer. Combat introduces **no new persisted timer** —
  it feeds the timer/countdown state and the dead-body-with-death-turn fields the snapshot already
  carries, **already pinned** in [`docs/adr/`](docs/adr/README.md)'s State-storage constraints
  (routed in §5.9/§7.9).

> **Not yet distilled.** Like §3–§7, §8's decided facts (the ratings model, the hit sequence and
> break point, the wound rule, fortification capacities, the prisoner/escape and attitude rules) wait
> on the orders pass (§10) before promotion to a `reference/model/` page — the orders that drive them
> (`ATTACK`, `DEFEND`, `BEHIND`, `HOSTILE`/`NEUTRAL`/`DEFAULT`, `ADMIT`, `SURRENDER`, `PILLAGE`) may
> still reshape the slots. The heroic combat-skill effects (§8.1 → §7), per-spell/aura combat magic
> (§7.6), the §2.4 weather-variance model feeding missile penalties (§8.5), and the §11 placement of
> combat within the turn remain 🟡/❓ and are carried forward.

## 9. Ships 🟡

Ships as **mobile, crewed entities** that bridge §2's map (ocean travel), §6's economy (construction,
ferries), and §8's combat (battle at sea). This section owns the **two ship types** and their authored
stats; how a ship is **owned and what it holds** (occupants + cargo, measured in §4.3 weight); how one
is **built** (the §6.4 construction machinery, gated by Shipbuilding and confined to port cities); how
it **sails** (the crew gate, the §2.5 dock, the §2.4 wind-modified travel time); how it takes and sheds
**damage** (storms, rocks, `REPAIR`); and the **ferry** model (`FEE`/`BOARD`/`FERRY`/`UNLOAD`). It
**reuses, never re-decides**: ocean movement and docking are §2.4/§2.5, the fifths-staged build is §6.4,
and **all battle-at-sea rules are §8** (the galley/roundship defensive rating and 50-occupant shelter
§8.6, the pirate `+10` / knight–elite-guard `−25` shipboard modifiers §8.1, no siege engines at sea
§8.6, a freed prisoner leaping overboard §8.7). A ship's identity is an **entity number** (§3.2), not a
sub-location code — but it *acts as* a location that stacks hold and occupants ride. Like combat, the
travel hazards here are **seeded** (§2.9), never live entropy. Primary source:
[ships.md](docs/content/rules/ships.md); [tables.md](docs/content/rules/tables.md) supplies the
Shipcraft skill rows and the shelter capacity. Cross-refs: §2.4/§2.5 (ocean travel, docking), §3.2
(entity number), §4.3 (weight/capacity), §4.5 (per-instance state), §6.3/§6.4 (training sailors,
construction), §7 (Shipcraft mechanics, storm magic), §8 (combat at sea), §10 (`SAIL`/`BOARD`/… syntax),
§11 (resolution timing).

### 9.1 The two ship types ✅ (authored table)

- Two ship types exist — the **galley** (warship: slender, rowed) and the **roundship** (merchantman:
  deep, wide, sail-driven). Their stats are **authored reference data**, a fixed ship-type table read
  immutably by resolution, never mutated by it — the same treatment as the item-type table (§4.2) and
  skill table (§7.2): ✅

  | ship      | cargo capacity | crew (to sail) | build effort    | build material | combat shelter / defense |
  | --------- | -------------- | -------------- | --------------- | -------------- | ------------------------ |
  | galley    | 5,000          | 14 sailors     | 250 worker-days | 50 wood `[77]` | first 50 occupants (§8.6) |
  | roundship | 25,000         | 8 sailors      | 500 worker-days | 100 wood `[77]`| first 50 occupants (§8.6) |

- **Cargo capacity is a §4.3 weight budget**, not a unit count — the rulebook's "units of cargo" are
  the universal weight units (§4.3), the same scale that drives carry capacity and ferry fees. A
  ship's effective capacity is reduced by damage (§9.5). ✅
- A **damaged** ship's capacity falls **in proportion to its damage percentage**: a galley at 10%
  damage carries `5,000 × 0.9 = 4,500`. ✅

### 9.2 Ships as entities: ownership, occupants & cargo ✅ (model) / 🟡 (capacity accounting)

- A ship is an **entity** in the entity-number space (§3.2), minted a fresh number when built — **not**
  a province `(row, col)` and **not** a static sub-location. Yet functionally it **acts as a location**:
  nobles (with their men and items) stack *inside* it, and when it sails the whole contents move with
  it. The map model (§2) must therefore represent "inside ship `[n]`" as a **containment edge keyed by
  entity number**, distinct from province coordinates and authored sub-location codes. ✅
- **Ownership is positional**, as for buildings (§6.4): the **captain** is the controlling unit aboard,
  and entry by another faction's character is gated by `ADMIT` (§5.3, §8.8) — except via the ferry
  `BOARD` path (§9.6). ✅
- A ship carries **per-instance state** (§4.5-style): its **name** (player-supplied, sanitized at
  ingest §10), **damage percentage** (§9.5), **construction progress** while in-progress (§9.3), and
  the **occupants + cargo** it holds. This is distinct from the **authored** type stats of §9.1 —
  capacity/crew/cost are read from the table, never stored per ship. ✅
- **Capacity accounting (🟡):** capacity is spent by the weight of everything aboard. Whether the
  **required crew's own weight** (a sailor weighs 100, §4.3) counts against the cargo budget, or rides
  "for free" as part of the vessel, is left 🟡 — settled with the §2.4 overload model it feeds.

### 9.3 Building a ship ✅

Ship construction **reuses the §6.4 machinery** — it is not a second build system:

- Requires the **Shipbuilding `[602]`** sub-skill of Shipcraft `[600]` (§7.2), at least **three
  workers**, and the §9.1 materials. The builder **unstacks** to the outer level and issues
  `BUILD GALLEY "name"` / `BUILD ROUNDSHIP "name"`. ✅
- **Materials stage in fifths** exactly as §6.4: **one-fifth of the wood is deducted immediately**
  (10 wood for a galley, of 50 total), the next fifth at 20% complete, etc.; construction **halts**
  if the builder runs out. The builder and workers are placed **inside** the new ship, which displays
  as `…-in-progress, NN% completed` until the worker-days are invested, then is **christened
  seaworthy**. Resume a partial hull by **entering** it and re-issuing the `BUILD` order. ✅
- **Ships may be built only in port cities** (§2.5) — the one placement rule §6.4's building catalog
  does not already cover. ✅

### 9.4 Sailing ✅ (gates) / 🟡 (travel-time model)

- **Ocean travel requires a ship** (§2.4); piloting one requires the **Sailing `[601]`** sub-skill
  (§7.2), learnable in any port city. The pilot must be aboard. ✅
- **Crew is a movement gate, not an existence gate.** A ship may sit in port under-crewed, but `SAIL`
  **fails** unless the full complement is aboard — **14 sailors for a galley, 8 for a roundship**
  (§9.1). Sailors are men trained from peasants via `TRAIN` under Sailing `[601]` (§6.3). ✅
- **`SAIL <direction | destination>`** is order priority **4** (§10); it moves the ship (and all it
  holds) along ocean routes. **Docking** is §2.5: a ship in an ocean province sails into an adjoining
  **land** province (1 day) and **cannot dock against mountains** (ocean↔mountain routes are
  `impassable`). ✅
- **Travel time is wind-modified (🟡).** Ocean routes are authored for "an ordinary ship in normal
  weather" (§2.4); wind speeds or slows a vessel, and a **roundship makes better time than a galley
  under favorable winds even when fully loaded**. §9 records *that* wind favors roundships; the
  concrete distribution is the **§2.4 variance model** (❓), deterministic from the §2.9 seed. The
  overload rules (§4.3) apply to a ship's contents as to any stack.

### 9.5 Damage & repair ✅ (model) / 🟡 (hazard rates)

- A sailing ship may be **damaged by storms or by submerged rocks in coastal waters** — a
  travel-resolution hazard, **deterministic from the §2.9 turn seed** (the same discipline as combat
  §8 and weather §2.4; the domain imports no entropy). The resulting **damage percentage** is
  per-instance ship state (§9.2) and **reduces effective capacity proportionally** (§9.1). ✅
- **`REPAIR [days]`** (order priority 3, §10) restores a damaged ship and **consumes one unit of
  pitch `[261]`**. ✅
- The **per-turn probability and magnitude** of storm/rock damage are **🟡** — they belong with the
  §2.4 weather-variance model (and interact with the storm-magic skills `[822]`/`[827]`/… of §7 that
  bind, direct, or dissipate storms). §9 records the hazard's *existence and effect*, not its rates.

### 9.6 Ferries ✅ (model) / 🟡 (order syntax & sync, with §10/§11)

A ship becomes a **commercial ferry** when its captain sets a fee:

- **`FEE <gold>`** sets the charge as **gold per 100 weight** of a passenger's stack (§4.3) — e.g.
  `fee 50` is ½ gold per weight. The **fee is a property of the captain (a noble), not the ship**
  (§9.2): it lives on the noble and is read from whichever captain is aboard. **`fee 0`** clears it;
  with **no fee set the ship is not a ferry** and `BOARD` is refused. ✅
- **`BOARD <ship> [max fee]`** (priority 2) pays the co-located captain the required fee, then moves
  the passenger's stack aboard. It **fails** if the ship is absent or not operating as a ferry. The
  captain **must not `ADMIT` paying passengers** — an `ADMIT`ted character could enter with `MOVE` and
  bypass the fee, so the ferry path deliberately routes around `ADMIT` (§9.2). ✅
- **Synchronization:** on arrival the captain issues **`UNLOAD`** (priority 3) to eject the current
  passengers, then **`FERRY`** (priority 1) to signal those waiting that they may now board; passengers
  use **`WAIT FERRY <ship>`** (never `WAIT SHIP`, which would `BOARD` before `UNLOAD` clears room).
  This four-order handshake is recorded here as a **model**; its exact syntax is finalized with §10 and
  its within-turn **ordering** (the priorities above, and `UNLOAD`-before-`FERRY`-before-`BOARD`) is a
  §11 resolution concern that bears on idempotency. 🟡

### 9.7 Combat at sea — owned by §8

§9 introduces **no** new battle rules; it only marks the ship-specific cross-refs already decided in §8:

- A galley/roundship has a **defensive rating** that shelters its **first 50 occupants** in a battle
  fought aboard (§8.6) — the same fortification mechanic as a castle/tower, at capacity 50.
- The **shipboard combat modifier** (§8.1) applies: a **pirate gains `+10`** attack/defense, while
  **knights and elite guard take `−25`** — pirates are trained from sailors *on a ship* (§6.3).
- **Siege engines are not used at sea** (§8.6); a freed prisoner aboard a ship **leaps overboard and
  swims to a nearby shore** (§8.7). ✅

### 9.8 Architectural implications

The architectural consequences of §9 have moved to their correct homes per the routing rule in
[AGENTS.md](AGENTS.md); this section remains only as a pointer (later chapters' §X.9/§X.8 sibling
lists reference its anchor, e.g. §10.8, §11.9). Like combat (§8.10), ships introduce **no new** port,
open decision, or snapshot field — §9 is a *consumer* of seams §2–§8 already established, so every
consequence routes to a home those chapters already populated:

- **Ship-type stats are authored reference data, not stored state.** The §9.1 table (capacity, crew,
  cost, shelter) is loaded immutably alongside the map (§2.1), item-type table (§4.2), and skill table
  (§7.2) — the **same open authored-data artifact-and-loader concern** as the "Map artifact format"
  row and the planned `MapSource` port in [`docs/adr/`](docs/adr/README.md) (a sibling loader, not a
  separate decision — §13.7, exactly as the item-type and skill tables are routed in §4.9/§7.9).
  Resolution reads these stats, never mutates them. The descriptive ship-type model belongs in
  `reference/model/` with the deferred page (see the note below).
- **A ship is a mobile container entity.** Its identity is an entity number (§3.2), but the location
  model (§2) must express **"inside ship `[n]`"** as a containment edge — distinct from province
  coordinates and from authored sub-location codes — so that a single `SAIL` transform relocates the
  vessel and everything stacked within it. This is the one place the entity-number space and the
  spatial graph genuinely meet. This descriptive containment-model fact belongs in `reference/model/`
  with the deferred page, alongside the §2 map model.
- **Per-instance ship state lives on the entity; type stats do not.** Name, damage %, construction
  progress, captain (positional, §6.4), and the occupant/cargo manifest are carried per ship (the
  §4.5 per-instance-state pattern); capacity/crew/cost are read from the authored table. Effective
  capacity is a *derived* value (`base × (1 − damage)`), computed at use, not persisted. This
  descriptive per-instance model belongs in `reference/model/` with the deferred page.
- **Build and travel introduce no new timers or entropy.** Ship construction reuses §6.4's
  fifths-staged progress (no new persisted countdown), and storm/rock damage is a **pure function of
  recorded state + the §2.9 seed** — the standing **domain-purity** rule (the domain imports no
  entropy or clock — [AGENTS.md](AGENTS.md)), the draws going through the `RNG` port
  ([ADR 0003](docs/adr/README.md)). Re-running `(gameID, turn)` reproduces the same hull damage and
  the same sailing outcomes, which keeps ship movement idempotent under the `TurnLedger` — the
  rationale already lives in [explanation/idempotency.md](docs/content/explanation/idempotency.md).
  Ship build and damage add **no new persisted timer** — they feed the timer/countdown state the
  snapshot already carries, **already pinned** in [`docs/adr/`](docs/adr/README.md)'s State-storage
  constraints (routed in §5.9/§7.9).
- **The ferry fee is per-captain (noble) state, read at `BOARD` time.** It is *not* a property of the
  ship entity (§9.6); the `BOARD` use case reads the co-located captain's fee and effects a
  pay-then-move. This descriptive per-noble slot belongs in `reference/model/` with the deferred page.
  The `FEE`/`BOARD`/`UNLOAD`/`FERRY`/`WAIT FERRY` handshake is order-driven (§10), and its priority
  ordering within a turn is a §11 resolution concern that bears on idempotency.

> **Not yet distilled.** Like §2–§8, §9's decided facts (the two ship types, the entity/container
> model, the §6.4-reusing build, the crew gate and dock, the damage/repair and ferry models) wait on
> the orders pass (§10) before promotion to a `reference/model/` page — the orders that drive them
> (`BUILD`, `SAIL`, `REPAIR`, `FEE`, `BOARD`, `UNLOAD`, `FERRY`, `WAIT FERRY`) may still reshape the
> slots. The wind-modified ocean **travel-time model** and the storm/rock **hazard rates** (both §2.4's
> ❓ variance model), the **capacity accounting** for crew weight (§9.2), and the **§11 placement** of
> sailing, ferry sync, and damage within the turn remain 🟡/❓ and are carried forward.

## 10. Orders 🟡

Orders are the engine's **sole input channel** and its **untrusted-input boundary**. A player composes a
plain-text order file and mails it in; the engine ingests it, queues per-unit commands, and resolves them
on the turn run. §10 owns the **order file itself** — its grammar (the `begin`/`unit`/`end` envelope),
the **command catalog** (the ~80 verbs, each with a fixed *priority* and *time* attribute), the
**per-unit command queue** (250-deep, `UNIT`-replaced, `STOP`-interruptible), the **two-stage validation
split** (what the scanner checks at parse time vs. what is checked only at execution), the **player-entity
admin order set**, and the **sanitization of player-supplied names**. It **reuses, never re-decides**:
every command's *meaning* lives in the section that owns its mechanic — `MOVE`/`FLY` travel is §2.4,
docking §2.5, `FORM`/`PLEDGE`/permissions §3/§5.3/§8.8, item transfer (`GET`/`GIVE`/`BUY`/`SELL`/…) §4/§6.6,
construction (`BUILD`/`RAZE`/`IMPROVE`) §6.4, `STUDY`/`RESEARCH`/`USE` §7, `ATTACK` and its flags §8,
`SAIL`/`BOARD`/`FEE`/`FERRY`/`UNLOAD`/`REPAIR` §9 — so §10 records only the **vocabulary and grammar**
that carry those mechanics, never their resolution. The hand-off is sharp: §10 assigns each command its
**priority value**; **§11 owns the scheduler** that consumes those values, the within-turn *ordering*,
*time accrual*, `STOP`/interrupt timing, and the idempotency placement. Primary source:
[orders.md](docs/content/rules/orders.md). Cross-refs: §2.4/§2.5 (`MOVE`/`FLY`/`SAIL` semantics), §3.1/§3.2
(player entity, entity-number space, `FORM`), §3.4 (faction→noble hierarchy the `unit` sections address),
§4–§9 (the owning section of every command's mechanic), §11 (the scheduler, timing, idempotency), §12
(the acknowledgement reply, the order template, the seen-here list).

### 10.1 The order file: structure & grammar ✅

- An order file is a **line-oriented text DSL** with a fixed envelope: a single `begin <player-number>
  [password]` header, then one `unit <number>` block per unit (its queued commands indented beneath it),
  closed by a **single `end`** — the parser stops reading at `end`, and there is **never** an `end` per
  `unit`. The first `unit` block is the **player (faction) entity** itself (§10.6); subsequent blocks are
  its characters (§3.4). ✅
- The grammar is **forgiving by design**: case-insensitive, whitespace/indentation-insensitive, `#`
  begins an end-of-line **comment** the engine never interprets, and a multi-word argument **must be
  quoted** (`name "Osswid the Constructor"`). The mail **Subject:** line is ignored. ✅
- **`UNIT` replaces, it does not append.** Re-sending a `unit` block **clears that unit's pending queue**
  and queues the new commands; units omitted from a later file keep their existing queue untouched. A unit
  may hold at most **250 queued orders**; excess is silently dropped (a gap §10.8 flags for an explicit
  ingest warning). ✅
- This DSL — not YAML, not a structured-field email schema — is the **decided order-file content format**,
  recorded as the rulebook's custom line-oriented DSL in [ADR 0001](docs/adr/README.md). Mail *transport*
  (how the file arrives) is the separate "Mail transport" decision; the
  DSL is simply the body it carries. The **exact tokenizer/grammar spec** (quoting edge cases, numeric vs.
  entity-code argument forms) is 🟡 — pinned down when the orderfile adapter is built. ✅ (format) / 🟡 (spec)

### 10.2 The untrusted-input boundary ✅

- `internal/infra/orderfile/` is **the** validation boundary, exactly as CLAUDE.md and AGENTS.md mandate:
  it owns the raw bytes, parses the DSL, and **rejects malformed input with `cerr.ErrInvalidOrders`** —
  no raw bytes ever reach `app` or `domain`. Only typed, validated **`domain.OrderBundle`** values cross
  inward; the bundle is **per-player** (the `OrderSource.ReadOrders` port returns `[]domain.OrderBundle`,
  one per submitting player), each carrying that player's parsed `unit`-keyed command queues. ✅
- Because the bundle is the **only** form orders take inside the engine, every later concern — scheduling
  (§11), idempotency hashing (§10.8/§11), report echoing of the queue (§12) — operates on typed values, not
  text. The adapter is **dumb** (no game rules), per CLAUDE.md's "keep infra adapters dumb." ✅

### 10.3 Two-stage validation: parse-time vs. execute-time ✅

The scanner deliberately does **not** fully validate orders, and this split maps cleanly onto the SOUSA
layer boundary:

- **Parse time (orderfile / infra):** the scanner checks only that each **command exists** and that the
  parameters of the **scan-affecting orders** are well-formed — those involved in parsing (`begin`,
  `unit`, `password`, `email`, `vis_email`), in report formatting (`format`, `notab`), and those with
  **immediate secondary effects** (`resend`, `lore`). A structural or lexical failure here is
  **`cerr.ErrInvalidOrders`**. This is *lexical/structural* validity only. ✅
- **Execute time (app/domain resolution, §11):** every *other* command's parameters are validated **only
  when the command runs** during the turn — does the location offer this skill, are there peasants to
  recruit, does the target exist. These are *semantic* checks against world state, so they belong to
  resolution, not the parser, and surface as the **business-meaning `cerr` sentinels** (§ `cerr`), not
  `ErrInvalidOrders`. A command that fails semantically **takes zero time** (§10.4) and does not consume
  the unit's monthly study/production budget. ✅
- Consequence: a file can parse cleanly (every verb known, envelope well-formed) yet have individual
  orders fail at run — the **acknowledgement reply** (§12) reports parse-time errors immediately, while
  run-time failures appear in the turn report. The two error channels are distinct on purpose. ✅

### 10.4 The command catalog: priority & time as authored attributes ✅ (shape) / 🟡 (per-command meaning)

- Every command carries two **fixed, authored attributes** — a **priority** in `0–4` and a **time class** —
  independent of any single invocation. This catalog is **reference data** (the same treatment as the
  item-type §4.2, skill-type §7.2, and ship-type §9.1 tables): read immutably, never mutated by resolution.
  The full table is [orders.md](docs/content/rules/orders.md)'s command summary; §10 records the **rules
  that generate it**, not a second copy. ✅
- **Priority** is assigned by rule: permission commands (`ADMIT`, `HOSTILE`, `NEUTRAL`, `DEFEND`, …) are
  **0**; zero-time commands and `WAIT` are **1**; `MOVE`/`FLY` are **2**; everything else is **3**; `SAIL`
  alone is **4**. §10 owns these *values*; the **§11 scheduler** consumes them (start all ready prio-0
  before any prio-1, …; ties broken by **location order**, §11/§12) — that algorithm is *not* §10's. ✅
- **Time class** is one of: **`0 days`** (instantaneous), a **fixed** count (e.g. `1 day`, `7 days`),
  **`as given`** (the player supplies the day count as an argument, e.g. `RECRUIT [days]`), or **`varies`**
  (computed from world state, e.g. `BUILD`, `MOVE`, `USE`). **Failed orders generally take zero time**
  (§10.3). *How* time accrues across the turn and *when* a multi-day order is interruptible are **§11**. ✅
- Each verb's **semantics stay in its owning section** (§10 intro). The catalog is **vocabulary**: it tells
  the parser each verb's arity and the scheduler each verb's priority/time, nothing about what the verb
  *does*. 🟡 marks that those per-command meanings are distributed, not centralized here.

### 10.5 The per-unit command queue ✅ (model) / 🟡 (`STOP`/interrupt timing → §11)

- Commands **queue per unit** and execute in submitted order. A unit already mid-command is **not
  interrupted** by a freshly-submitted `unit` block **unless the first new order is `STOP`**; `STOP` itself
  **queues like any other order** and only takes effect when the turn runs (so a later file can still
  replace it before the deadline). The precise moment `STOP` bites, and how the queue drains against the
  day clock, are **§11** resolution concerns. ✅ (model) / 🟡 (timing)
- Orders may be queued for units **not yet under control** (a `BRIBE`/`TERRORIZE`/`PLEDGE` target, §3/§8) —
  they begin executing the moment the unit joins the faction. ✅
- Orders may be queued for a **noble that does not yet exist**: the turn report lists the **next entity
  numbers** that `FORM` will mint (§3.2's reserved entity-number space), the player `FORM`s one with that
  number, and a `unit <that-number>` block queues the newborn's first orders. This is the one place ingest
  must accept a `unit` number that is not yet a live entity. ✅
- **Same-priority ties resolve by location order** — the order units are listed in a location ("seen
  here"), oldest-resident first, a freshly-arrived or freshly-`UNSTACK`ed unit reinserted per the §12
  rules. §10 only notes that location order is the tiebreaker the **§11 scheduler** reads; maintaining the
  list is §11/§12. ✅ (the tiebreaker exists) / 🟡 (consumed in §11)

### 10.6 The player (faction) entity & its restricted order set ✅

- The faction is itself an **entity** (§3.1/§3.2) — invisible, in no location, receiving no location
  report — used mostly as a placeholder. It may issue **only administrative orders**: `ACCEPT`, `ADMIT`,
  `DEFAULT`, `DEFEND`, `FORMAT`, `HOSTILE`, `MESSAGE`, `NAME`, `NEUTRAL`, `NOTAB`, `PRESS`, `REALNAME`,
  `RUMOR`, `TIMES`, `QUIT`. **Characters** issue everything else; one must **not** `FORM` or `RECRUIT` with
  the player entity. Renaming the faction is `NAME` issued in the player entity's own `unit` block. ✅
- **`QUIT`** is issued for the player entity to leave the game; **no turn report is sent** for the turn in
  which a player quits. ✅
- A distinct sub-class of orders acts at **scan/account level, not as queued unit commands**: `begin`/`unit`
  (envelope/routing), `password`/`email`/`vis_email`/`realname`/`format`/`notab`/`times` (player-account &
  report-format settings), and `resend`/`lore`/`players`/`public` (immediate secondary effects — re-mail a
  past report, request a lore sheet, the player list). These configure the **account and the scan reply**,
  not the per-turn world simulation — a seam §10.8 calls out. Changing `email`/`vis_email`/`password` mails
  a confirmation to **both** the old and new addresses as a security measure. ✅

### 10.7 Player-supplied names are untrusted → sanitized at ingest ✅ / 🟡 (charset)

*(Promoted from the early decision recorded ahead of this pass.)*

- Players name nobles and other createable entities through orders (`NAME`, `FORM`, `BANNER`, ship names in
  `BUILD`, §9.2). Because the order file is the **untrusted-input boundary** (§10.2), a name carries
  whatever bytes a player typed and could otherwise inject markup/script (XSS) or terminal/PDF control
  sequences when embedded in a report (text, PDF, HTML view). ✅
- **Mechanism:** names are **sanitized at ingest**, in the orderfile adapter — so only a safe, typed name
  reaches `domain` and **every render target (§12) inherits a clean string**. Render-time escaping, if
  added, is defense-in-depth, not the primary control. The exact **allowed character set and transform**
  (reject vs. strip vs. escape on the way in) is 🟡 — fixed with the §10.1 tokenizer spec. ✅ (boundary) /
  🟡 (charset)

### 10.8 Architectural implications

The architectural consequences of §10 have moved to their correct homes per the routing rule in
[AGENTS.md](AGENTS.md); this section remains only as a pointer (§11.9's sibling list references its
anchor). Unlike the consumer chapters §3–§9, §10 *owns* the order-ingest machinery — but the
`OrderSource` port and the `internal/infra/orderfile/` untrusted boundary were already decided, so all
but one consequence route to homes that already hold them:

- **The order file is the engine's one untrusted boundary, and it lives in `orderfile/`.** Raw bytes
  exist *only* inside `internal/infra/orderfile/`, which parses the DSL, rejects malformed input with
  `cerr.ErrInvalidOrders`, and emits typed `domain.OrderBundle` values through the `OrderSource` port
  (`ReadOrders → []domain.OrderBundle`, one per player); `app` and `domain` never see text. This is the
  standing untrusted-input rule in [AGENTS.md](AGENTS.md) ("Input files are untrusted"); the port's
  descriptive signature and its `ErrInvalidOrders` return are already in
  [reference/ports.md](docs/content/reference/ports.md) and
  [reference/errors.md](docs/content/reference/errors.md).
- **Two-stage validation *is* the layer boundary.** Lexical/structural validity (verb exists, envelope
  well-formed, scan-order params) is the **parser's** job and fails as `ErrInvalidOrders`; semantic
  validity (skill offered here? peasants present? target real?) is **resolution's** job (§11) and fails as
  the business-meaning `cerr` sentinels at run, taking zero time. The descriptive error split lives in
  [reference/errors.md](docs/content/reference/errors.md); the rationale — the parser is the boundary, so
  the domain never re-validates well-formedness — is in
  [explanation/use-cases.md](docs/content/explanation/use-cases.md).
- **The command catalog is authored reference data with two consumers.** Each verb's priority and time
  class are immutable authored facts (like the item/skill/ship-type tables); the **parser** reads arity and
  parse-time params, the **§11 scheduler** reads priority and time. *Where* the catalog physically lives —
  a domain constant table vs. an authored artifact loaded like the map (§2.1) — is the **same open
  authored-data artifact-and-loader concern** as the "Map artifact format" row and the planned `MapSource`
  port in [`docs/adr/`](docs/adr/README.md) (a sibling loader, not a separate decision — exactly as the
  item-type and skill tables route in §4.9/§7.9). The descriptive catalog model belongs in
  `reference/model/` with the deferred §10 page (see the note below).
- **Account/scan directives may need a channel separate from the `OrderBundle`.** `begin`/`unit`/`end`,
  the account/report-format settings (`password`/`email`/`vis_email`/`realname`/`format`/`notab`/`times`),
  and the immediate-effect directives (`resend`/`lore`/`players`/`public`) configure the *account and the
  scan reply*, not the simulated turn (§10.6) — so they probably should **not** ride inside the per-unit
  `OrderBundle` the §11 resolver drains. Whether the parser emits a second typed value (an
  "account-directives" struct alongside the bundle) is an **open `OrderSource`-output shape**, now tracked
  in [`docs/adr/`](docs/adr/README.md)'s open-decisions register; settle it when the orderfile adapter and
  the `OrderSource` port are built.
- **`OrderBundle` + an input hash is the idempotency key's raw material.** Hashing the accumulated, parsed
  bundles for `(gameID, turn)` gives `TurnLedger` its input hash — the Application-level idempotency
  concern whose rationale already lives in
  [explanation/idempotency.md](docs/content/explanation/idempotency.md) and
  [explanation/use-cases.md](docs/content/explanation/use-cases.md). *Which* bundles accumulate by the
  deadline depends on mail arrival — the `UNIT`-replace race the rulebook's `PASSWORD` trick exploits — so
  arrival ordering is an **ingest-time** input to the hash, deterministic *given* the assembled mailbox (a
  §12 mail-ingest detail carried with the report pass); the resolver downstream stays a pure function of
  that hash.
- **The 250-order cap and its silent drops want an explicit ingest warning.** Truncation that reads as
  "accepted everything" should instead surface the dropped orders in the acknowledgement reply (§12) rather
  than discarding them quietly — a reporting-behavior constraint the §12 acknowledgement-reply mechanic
  owns.

> **Not yet distilled.** Like §2–§9, §10's decided facts (the `begin`/`unit`/`end` DSL and its
> forgiving grammar, the untrusted boundary and `OrderBundle` flow, the parse-time/execute-time split,
> the priority/time catalog rules, the per-unit queue and `STOP` model, the player-entity admin set, and
> name sanitization at ingest) wait on **§11 (turn resolution)** before promotion to a `reference/model/`
> page — the scheduler that consumes priorities, the `STOP`/interrupt and time-accrual timing, and the
> idempotency hashing may still reshape the slots. The **tokenizer/grammar spec** and **sanitization
> charset** (§10.1/§10.7) remain 🟡 and are carried forward. The two *architecture* shapes §10.8 surfaced
> — the **physical home of the command catalog** and whether **account/scan directives** travel as a
> separate typed channel out of the parser — have moved to their home in
> [`docs/adr/`](docs/adr/README.md) (the authored-data concern and the open `OrderSource`-output row,
> respectively).

## 11. Turn resolution 🟡

§11 is the engine's **global sequencer** — the section every earlier one deferred its *timing* and
*ordering* questions to. §6.1 deferred *when* tax is computed, upkeep debits, and the castle owner is
paid; §6.2 fixed upkeep "at end of month" but left its slot relative to combat; §8.2 deferred *when* a
battle resolves relative to movement, market clearing, and pillaging; §9.4 deferred the ship travel-time
model; §10.4/§10.5 handed over the **scheduler** that consumes each command's priority and the moment a
queued `STOP` bites. §11 is where those compose into one deterministic month. It **owns**: the **turn
clock** (30 game days, §11.1), the **three turn phases** (§11.2), the **priority scheduler** (the 0–4
bands §10 assigns, run ascending, location-order tiebreak — §11.3), **parallel execution & time accrual**
(§11.4), **`STOP`/interrupt timing and carry-over** (§11.5), the **turn-boundary events** (new-player
addition, `FORM` minting, NP award, the monthly money flow — §11.6), the **determinism contract** (no
`time.Now`, no ambient randomness; a seeded PRNG, §11.7), and the **idempotency placement** (`ProcessTurn`
+ an input hash, §11.8). It **reuses, never re-decides** every *mechanic*: combat math is §8, travel cost
§2.4/§9.4, upkeep amounts §6.2, tax base §5.2/§6.1, study/production §6/§7 — §11 fixes only **when each
fires and in what order**, never what it computes. Primary sources: [playing.md](docs/content/rules/playing.md)
(the parallel-execution and 30-day-month model, the turn schedule, global-only re-runs) and
[orders.md](docs/content/rules/orders.md) (the priority/time table, the `STOP` interrupt rule). §11 is the
**`ProcessTurn`** use case — the `process` pipeline stage (`internal/app`, see the idempotency and
use-case explanation docs). Cross-refs: §2.9 (the established turn-seed discipline), §3.7 (NP at multiples
of 8), §6.1/§6.2 (money-flow & upkeep timing), §8.2 (battle placement), §9.4 (ship travel), §10.4/§10.5
(priority bands & the per-unit queue), §12 (the report the resolved turn feeds).

### 11.1 The turn is a fixed 30-day month ✅

- One turn = one Olympian month = **30 game days** ([playing.md](docs/content/rules/playing.md)). The day
  is the unit of accrual: each command consumes a **time class** (§10.4) — `0 days`, a fixed count,
  `as given` (player-supplied), or `varies` (computed from world state). ✅
- The real-world cadence (turns run Mondays at noon AEST/AEDT) is **operator scheduling, outside the
  engine**. Resolution never reads a wall clock (§11.7): it knows only the integer **turn number** and the
  30-day budget. The turn number drives the calendar/season (the eight-month Olympian year,
  [playing.md](docs/content/rules/playing.md)) and the multiple-of-8 NP award (§3.7/§11.6). ✅

### 11.2 The turn's three phases: open → execute → close ✅ (frame) / 🟡 (sub-step placement)

A turn resolves in three ordered phases:

1. **Open** — turn-boundary *entry* events that must precede order execution: **new players are added**
   ([playing.md](docs/content/rules/playing.md): "additions are performed when the turn is run"), the
   **`FORM`-number reservation list** is computed and published (§3.2/§10.5), and **NP is awarded** on
   turns that are a multiple of 8 (§3.7).
2. **Execute** — the 30-day order-execution loop driven by the priority scheduler (§11.3) over every
   unit's command queue (§11.4).
3. **Close** — end-of-month *exit* bookkeeping: the **monthly money flow** (tax base → garrison upkeep →
   castle owner, §6.1), **men upkeep** charged to holding nobles with deterministic desertion/starvation
   (§6.2), then hand-off to **report rendering** (§12).

The **phase boundaries are decided**; the **exact slot of several sub-steps inside Execute vs. Close is
🟡**. §6.1 explicitly deferred *when* tax is computed and paid; §6.2 fixed upkeep at month's end (→ Close);
§8.2 deferred where battles fall relative to movement/market/pillage (→ within Execute). §11 names the
phases — pinning every sub-step's slot is the remaining work this section tracks.

### 11.3 The priority scheduler ✅ (bands & tiebreak) / 🟡 (tick granularity)

- Within **Execute**, commands run in **ascending priority** (§10.4): all ready **prio-0** (permission /
  attitude — `ADMIT`, `HOSTILE`, `NEUTRAL`, `DEFEND`, `CONTACT`, `DECREE`, `DEFAULT`) before any
  **prio-1** (zero-time commands and `WAIT`), before **prio-2** (`MOVE`/`FLY`/`BOARD`), before **prio-3**
  (everything else), before **prio-4** (`SAIL` alone). The bands exist so a declared attitude is in force
  **before** the movement and combat it governs, and so sailing settles last. ✅
- **Same-priority ties resolve by location order** — the "seen here" list, oldest-resident-first (§10.5);
  a freshly-arrived or `UNSTACK`ed unit reinserts per the §12 ordering rules. This is the **sole
  tiebreaker**, making the within-band order a deterministic function of recorded state — never `rand`,
  never Go map-iteration order (§6.8). ✅
- **Tick granularity (🟡):** the rulebook presents execution as happening "in parallel"
  ([playing.md](docs/content/rules/playing.md)) — player-facing fiction. The engine must impose a
  **deterministic total order**. Whether that is a literal **day-by-day loop** (advance every unit one
  day, re-evaluate ready commands each day, in priority/location order) or an **event-merge** (compute each
  command's completion day, then process events in `(day, priority, location)` order) is the **central open
  modeling choice** — it dictates exactly how movement, combat (§8.2), and market clearing (§6.6)
  interleave mid-month. The bands and tiebreak are fixed; the tick model is deferred.

### 11.4 Parallel execution & time accrual ✅ / 🟡

- Each unit drains its queue **independently**: a command starts, consumes its time class in game days, the
  next begins, "as many orders as game time allows" until the unit's 30-day budget is spent
  ([playing.md](docs/content/rules/playing.md)). ✅
- **Failed orders generally take zero time** (§10.3/§10.4): a semantically-invalid command (target gone,
  skill not offered here, no peasants to recruit) fails with a business-meaning `cerr` sentinel at run,
  consumes **no days**, and the queue advances to the next order. ✅
- **`varies`/`as given` durations are evaluated at start-of-command against current world state** — travel
  cost (§2.4/§9.4), `BUILD` (§6.4), `USE` (§7) — never precomputed once for the whole turn. The
  cost-functions live in their owning sections; §11 fixes only that they read **the state as it stands when
  the command begins**, deterministically. 🟡 where those functions are themselves still 🟡 (the §2.4
  movement-variance model, the §9.4 ship travel-time model). ✅ (the *when*) / 🟡 (the *how much*)

### 11.5 `STOP`, interrupts & carry-over ✅

- A re-submitted `unit` block **does not interrupt** an in-flight command **unless its first queued order
  is `STOP`** (§10.5). `STOP` queues like any other order and **bites only when the turn runs** — so a
  later file can still replace it before the deadline. At run, before draining a unit's queue, the
  scheduler checks for a leading `STOP`: present → **abandon the in-flight command** (its accrued days are
  forfeit — no partial credit) and proceed; absent → let the in-flight command **finish first**, then
  drain the new queue. ✅
- **Carry-over:** a command still executing when the 30-day budget is exhausted **continues into the next
  turn** ([playing.md](docs/content/rules/playing.md)); its **remaining days are preserved as in-flight
  state in the turn snapshot** and it resumes at the start of the next turn's Execute phase. Unstarted
  queued orders simply remain queued. This is why **per-unit in-flight progress is part of persisted turn
  state**, not transient scheduler memory (§11.9). ✅

### 11.6 Turn-boundary events: additions, minting, NP & money flow ✅ / 🟡

Collects the entry/exit events earlier sections deferred to the turn boundary (§11.2):

- **New-player addition** (Open): players who joined since the last run enter at run, placed in a
  safe-haven city (§2.8, where combat/magic are forbidden), seeded with starting gold and NP **including
  catch-up NP** so all players hold roughly equal NP (§3.7). ✅
- **`FORM` minting** (Open + Execute): the next-N entity numbers are reserved and **published in the report**
  (§3.2/§10.5) so a player can pre-queue a `unit <minted-number>` block; the `FORM` itself (prio-3, 7 days)
  runs in Execute and the newborn begins draining its queue the moment it appears. ✅
- **NP award** (Open): **+1 NP on every turn that is a multiple of 8** (§3.7). ✅
- **Recorded countdown & replenishment timers** (Close): the monthly tick that advances every
  deterministic per-turn counter the earlier sections deferred here — **loyalty-bond decay** (contract,
  fear, oath; §3.5/§3.6), **body decomposition and noble-return** windows (§4.5/§3.6), **relic/realm and
  pillage-recovery** timers (§5.2/§5.8), and **illness recovery / health regeneration** (§8). All advance
  by **fixed rule or §2.9-seeded draw**, never live entropy. ✅ (that they tick in Close) / 🟡 (per-timer
  slot order, with their owning sections)
- **Money flow** (Close): the tax base is computed, garrison upkeep debited, the castle owner paid
  **⌊(base − garrison upkeep)/2⌋**, the un-forwarded half lost (§6.1); men upkeep is charged to holding
  nobles, and **one-third of unpaid soldiers leave / unpaid peasants starve**, the affected men selected
  **deterministically from recorded state** (§6.2), never by `rand`. ✅ (the events) / 🟡 (their exact slot
  inside Close, carried with §6.1)
- **Un-garrisoned castle collection — settles the §6.1 pin:** a castle with **no functional garrison**
  (≥10 men, §5.5) in its own province is **undefended, and its tax goes uncollected** that month —
  collection is gated on the same garrison presence that gates protection. This **closes the §6.1 open
  item** (which pinned the question here for confirmation). ✅

### 11.7 Determinism: the seeded PRNG and the no-ambient-input rule ✅ (rule) / 🟡 (PRNG integration)

- Resolution is a **pure function of recorded state, the assembled orders, and a recorded seed** —
  conceptually `resolve(state, orders, seed) → state'`. CLAUDE.md forbids `time.Now` and ambient
  randomness in the domain; AGENTS.md's **`Clock` port** abstracts time. So **no wall clock and no
  `math/rand` global** enters resolution: the turn number, not a date, drives the calendar (§11.1). ✅
- **Two distinct flavors of "the engine decides", both reproducible:**
  - **Deterministic-from-state selection** — *which* men desert/starve when a noble underpays (§6.2),
    *which* of several equal trades matches (§6.6), the within-band scheduling tie (§11.3). These follow a
    **fixed rule over recorded order**, with **no PRNG draw at all**.
  - **Genuinely stochastic mechanics** — combat hit rolls and random target selection (§8.2), the 25%/week
    skill-acquisition roll (§7.3), randomized return/decay windows (§4.5/§5.8), exploration/quest outcomes.
    These draw from the **seeded PRNG**.
- **Randomness is a port, realized now.** `app/ports.go` declares an **`RNG` port** (`Roll`, `RollDice`) —
  "a deterministic random source used by use cases that involve dice or stochastic decisions" —
  implemented by the **`internal/infra/prng`** adapter, a **splittable PCG** that marshals its state. This
  mirrors the **`Clock` port**: stochastic draws are a **use-case (app) concern reached through a port**,
  so the **domain imports no RNG and no entropy source**, and the SOUSA domain-import check needs no
  exception. **RNG state persists in the game-state snapshot, round-tripped by `GameStateStore`** — *not*
  the `TurnLedger`, whose key stays `(gameID, turn, inputHash)` — and the port deliberately hides
  marshal/scan. **Substream `Split()`** (per game / stage / player, for order-independent reproducibility —
  a battle's stream does not shift because an unrelated unit rolled earlier) is a **Runtime wiring-time**
  operation on the concrete adapter today, promoted to the port only if a use case needs mid-turn fan-out.
  🟡 remains on the **seed-derivation rule** and the **substream-assignment scheme**. ✅ (the port & its
  realization) / 🟡 (seeding/splitting policy)
  > **Reconciles §2.9 / §2.4 ✅:** those sections, written before the ports existed, placed the seed "in
  > the `TurnLedger`" and had "the domain derive variance via a seeded PRNG." The implemented ports refine
  > both — RNG **state lives in the snapshot (`GameStateStore`)**, and draws are a **use-case concern
  > through the `RNG` port**, never a domain import. The determinism *intent* (recorded state, no live
  > entropy) is unchanged; only the mechanism is now concrete.

TODO: The note above is questionable. Seeds may live in the GameStateStore, but they have to be carried in the TurnLedger for deterministic restarts. I am concerned about the game design capturing implementation details. We need to circle back on both. 🟡

### 11.8 Idempotency: `ProcessTurn`, the input hash & global-only re-runs ✅

- `ProcessTurn` is the state-mutating use case, and **idempotency is an Application concern** (CLAUDE.md,
  the idempotency explanation doc): it hashes `(prior state, validated orders)` — and the **prior state
  already includes the RNG state** carried in the snapshot (§11.7), so the random stream is part of the
  hashed input without a separate seed field — asks
  `TurnLedger.AlreadyProcessed(gameID, turn, inputHash)`, **short-circuits** with
  `cerr.ErrTurnAlreadyProcessed` on a match, otherwise resolves and `Record`s the hash. Re-running after an
  operator fixes garbled orders **changes the hash**, so it does **not** short-circuit — exactly the
  intended escape hatch. ✅
- **Re-runs are global, never per-player** ([playing.md](docs/content/rules/playing.md): "it is not
  possible to re-run a turn for a single player … only for all players"). This matches the hash model
  directly: the hash is over the **whole assembled mailbox plus prior state**, so the idempotency unit is
  the **turn**, not the player. ✅
- Because **RNG state travels inside the snapshot** (§11.7), a re-run from the same prior state replays the
  same stochastic draws automatically — there is no separate seed to forget to hash, and an idempotent
  re-run cannot silently diverge. ✅
- Mail-arrival ordering feeds the `UNIT`-replace race **at ingest** (§10.8); by the time `ProcessTurn`
  runs, the bundle set is fixed and resolution is pure. **Late orders simply queue for the next turn**
  ([playing.md](docs/content/rules/playing.md)) — they never retroactively alter a resolved turn. ✅

### 11.9 Architectural implications

The architectural consequences of §11 have moved to their correct homes per the routing rule in
[AGENTS.md](AGENTS.md); this section remains only as a pointer (§8.10/§9.8/§12.9's sibling lists and
several body cross-refs reference its anchor). §11 carries the determinism and idempotency seams the
whole engine turns on, but earlier passes already settled them, so every consequence routes to a home
that already holds it:

- **`ProcessTurn` is the `process` stage and must stay a pure transform.** It takes the prior snapshot
  and the turn's `[]OrderBundle` and returns the next snapshot; it reads the world only through
  `GameStateStore`, consults `TurnLedger` for idempotency, and touches no renderer, mailer, or clock —
  if resolution ever "needs" `time.Now`, live entropy, or a concrete store, the boundary moved to the
  wrong layer. That it is a pure sequential transform (turn N's snapshot is turn N+1's input) is the
  concurrency decision in [ADR 0002](docs/adr/README.md); the use-case shape and its ports-not-adapters
  discipline are described in [explanation/use-cases.md](docs/content/explanation/use-cases.md).
- **Randomness is a port, like time — decided.** `app/ports.go` declares the **`RNG` port**
  (`Roll`/`RollDice`), realized by the **`internal/infra/prng`** seeded-PCG adapter, exactly mirroring
  `Clock`; the domain imports no concrete RNG, so the domain-import conformance command passes with no
  relocation and no domain-side `Roller` interface. This is [ADR 0003](docs/adr/README.md), and the
  port's descriptive signature now lives in [reference/ports.md](docs/content/reference/ports.md). The
  residual open points — the seed-derivation rule, the substream-assignment scheme, and whether
  domain-resident combat math (§8.2) rolls in the use case or through an injected domain interface — are
  recorded with [ADR 0003](docs/adr/README.md).
- **RNG state lives in the snapshot, so the idempotency hash captures it for free** (§11.7/§11.8):
  because `GameStateStore` round-trips RNG state as part of game state (an
  [ADR 0003](docs/adr/README.md) consequence, pinned among the State-storage snapshot constraints in
  [`docs/adr/`](docs/adr/README.md)), hashing `(prior state, orders)` already pins the random stream —
  there is no separate seed field to forget, and a re-run from the same prior state reproduces the same
  draws. The `(state, orders)` input-hash rationale is in
  [explanation/idempotency.md](docs/content/explanation/idempotency.md).
- **In-flight command progress is persisted turn state** (§11.5 carry-over): the per-turn snapshot must
  carry each unit's *remaining days on the running command*, not just its pending queue, or a turn
  cannot resume a multi-day order. This is **already pinned** among the State-storage snapshot
  constraints in [`docs/adr/`](docs/adr/README.md) ("per-unit in-flight command progress").

The remaining items §11 surfaced are **game-design timing questions, not engine-architecture
decisions**, so they stay tracked in this chapter's body rather than moving to `docs/adr/`:

- **The intra-turn tick model** — day-by-day loop vs. event-merge (§11.3) — is the largest undecided
  mechanic; it gates how movement, combat (§8.2), market clearing (§6.6), and the money-flow/upkeep
  slots (§6.1/§6.2) interleave mid-month. It is a game-time modeling choice, not an infra-adapter
  choice, so it remains a 🟡 in §11.3.
- **Money-flow and upkeep slot placement** in the Close phase (§11.6, with §6.1/§6.2) is the last
  turn-timing decision; the same recorded state must always debit and pay in the same order. It remains
  a 🟡 carried with those sections.

> **Not yet distilled.** §11's decided facts (the 30-day clock, the open→execute→close phase frame, the
> ascending-priority/location-order scheduler, parallel time-accrual with zero-time failures, the `STOP`
> bite and carry-over of in-flight progress, the turn-boundary event set, the seeded-PRNG determinism
> contract, and the `ProcessTurn`/input-hash/global-re-run idempotency model) wait on **§12 (turn reports)**
> before promotion to a `reference/model/` page — the report is the resolved turn's only output, and the
> "seen here" / location-ordering rules the scheduler's tiebreak depends on are co-owned with §12. Still
> 🟡 and carried forward: the **intra-turn tick granularity** (§11.3), the **exact Close-phase slot** of
> money-flow vs. upkeep (§11.6, with §6.1/§6.2), the **PRNG seed-derivation and substream scheme** plus its
> **package location** (§11.7/§11.9, recorded with ADR 0003), and the §2.4/§9.4 cost-functions the accrual
> reads (§11.4).

## 12. Turn reports 🟡

§12 is the engine's **output projection** — the resolved turn's only player-visible product, the thing
§11.2's **Close** phase hands off to once the month is resolved. Where §11 owns *when each mechanic fires*,
§12 owns *how the result is shown to each player*. It **owns**: the **per-player perspective** that decides
what a faction may see (§12.2), the **anatomy** of the report — header, per-noble narrative, per-location
blocks, order template (§12.3), the **location-ordering rule** the "Seen here" list obeys (co-owned with the
§11.3 scheduler tiebreak — §12.4), the **formatting contract** for the canonical text product (80 columns,
2-space tab stops — §12.5), the **three delivery products** (text, PDF, JSON — §12.6), the
**always-generate-and-store** durability rule (§12.7), and **delivery to the registered address** with the
render/dispatch split (§12.8). It **reuses, never re-decides** every *mechanic* whose output it shows:
visibility into inner locations is §2.5, hidden-route disclosure §2.6, the prisoner-opacity rule §8.7, the
market report §6.6, the inventory/skills listings §7, the "next N nobles to be formed" list §3.2/§11.6.
§12 fixes only **what is projected, in what order, and how it is formatted** — never what the underlying
numbers are. Primary sources: [playing.md](docs/content/rules/playing.md) (report length & cadence, the
"Seen here" block, the Olympian calendar), [geography.md](docs/content/rules/geography.md) (the location-report
anatomy — routes, inner locations, skills taught, ships docked, market report), [orders.md](docs/content/rules/orders.md)
(the location-order rule, command precedence, the order template, the FORM-reservation line),
[markets.md](docs/content/rules/markets.md) (the market report), [combat.md](docs/content/rules/combat.md)
(prisoner display & the column-wrap example), [skills-magic.md](docs/content/rules/skills-magic.md)
(inventory and skills-taught listings). §12 is the **`RenderReports`** use case (the `render` pipeline stage)
feeding **`DispatchReports`** (the `dispatch` stage) — two stages, two ports, per AGENTS.md and CLAUDE.md.
Cross-refs: §2.5/§2.6 (visibility & hidden routes), §3.2/§11.6 (FORM reservation in the header), §3.7/§11.1
(NP and the calendar/month name), §6.6 (market report), §8.7 (prisoner opacity), §10.5/§11.3 (the per-unit
queue & location-order tiebreak), §11.2 (Close → render), §11.8 (idempotency & global-only re-runs).

### 12.1 The report is the resolved turn's only output ✅

- After Close (§11.2) the resolved snapshot is fixed; `RenderReports` produces **one report per player** from
  it. Rendering is a **pure, deterministic function of the resolved snapshot** — `render(snapshot, playerID)
  → PlayerReport` — reading no wall clock and drawing no randomness (the same no-ambient-input contract as
  §11.7). The header's month and season come from the integer **turn number**, not a date (§11.1: the
  eight-month Olympian year, [playing.md](docs/content/rules/playing.md)). ✅
- **Two distinct steps, two layers.** Projecting the snapshot into a per-player **`domain.PlayerReport`** —
  applying visibility (§12.2), ordering locations (§12.4), assembling narrative and listings — is a **pure
  domain transform** (testable, deterministic, no I/O). Turning a `PlayerReport` into bytes is **infra**
  (`ReportRenderer`, §12.6). The split matters: *what a player may see* is a **game rule that lives in
  domain**, never a decision a formatter makes (§12.9). ✅
- Reports are sized for play, not the engine: a three-noble starter report runs ~5 pages of 66 lines,
  typical reports 15–25 pages ([playing.md](docs/content/rules/playing.md)). Length is an output of faction
  size, not a configured limit. ✅

### 12.2 Per-player perspective: the report is what your faction can see 🟡

The report is **not** a world dump — it is assembled strictly from what the player's faction can observe:

- **Locations the faction occupies, and what is visible from them.** A unit reports its immediate location
  and the surrounding province per the §2.5 inner-location visibility rules: a unit in a sub-location sees
  that sub-location and its immediate surround, **not** into sibling inner locations, and an outer-province
  unit cannot see into an inner location without entering (§2.5, [geography.md](docs/content/rules/geography.md)). ✅
- **Other characters appear only as much as visibility allows** — the "Seen here" block lists co-located
  units (§12.4), but a foreign noble's owning faction is not disclosed (hence the §forwarding service keyed
  on entity number, [playing.md](docs/content/rules/playing.md)). The depth of detail shown for foreign units
  vs. the player's own is **🟡** (own units get full inventory/skills/queue; foreign units get name, banner
  text, and visible men only). 🟡
- **Prisoners contribute nothing.** A captured unit reports neither location nor sightings to its own
  faction; it shows in the holder's report marked `prisoner`, "little else" (§8.7,
  [combat.md](docs/content/rules/combat.md)). The faction owning the prisoner sees only that the unit is
  held. ✅
- **Hidden routes are disclosed per-faction.** A hidden route appears in a location's route list **only for
  factions that have traversed it** (or stacked across it); knowing the destination's entity number is not
  enough to see or use it (§2.6, [geography.md](docs/content/rules/geography.md)). ✅

### 12.3 Anatomy of the report 🟡 (section roster) / ✅ (the blocks)

The report is a sequence of blocks; the **individual blocks are well-attested in the rulebook**, their exact
top-level **roster and ordering is 🟡**. The working structure:

1. **Header** — turn number, Olympian month & season (§11.1), faction summary (gold, NP and any catch-up NP
   awarded this turn — §3.7/§11.6), and the **FORM reservation line**: `The next five nobles formed will be:
   …` (§3.2/§11.6, [orders.md](docs/content/rules/orders.md)) so a player can pre-queue a `unit
   <minted-number>` block. ✅
2. **Per-noble narrative** — for each of the player's own units, an **echo of the orders run with their
   outcomes**, in execution order. Echoed input lines are prefixed `>`; engine responses follow on their own
   lines (e.g. `> buy 79 5 10` / `Try to buy five iron [79] for 10 gold each.` — §6.6,
   [markets.md](docs/content/rules/markets.md)). A failed order reports its business-meaning outcome and (per
   §11.4) consumed no time. Own-unit detail includes **Inventory** and **skills** listings in the rulebook's
   tabular form (§7, [skills-magic.md](docs/content/rules/skills-magic.md)). ✅
3. **Per-location reports** — one block per location the faction can see (§12.2), each in the
   [geography.md](docs/content/rules/geography.md) shape: the location line (`Plain [ae48], plain, in region
   Tollus, civ-1`), **Routes leaving** (with hidden routes per §2.6), **Inner locations**, **Skills taught
   here** (§7.x), **Seen here** (§12.4), **Ships docked at port** (§9), and the **Market report** (§6.6). A
   block omits sub-sections that are empty. ✅
4. **Order template** — at the **bottom** of the report, a template listing every unit in the faction with
   its still-pending queued orders, ready to edit and resubmit ([orders.md](docs/content/rules/orders.md)).
   Carry-over in-flight commands (§11.5) surface here as still-running. ✅

### 12.4 Location order & the "Seen here" list ✅

The order in which units appear within a location is **recorded state**, governed by one rule set — and it is
the **same ordering the §11.3 scheduler reads for its same-priority tiebreak**. §11.3 and §12 are co-owners;
this is the single definition both consume:

- **Longest-resident first.** A unit entering a location is **appended to the end**; the unit present longest
  sorts to the top ([orders.md](docs/content/rules/orders.md)). ✅
- **Leave and return → back of the line.** A unit that departs and later returns is re-appended at the end. ✅
- **Unstack reinserts after the parent**, not at the end: a unit unstacking from beneath another appears
  **immediately after that unit**, preserving locality ([orders.md](docs/content/rules/orders.md)). ✅
- **New player characters join at the top** of their safe-haven's list (not the bottom); nobles they later
  `FORM` are appended at the bottom as usual ([orders.md](docs/content/rules/orders.md)). ✅
- Stacking is shown by indentation and `accompanied by:`; this ordering is a deterministic function of
  recorded arrival/unstack events — **never `rand`, never map-iteration order** (§6.8). Because the scheduler
  tiebreak (§11.3) and the report's "Seen here" block are the **same list**, location-order precedence (e.g.
  who `HARVEST`s first — [orders.md](docs/content/rules/orders.md)) is exactly what the player reads. ✅

### 12.5 Formatting contract: the canonical text product ✅

The text report is the canonical human format; its layout is **infra (the text renderer's contract)**, not a
domain concern. Standardized for opyl:

- **80-column width.** The rulebook's writer wraps at column 79; opyl assumes the same 80-column field — not
  a hard limit of the medium, but it reads well and keeps reports diff-stable. ✅
- **Automatic wrap with aligned continuation.** A line too long to fit wraps onto continuation lines indented
  to align under the wrapped content — as in the prisoner example, where a long "Seen here" entry wraps and
  the continuation sits under the name ([combat.md](docs/content/rules/combat.md)). ✅
- **2-space tab stops, standardized.** The rulebook examples mix 2- and 4-space indents; opyl uses a **single
  2-space indent unit** for every nesting level (routes under a location, stacked units under their parent,
  listings under their heading). ✅
- **Entity codes in brackets** trail every name (`Osswid the Destroyer [5499]`, `Gold [1]`, `City of the
  Lost [gx14]`), per [playing.md](docs/content/rules/playing.md). ✅
- The **PDF** product is a typeset rendering of the same `PlayerReport` and is **not** column-bound; the
  **JSON** product (§12.6) is structured data and carries no layout at all. The 80-column / 2-space contract
  binds the **text renderer only**. ✅

### 12.6 Three delivery products: text, PDF, JSON ✅ (text/PDF) / 🟡 (JSON shape)

All three are **renderings of one `domain.PlayerReport`** behind the **`ReportRenderer`** port (turn data →
bytes + MIME type); none re-derives game state, and adding a format is an infra change only (AGENTS.md):

- **Text** — the canonical human report (§12.5), `text/plain`. ✅
- **PDF** — typeset human report, `application/pdf`; the PDF library choice stays an open infra decision
  (AGENTS.md). ✅
- **JSON** — **machine-readable turn results, emailed alongside the human report** (a modern addition), MIME
  `application/json`. It is the most faithful serialization of `PlayerReport`, for players who script against
  their results. Its concrete schema is **🟡** — it should be a stable projection of `PlayerReport`, versioned
  so format changes do not silently break consumers. 🟡
- A future **SQLite export** of results is **explicitly out of scope** — it depends on too many unsettled
  factors to design now; noted only so the JSON schema is not over-fitted to email. ❓ (deferred)

### 12.7 Reports are always generated and stored 🟡

A modern departure from re-rendering on demand:

- **Every turn's reports are rendered and persisted** when the turn resolves, in all delivered formats, so a
  later code change can **never alter a past turn's results** — the stored report is the player's record of
  what happened, frozen against renderer evolution. ✅
- **The resolved snapshot remains the source of truth**; a stored report is a reproducible projection of it
  (§12.1). This is why storage is safe: regenerating from the same snapshot with the same renderer yields the
  same bytes. ✅
- **The GM can remove and regenerate a bad report.** Regeneration re-renders from the stored snapshot with
  current renderer code — a deliberate operator act, distinct from re-running the turn (§11.8). It does **not**
  re-resolve the month; only a global turn re-run (§11.8) changes results, and that regenerates **all**
  players' reports. ✅
- Persisting rendered artifacts is a capability the current port set does not cover — it implies a new
  **report-store port** (§12.9). 🟡

### 12.8 Delivery: the registered address, render and dispatch kept separate ✅

- Reports and JSON results are **sent to the player's registered email address** — identity is **routing
  metadata** (`domain.PlayerID` → email), not a security principal (AGENTS.md operator-trust model). **Updating
  a player's email address is out of scope** for the engine. ✅
- **Render and dispatch are two ports, never one** (CLAUDE.md/AGENTS.md): `ReportRenderer` produces bytes;
  **`ReportDispatcher`** sends an attachment to a recipient. `DispatchReports` wires them, so a **dry run —
  render and store without sending — is free** (and is exactly what §12.7's always-store path does even when
  dispatch is skipped). ✅
- **`DispatchReports` is idempotent** (AGENTS.md): re-invoking for the same `(gameID, turn)` must not double-send.
  Like `ProcessTurn` it short-circuits via `TurnLedger` + an input hash (here over the stored report set), so an
  operator rerun of the dispatch stage is safe. ✅
- The mail transport (SMTP vs. SES/SendGrid vs. drop-EML-to-`/outbox`) stays an open infra decision behind
  `ReportDispatcher` (AGENTS.md). ✅

### 12.9 Architectural implications

These follow from §12 and join §2.9 / §3.8 / … / §11.9 in AGENTS.md's "Open architectural decisions" table:

- **Visibility is a domain rule, not a renderer responsibility.** Building the per-player `domain.PlayerReport`
  — fog of war (§12.2), prisoner opacity (§8.7), per-faction hidden-route disclosure (§2.6) — is a **pure
  domain projection** of the resolved snapshot. A `ReportRenderer` receives an **already-filtered**
  `PlayerReport` and only formats it; it must never hold "is this secret?" logic. If a formatter ever needs to
  decide what a player may see, the filter is in the wrong layer.
- **A new report-store port is needed** (§12.7). The current ports — `ReportRenderer`, `ReportDispatcher` —
  cover *make bytes* and *send bytes*, not *durably keep bytes*. Add a small port in `app/ports.go` (working
  name **`ReportStore`**: persist, retrieve, and remove rendered artifacts keyed by `(gameID, turn, playerID,
  format)`) before implementing an adapter. Whether stored reports live with the per-turn snapshots
  (`GameStateStore`) or in their own store interacts with the open **State storage** decision (SQLite vs.
  per-turn directory).
- **JSON is a third `ReportRenderer` format, not a new pipeline.** It plugs in behind the existing port
  (bytes + `application/json`); the render/dispatch split (§12.8) and the always-store path (§12.7) apply
  unchanged. The deferred SQLite export (§12.6) is a *different* capability and must not warp the JSON schema.
- **Location order is single-sourced and feeds two consumers** (§12.4): the §11.3 scheduler tiebreak and the
  report's "Seen here" block read the **same** recorded ordering. It must be derived from recorded
  arrival/unstack events in domain, never recomputed independently in the renderer — or the precedence a
  player reads could drift from the precedence the engine ran.
- **The order template closes the orders loop** (§12.3): the bottom-of-report template is the editable
  successor to this turn's queue (§10.5) including carried-over in-flight commands (§11.5), so the snapshot's
  per-unit pending-queue and in-flight state (§11.9) must be projectable back into the report.

> **Not yet distilled.** §12's decided facts (rendering as a pure projection of the resolved snapshot; the
> domain `PlayerReport` vs. infra `ReportRenderer` split; the per-faction visibility, prisoner-opacity and
> hidden-route rules; the single location-ordering shared with §11.3; the 80-column / 2-space text contract;
> the text/PDF/JSON product set; always-generate-and-store with GM regenerate; delivery to the registered
> address with the render/dispatch split) wait on the **report-store port** and the **JSON schema** before
> promotion to a `reference/model/` page (and a how-to for "switch the PDF renderer" / "regenerate a turn's
> reports"). Still 🟡 and carried forward: the **top-level section roster/ordering** of the report (§12.3),
> the **own-vs-foreign unit detail depth** (§12.2), the **versioned JSON schema** (§12.6), the **report-store
> placement** against the open State-storage decision (§12.7/§12.9), and the deferred **SQLite export** (§12.6).
> With §12 settled to here, §11's decided facts are now unblocked for joint promotion (§11.9's note).

## 13. Open decisions carried from AGENTS.md 🟡

§13 is the **reconciliation register** — it introduces no new mechanics. For each row of AGENTS.md's
"Open architectural decisions" table it records where the design work in §1–§12 now leaves the decision:
**decided**, **constrained but open**, or **untouched**. AGENTS.md remains the authoritative table; this
section is the design layer's verdict feeding back into it, and where a row is **decided** here the
ADR-style note belongs in AGENTS.md itself (per its own "Add a short ADR-style note … when each is
settled"). The §X.9 "Architectural implications" notes throughout this file are the raw material; §13
collates them against the six table rows and lists the decisions the design **surfaced** that the table
does not yet carry. Cross-refs are to the sections that did the deciding; nothing here overrides them.

### 13.1 State storage 🟡 (contents fixed, backend open)

SQLite **vs.** a directory of versioned JSON/YAML per turn — **still open**. The design has not chosen a
backend, but it has fully pinned what the backend must hold:

- The unit of persistence is the **per-turn snapshot**, and the snapshot is the **source of truth** the
  report is a reproducible projection of (§12.1/§12.7). It must round-trip: **RNG state** (§11.7), each
  unit's **in-flight command progress** for carry-over (§11.5/§11.9), the full set of **timer/countdown
  state** (decomposition & noble-return §4.9, relic/realm & pillage-recovery §5.9, opium/mine timers §6.8,
  loyalty-bond decay §3.8), and the **per-location arrival-order list** the scheduler tiebreak and "Seen
  here" block share (§6.8/§12.4). ✅ (what it must contain)
- Idempotency rides on top, not in the store: `ProcessTurn` hashes `(prior snapshot, validated orders)` and
  the **`TurnLedger`** keyed `(gameID, turn, inputHash)` is a **distinct** persistence concern (§11.8).
  Alongside sit two more stores the design surfaced — the **report store** keyed `(gameID, turn, playerID,
  format)` (§12.7/§13.7) and the authored **map artifact** read through `MapSource` (§2.1/§13.7). ✅ (the
  separate stores)
- The git-diffable, human-inspectable audit appeal cited for the authored map (§2.1) is the same appeal the
  per-turn-directory option carries, but **nothing in the rules forces it**; both backends can satisfy the
  contract above. The decision is genuinely open — the design constrains the **contents**, not the engine. 🟡

### 13.2 PDF library 🟡 (constrained, open)

`gofpdf` / `gopdf` **vs.** `typst` CLI **vs.** `chromedp` — **open**, an infra choice fully contained behind
`ReportRenderer` (§12.6). One design-side constraint has emerged: §12.7 stores reports and lets the GM
**regenerate** them, so reproducible, deterministic byte output (same snapshot + same code → same bytes) is
desirable — a point favoring a pure-Go, version-stable library over an external binary or headless browser
whose output can drift with environment. Not decisive, not decided. 🟡

### 13.3 Order file format ✅ (format) / 🟡 (spec)

**Decided** by §10.1 toward the rulebook's **custom line-oriented DSL** — not YAML, not a structured-field
email schema. The envelope (`begin <player> [password]` … `unit <number>` blocks … a single `end`), the
forgiving grammar (case- and whitespace-insensitive, `#` comments, quoted multi-word args), **`UNIT`-replaces-not-appends**,
and the 250-order-per-unit cap are fixed (§10.1). What remains 🟡 is the **exact tokenizer/grammar spec**
(quoting edge cases, numeric-vs-entity-code argument forms), pinned when the `internal/infra/orderfile/`
adapter is built. This is distinct from §13.4 — the DSL is the body, mail transport is its carrier. **→ ADR
note due in AGENTS.md.** ✅

### 13.4 Mail transport 🟡 (open)

Direct SMTP **vs.** SES / SendGrid **vs.** drop-EML-to-`/outbox` — **open**. It is the transport behind two
ports, not one: outbound through **`ReportDispatcher`** (§12.8) and inbound as the over-the-wire arrival of
order files feeding **`OrderSource`** (email/scp/dropbox, §10.2). The constraint is containment: the choice
must stay entirely inside those adapters, and `DispatchReports` **idempotency** (no double-send on rerun,
§12.8) lives in the **app** layer above transport, so it holds regardless of which transport lands. 🟡

### 13.5 CLI framework ✅ (design-neutral — defer freely)

stdlib `flag` **vs.** `cobra` — **carries no game-design stake.** The four pipeline stages are use cases
(`IngestOrders → ProcessTurn → RenderReports → DispatchReports`) and the CLI is a thin `internal/delivery/cli`
layer over them (CLAUDE.md); the framework choice touches neither domain nor app. The only design-side
requirement is structural and satisfied by either: **each stage invocable as its own subcommand**, plus a
`pipeline` subcommand running all four (CLAUDE.md/AGENTS.md). Decide at implementation time. ✅ (nothing here
blocks it)

### 13.6 Concurrency model ✅ (confirmed)

"Turns processed serially per game; multiple games in parallel" — **confirmed by the §11 determinism
contract**, not merely accepted:

- `ProcessTurn` is a **pure sequential transform** imposing a deterministic total order (§11.3), and turn
  N's resolved snapshot **is** turn N+1's input — so two turns of one game **cannot** overlap. ✅
- Distinct games share no state, so they may resolve **in parallel** with no coordination. ✅
- The **RNG substream `Split()`** (per game / stage / player, §11.7) gives order-independent reproducibility,
  so any future within-turn fan-out would stay deterministic — but a single turn's resolution adds **no
  goroutines** today; AGENTS.md's "confirm before adding goroutines" is answered: not inside a turn. **→
  confirm in AGENTS.md.** ✅

### 13.7 Decisions the design surfaced — to add to the table 🟡

§1–§12 raised infra decisions the AGENTS.md table does not yet list; they belong in it:

- **Map artifact format + `MapSource` port** (§2.1/§2.9) — the authored province graph is loaded by an infra
  adapter as immutable domain input; the on-disk format (JSON/YAML/custom) is undecided. 🟡
- **Report-store port + format** (§12.7/§12.9) — persist / retrieve / remove rendered artifacts keyed
  `(gameID, turn, playerID, format)`; placement interacts with §13.1 (its own store, or alongside the
  snapshots). 🟡
- **JSON results schema** (§12.6) — a **versioned** projection of `domain.PlayerReport`, so format changes do
  not silently break scripted consumers; the deferred SQLite export must not warp it. 🟡

The largest **mechanics** still open — the **intra-turn tick model** (day-by-day vs. event-merge, §11.3) and
the **Close-phase money/upkeep slot** (§6.1/§6.2/§11.6) — are tracked in their owning sections, **not** this
infra table; they gate the §6.1/§6.2/§8.2/§6.6 interleavings but are not technology choices.

### 13.8 Already resolved out of the table (for the record) ✅

- **Randomness source** — closed. Stochastic draws go through the **`RNG` port** (`app/ports.go`) realized by
  the **`internal/infra/prng`** PCG adapter, mirroring `Clock`; RNG state round-trips with the snapshot via
  `GameStateStore` (§2.9/§11.7/§11.9). Removed from the open table and promoted to AGENTS.md's **Ports** list.
  This section newly adds **§13.3 (order file format)** and **§13.6 (concurrency)** to the resolved set.

> **AGENTS.md sync pending.** §13 records the verdicts; the table itself still shows every row as open. To
> finish the loop, AGENTS.md wants: an ADR note marking **Order file format** decided (custom DSL, §10.1) and
> **Concurrency** confirmed (serial-per-game, §11), three **new rows** for the surfaced decisions (§13.7:
> map-artifact format / report-store / JSON schema), and the **PDF-library** and **Mail-transport** rows
> annotated with the constraints §13.2/§13.4 add. Until then this section is the single place that reconciles
> the design against the table; promotion of §1–§12 to the `reference/model/` and how-to docs can proceed in
> parallel, since none of these open infra choices changes a domain rule.
