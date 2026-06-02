---
title: Map
weight: 1
prev: /reference/model
---

The world is a grid of flat-topped **hexagonal provinces** — each adjacent to up
to six others, one across each edge. Each province has a terrain type, a
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

- A province's identity is its **axial coordinate `(q, r)`** — a pair of signed
  integers (Red Blob Games flat-top axial convention). The implied third cube
  axis is `s = −q − r`.
- The **origin `(0, 0)`** is a hex near the centre of the authored world (the
  GM's choice). Coordinates run negative and positive in every direction; the
  world can extend outward without renumbering.
- The bracketed **display code** *is* the coordinate, printed as `[q,r]` (e.g.
  `[8,-5]`, `[0,0]`). No rule depends on a separate encoding; the province *is*
  its `(q, r)` pair.
  - **Emit strict:** the canonical form has no spaces, no leading `+`, and `-0`
    normalised to `0`.
  - **Accept lenient:** order-file parsing tolerates interior whitespace
    (`[ 8 , -5 ]`) and normalises to canonical.
- Sub-location codes are **entity numbers** (see the entity model), not
  coordinates, and carry **no** spatial meaning.

For the reasoning behind axial hex coordinates and the central origin, see
[Map coordinates]({{< relref "/explanation/map-coordinates" >}}).

## Terrain

Every province has exactly one of six terrain types:

`plains` · `forest` · `swamp` · `mountain` · `desert` · `ocean`

Terrain affects movement cost (below). Other terrain effects (resource yield,
defense, sighting) are not specified in this reference.

## Routes and movement

- Each province has **up to six edges**, in the directions **North, Northeast,
  Southeast, South, Southwest, Northwest**. There is no due east or west and no
  diagonal travel: a step crosses exactly one edge, and the six edges are
  structurally equal — none costs extra for its direction. The axial direction
  vectors are:

  | Direction | Abbr | Δ(q, r) |
  | --------- | ---- | ------- |
  | North     | N    | (0, −1) |
  | Northeast | NE   | (+1, −1) |
  | Southeast | SE   | (+1, 0)  |
  | South     | S    | (0, +1)  |
  | Southwest | SW   | (−1, +1) |
  | Northwest | NW   | (−1, 0)  |

  Where the authored world ends, a province simply has no route in that
  direction (a hole).
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

- **`maxNeighborCiv`** is the maximum civ level among the **six hex-adjacent
  neighbours**, **read from the previous turn's values**. Missing and hole
  neighbours count as `0`. The computation is a single pass with no fixpoint, so
  civilization spreads at most one hop per turn.
- At turn zero, civ comes from the authored map; absent an authored value, the
  first computation uses `buildings(p)` only.

## See also

- [Geography & Movement]({{< relref "/rules/geography" >}}) — the player-facing
  rulebook these facts are distilled from.
- [Map artifact format]({{< relref "/reference/model/map-artifact" >}}) — the
  JSON the `MapSource` port loads these facts from.
- [Map coordinates]({{< relref "/explanation/map-coordinates" >}}) — why axial
  hex coordinates and a central origin.
