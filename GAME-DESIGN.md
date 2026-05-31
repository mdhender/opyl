# GAME-DESIGN.md — opyl

The living source of truth for opyl's **game rules** while they are being designed.

This file is design, not published documentation. As sections stabilize they feed the
Diataxis docs in `docs/content/`:

- Mechanical facts (order syntax, phase order, entity attributes) → **reference** pages
- Rationale and trade-offs ("why X resolves before Y") → **explanation** pages
- It also resolves the open decisions tracked in [AGENTS.md](AGENTS.md).

Status legend: ✅ decided · 🟡 partially specified · ❓ open question, undecided.

---

## 1. Concept & victory

opyl is an **open-ended fantasy game**. ✅

- Each player controls a single **nation**. ✅
- Players are **free to pursue their own goals** — there is no single imposed win
  condition. ✅
- A nation is managed through its **characters**: entities with well-defined attributes
  that **accept orders** (a nation starts with one and grows — see §3). ✅
- Characters command **minions**, who carry out the orders. ✅
- **Scoring** exists but is **deferred** — to be designed later. 🟡
- **A game ends** when no players are submitting turns, **or** when the GM retires. ✅

## 2. The map

### Grid & coordinates

- **Hex grid**, **axial coordinates** `(q, r)`, **flat-top** hexes
  ([Red Blob Games](https://www.redblobgames.com/grids/hexagons/) terminology). ✅
- Map **center is `(0, 0)`**. ✅
- Extends an **unspecified number of hexes** in all directions. ✅
- A nation is **assigned a starting hex** at creation, **usually a major city**. ✅

### The six directions

Flat-top hexes have six neighbours, in opposite pairs (N↔S, NE↔SW, SE↔NW). ✅

| Direction | `(Δq, Δr)` |
| --------- | ---------- |
| North (N)      | `( 0, -1)` |
| Northeast (NE) | `(+1, -1)` |
| Southeast (SE) | `(+1,  0)` |
| South (S)      | `( 0, +1)` |
| Southwest (SW) | `(-1, +1)` |
| Northwest (NW) | `(-1,  0)` |

An edge in direction `d` from hex `(q, r)` leads to `(q + Δq_d, r + Δr_d)`.

Distance between two hexes is the orientation-independent axial/cube distance
`(|Δq| + |Δq + Δr| + |Δr|) / 2`.

### Terrain

Every hex has exactly one **terrain** type. ✅

| Land | Water |
| ---- | ----- |
| Plains | Lake |
| Forest | Sea |
| Mountains | Ocean |
| Hills | |
| Forested Hills | |
| Wasteland | |

> ❓ **Open — terrain effects.** Beyond movement cost, what does terrain drive — resource
> yield, defence, visibility/sighting range? Which water types are passable by which
> travel modes?

### Hex contents

Besides terrain, a hex may hold: ✅

- **Resources** — e.g. minerals, grain, lumber.
- At most **one settlement** — a single entity with a **size** that grows
  **village → town → city**. ✅
- **Plus** at most **one** of: a **castle** *or* a **wizard's tower**. ✅
- **Plus** optionally **one ruins**, which are **usually hidden** from players. ✅
- **Occupants**: player characters/minions, **NPCs**, and **monsters**.

So the structure slots are: `[1 settlement?] + [1 castle | 1 tower]? + [1 ruins?]`.

> ❓ **Open — resource model.** Are resources quantities held by the hex? Do they
> regenerate? How are they harvested into a nation's economy?
> ❓ **Open — structure attributes.** What attributes do settlement / castle / tower /
> ruins carry (owner, garrison, size thresholds for growth, what hides ruins and how
> they're discovered)? These fill `domain.GameState` and become *Game entities* reference.

### Movement

- A character travels between hexes by **walking, riding, or flying** (its minions move
  with it as inventory). ✅
- **Movement points (MPs)** are spent to move between hexes. ✅
- Movement cost is **directional (asymmetric)**: the cost to move `(0,0) → (0,-1)` may
  differ from `(0,-1) → (0,0)`. Cost is therefore a property of the **directed edge**,
  not the hex pair. ✅
- Asymmetry arises from things like **downhill vs uphill**, a **one-way ford** on a
  river, or **magical** effects. ✅

**Provisional MP budgets** (🟡 strawman — fill the section and tune later, not settled):

| Mode | MP / turn | Notes |
| ---- | --------: | ----- |
| Walking | 12 | baseline foot movement |
| Riding  | 24 | mounted; may be barred by some terrain |
| Flying  | 36 | ignores terrain entry cost (pays a flat 1/hex) |

**Cost model — two layers.** ✅

1. **Default layer (dense, terrain-driven).** A lookup keyed by the **terrain transition**
   (`from-terrain → to-terrain`) yields the base `(walk, ride, fly)` cost — e.g. the cost
   of "Plains → Forest". This is the cost almost every edge uses.
2. **Override layer (sparse, edge-driven).** A directed edge is identified by the tuple
   **`(q, r, direction)`** and stores a `(Δwalk, Δride, Δfly)` **delta** applied on top of
   the default. This is where slope, one-way fords, and magic live. **Only edges whose
   delta ≠ `(0,0,0)` are stored** — the vast majority fall through to the default.

So `cost(edge, mode) = default[fromTerrain → toTerrain][mode] + override[(q,r,dir)][mode]`,
with the override term zero when absent.

Provisional default `(walk, ride)` entry cost by destination terrain (🟡 strawman):
Plains 1, Hills 2, Forest 2, Forested Hills 3, Wasteland 2, Mountains 4; water impassable
on foot/horse. Flying pays a flat 1/hex.

> Note: the sparse override table is graph-shaped, queryable, mostly-empty state — a data
> point for the AGENTS.md **state storage** decision (favours a keyed table / sparse map).

## 3. Nation, characters, minions

| Concept | Specified | Open |
| --- | --- | --- |
| **Nation** | A **collection of player-controlled characters** — no separate nation object beyond its characters. Starting hex (usually a city); starts with **1 character**; earns **growth points** computed via the GURPS **[City Stats](https://www.sjgames.com/gurps/books/citystats/)** supplement, from **territory (hex count) and resources controlled**, spent to recruit more characters. **No separate treasury**: a nation's money lives in **one character's inventory** — a deliberate single-point-of-failure risk wise players will mitigate. ✅ | ❓ How is "territory controlled" determined (occupied/adjacent hexes)? |
| **Character** | The core entity and the unit that accepts orders. Has **attributes** (ST/DX/IQ/HT), **skills**, and an **inventory**; is **always in a location** (a hex). Player characters are created by spending nation growth points. **On death**, her inventory is **distributed to the characters and NPCs at her location**; if none are present, it is **lost** (this is what gives the single-inventory treasury its bite). ✅ | ❓ Cap on character count? |
| **NPC** | A **game-controlled character** — *the same entity model as a player character*, run by the engine with **its own goals**. NPCs **command monsters and minions** and can be **hired** (mercenaries, teachers). ✅ | ❓ Hiring mechanics (cost, duration, loyalty)? ❓ Does an NPC belong to a nation or stand alone? |
| **Minion** | A **typed, counted inventory item**, **transferable between characters** like any inventory. Types: **warrior, priest, scout, laborer** — each is **like a GURPS template**, carrying a set of **skills and combat values**. **Recruited with money or the Charisma skill.** **No health** — alive or dropped from inventory. ✅ | — |
| **Monster** | First-class entity tracking **number, health, combat abilities**; typically **commanded by NPCs**. **Attracted to wilderness, ruins, and food sources**; **avoids civilisation** (an **aggression knob** tunes how strongly) but will **raid for food and supplies**. Engine grows numbers in wild areas, shrinks them as civilisation encroaches. ✅ | 🟡 **Monster agents (order-generation logic) — TBD.** ❓ Exact attraction/aggression model and raid triggers? |

### Character model

**One entity, two controllers.** Player characters and NPCs share this *same* model; they
differ only in who issues their orders (a player vs. the engine). Build it once. ✅

A character carries: ✅

- **Attributes** — borrowed from GURPS:
  - **ST** (Strength), **DX** (Dexterity), **IQ** (Intelligence), **HT** (Health).
- **Skills** — e.g. **Combat, Stealth, Diplomacy, Magic**, and **Charisma** (used to
  recruit minions). The system **leans heavily on GURPS** for both attributes and skills.
  ✅ 🟡 *(full skill list, levels, and which orders each gates/modifies still TBD.)*
- **Location** — the hex the character occupies; always present. ✅
- **Inventory** — everything the character holds:
  - **Minions** as typed, counted lines (e.g. `ogre × 3`), transferable.
  - **Money** — a nation's treasury is simply the money in some character's inventory.
  - 🟡 *(other item kinds — equipment, carried resources — TBD.)*

Orders attach to a **character**; the character's minions and inventory are the means of
carrying them out.

> 🟡 **GURPS dependency.** Attributes (ST/DX/IQ/HT) and skills are GURPS-derived. Worth an
> *explanation* page later on how much GURPS we adopt vs. simplify, and a licensing note.

## 4. Orders 🟡

Orders are issued **to characters**. Each verb is documented with a fixed template:
**Syntax · Parameters · Effect · Errors** — the same shape the *Order catalog* reference
page will mirror.

> ❓ **Open — order file format.** Still undecided in AGENTS.md (custom DSL vs YAML vs
> structured email). The character-issues-orders model leans toward a readable
> line-oriented form like `<character>: <verb> <args>`, used below provisionally.
> ❓ **Open — full catalog.** Beyond `move`: build, recruit, scout, attack, trade, train,
> hire (NPC)… to be specced one at a time.

### `move` — relocate a character across the map

Spec **✅ accepted.** Built from already-decided rules (§2 directions, modes, MP budgets,
two-layer cost; §3 location). Only the strawman MP/cost *numbers* remain 🟡 pending tuning.

**Syntax**

```
<character>: move <mode> <direction> [<direction> ...]
```

Example: `Roderik: move ride N N NE`

**Parameters**

| Parameter | Values | Meaning |
| --------- | ------ | ------- |
| `<character>` | a character name | The acting character; must belong to the issuing nation. |
| `<mode>` | `walk` \| `ride` \| `fly` | Sets the MP budget and which default-cost column applies. |
| `<direction>…` | one or more of `N NE SE S SW NW` | An **ordered path of steps** from the character's current hex. |

The path is expressed as **directions, not a destination** (🟡 design choice) — this fits
the directed-edge cost model and keeps resolution deterministic without the engine
choosing routes.

**Effect**

1. The character **and its entire inventory** (minions, money, items) move together,
   one step at a time along the path. (To leave minions behind, use a transfer/drop
   order — there is no implicit split.) ✅
2. Each step from hex `H` in direction `d` (to hex `H'`) costs
   `default[terrain(H) → terrain(H')][mode] + override[(H.q, H.r, d)][mode]`.
3. The MP budget for the turn is the mode budget (🟡 strawman: walk 12 / ride 24 / fly 36).
   Steps are taken until the next step's cost exceeds remaining MP, then the character
   **stops at the last reached hex** (partial move).
4. Entering a hex reveals it and its surroundings per fog-of-war sighting (§6).

> ✅ **Resolved by §5 phases.** A character gets one `move` per turn, governed by its MP
> budget (phase 2a). Combat is a separate phase (2b), *not* paid from MP — so a character
> can move and then fight in the same turn.

**Errors / outcomes**

| Condition | Result |
| --------- | ------ |
| Unknown character, or not owned by the issuing nation | rejected at validation (`cerr.ErrInvalidOrders`) |
| `<mode>` not `walk`/`ride`/`fly`, or bad direction token | rejected at validation (`cerr.ErrInvalidOrders`) |
| Mode cannot traverse the next step (e.g. `walk`/`ride` into water) | character **stops before** that step; remainder of path ignored; reported in the turn report |
| Insufficient MP for even the first step | character does not move; reported |

> ✅ **Accepted:** (a) directions-not-destination, (b) whole inventory always moves with
> the character, (c) impassable/over-budget steps are a **soft stop reported in the
> report**, not a hard order error.

## 5. Turn resolution 🟡

**Simultaneous resolution.** ✅ All nations' orders for a turn are resolved together —
classic PBEM. No nation moves "before" another; order of submission does not matter.

A turn passes through these phases, in order: ✅

| # | Phase | Sub-phase | Does |
| - | ----- | --------- | ---- |
| 1 | **Initialize** | 1a. **Heal** | Recover health (characters, monsters). |
|   |                | 1b. **Monster Growth** | Engine grows monster numbers in wild areas (§3). |
| 2 | **Orders** | 2a. **Move** | Resolve all `move` orders (§4). |
|   |            | 2b. **Combat** | Resolve combat between forces now sharing a hex. |
| 3 | **Economy** | | Growth points (City Stats), production, recruiting costs, upkeep. |
| 4 | **Reporting** | | Produce each player's report (§6). |

Because movement (2a) completes before combat (2b), simultaneous moves settle first and
combat then resolves wherever opposing forces ended up co-located — the standard way
simultaneous-movement PBEM avoids "who moved first" disputes. ✅

> ❓ **Open — intra-phase determinism.** Within a simultaneous phase, conflicts still need
> a deterministic tie-break (two forces want the same hex; multi-party combat order). What
> canonical ordering applies — by character id, by a seeded RNG recorded in the turn
> ledger? This is the crux of "resolve deterministically."
> ❓ **Open — Heal & Combat rules.** What heals and by how much (GURPS HT-based?); what
> *triggers* combat (co-location + hostility?) and how it is resolved.
> ❓ **Open — Economy.** Depends on the still-undefined resource/economy model (§2).

## 6. The player report 🟡

There **is a fog of war**: players see only what they can observe. ✅
A turn report includes **what the player's characters see** and **what their minions
report in the current turn**. ✅

> ❓ **Open — report detail.** Beyond character sightings and minion reports, what else
> is included — character/minion status, order outcomes (success/failure), economy,
> events/messages? What exactly does "see" cover (current hex + adjacent? sighting range
> by terrain)? How is stale knowledge of previously-seen hexes presented? This becomes
> the *Report contents* reference and fills `domain.PlayerReport`.

## 7. Open decisions carried from AGENTS.md

These remain unresolved and intersect the game design:

| Decision | Note |
| --- | --- |
| Order file format | See §4. The character/order model constrains it. |
| State storage | Hex map + nations + characters is graph-ish, queryable state — informs SQLite vs per-turn JSON. |
| Concurrency | Turns resolve serially per game (§5 simultaneity). |

---

## Next design steps (in order)

Done so far: §1 concept; §2 map (grid, directions, terrain set, hex contents, two-layer
cost); §3 entity roster + character model; §4 `move` (accepted); §5 phase sequence +
simultaneity.

1. **§5 detail** — intra-phase determinism (tie-break / seeded RNG in the turn ledger),
   Heal & Combat rules. The determinism rule blocks any combat-bearing verb.
2. **§4 next verbs** — recruit, build, attack, hire… one at a time.
3. **Economy model** (§2/§3) — resources/production that Economy (phase 3) and growth need.
4. **§6 report contents** — what the turn report shows.
5. **§2/§3 leftovers** — terrain effects, structure attributes, sighting rules.

Only after a slice is settled (and ideally reflected in `internal/domain`) do we write the
corresponding **reference** page that *describes* it.
