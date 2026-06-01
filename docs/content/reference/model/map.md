---
title: Map
weight: 1
prev: /reference/model
---

The world is a square grid of **provinces**. Each province has a terrain type, a
civilization level, the **region** it belongs to, a set of outgoing routes, and
zero or more inner sub-locations.

## The map artifact

- The map is a **fixed, authored artifact**, loaded before turn resolution and
  treated as immutable input to it. It is not generated at run time.
- Loading the artifact is an infra concern; the on-disk format is not specified
  in this reference.
- Map **dimensions are a property of the artifact**, not fixed by the rules.
- **Regions** are named collections of provinces, authored with the map: every
  province is assigned to exactly one region. Region membership is loaded as
  immutable map data; how regions function politically (garrison binding,
  king-hood) belongs to the rules, not this reference.
- The artifact also carries **seed data** loaded as immutable input: the
  locations of the **initial settlements** where new factions begin, and the
  list of **cities that are safe havens**. Their mechanics belong to the rules,
  not this reference.

## Province identity and coordinates

- A province's identity is a **one-based `(row, col)` integer pair**. `(1, 1)` is
  the top-left (northwest) corner; rows increase **south**, columns increase
  **east**, out to `(n, n)`.
- The bracketed **display code** (e.g. `a1`, `ae48`) is presentation only. No
  rule depends on it; the province *is* its `(row, col)` pair.
- A display code is the **row letters followed by the column number**,
  compressed: no fixed width, no leading zeros.
  - The **row** is a bijective base-22 numeral over the alphabet
    `a b c d e f g h k m n p q r s t u v w x y z` (the Latin alphabet without
    `i`, `j`, `l`, `o`). `a` = 1 … `z` = 22, `aa` = 23, `ab` = 24, …. There is no
    zero digit and no upper bound.
  - The **column** is a plain decimal ordinal.
  - `(1, 1)` renders as `a1`.
- Sub-location codes are arbitrary and carry **no** coordinate meaning.

For the encode/decode procedure and the reasoning behind the scheme, see
[Map coordinate compression]({{< relref "/explanation/map-coordinates" >}}).

## Terrain

Every province has exactly one of six terrain types:

`plains` · `forest` · `swamp` · `mountain` · `desert` · `ocean`

Terrain affects movement cost (below). Other terrain effects (resource yield,
defense, sighting) are not specified in this reference.

## Routes and movement

- Provinces are adjacent in the **four orthogonal directions** (north, east,
  south, west). Diagonal travel costs two moves. Map edges are impassable.
- Each route carries a **nominal cost in whole days**, authored per route.
- Land travel auto-selects the **fastest available mode**: horseback when the
  whole party can be mounted, otherwise on foot; rough terrain may negate the
  horse benefit. Ocean travel requires a ship.
- Actual cost may differ from the nominal cost. The variance is a
  **deterministic function of the per-turn seed** recorded for the turn. The
  variance model (per-mode, per-terrain, and weather modifiers) is not specified
  in this reference.
- Ocean↔mountain routes are **impassable**: a ship cannot dock against a
  mountain province. A ship docks into an adjoining land province in 1 day.
- `FLY` exists in the order set; its movement rules are not specified in this
  reference.

## Inner locations and visibility

- A province may contain **sub-locations** (city, inn, port, island, …), entered
  from the surrounding province. `MOVE IN` enters the first listed sub-location
  when unambiguous.
- Entering a sub-location costs the route's listed days, or 0 when no time is
  listed.
- Occupants of a sub-location receive a report for the **immediately surrounding
  location** only. Outsiders cannot see into a sub-location without entering it.
- Co-located characters interact without travel.
- A **port city** gates ocean access: the surrounding province cannot reach the
  ocean directly; ships sail into and out of the port city itself.

## Holes and hidden routes

- A province may have **no route** in a given direction — a hole, whether
  permanently impassable or merely undiscovered.
- A **hidden route** is discoverable via `EXPLORE`. Once found, it is usable by
  the **whole owning faction** and by no other faction, even one that knows the
  destination code.
- Stack-mates crossing a hidden route learn it; prisoners do not.

## Civilization level

Every province has an integer civilization level ≥ 0. A level of `0` is
wilderness. There is no upper cap.

Each turn:

```
civ(p) = max( buildings(p), floor( maxNeighborCiv / 2 ) )
```

- **`buildings(p)`** sums the contribution of each feature present, counting only
  the **first of each type**; fractional remainders are dropped **after**
  summing.

  | Feature    | Contribution                  |
  | ---------- | ----------------------------- |
  | Safe Haven | 2                             |
  | Castle     | 1.5 + improvement level / 4   |
  | City       | 1                             |
  | Tower      | 1                             |
  | Temple     | 1                             |
  | Inn        | 1                             |
  | Mine       | 1                             |

- **`maxNeighborCiv`** is the maximum civ level among the four orthogonal
  neighbors, **read from the previous turn's values**. Off-map and hole
  neighbors count as `0`. The computation is a single pass with no fixpoint, so
  civilization spreads at most one hop per turn.
- At turn zero, civ comes from the authored map; absent an authored value, the
  first computation uses `buildings(p)` only.

## See also

- [Geography & Movement]({{< relref "/rules/geography" >}}) — the player-facing
  rulebook these facts are distilled from.
- [Map coordinate compression]({{< relref "/explanation/map-coordinates" >}}) —
  why identity and display code are kept separate.
