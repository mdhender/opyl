---
title: Map artifact format
weight: 2
prev: /reference/model/map
---

The authored map is a single JSON document — the input the `MapSource` port
loads and treats as immutable. It is produced by `cmd/woly` from a Worldographer
(`.wxx`) source and is never hand-edited as the source of truth.

The document is **deterministic**: every collection is a sorted array (no JSON
objects-as-maps), every numeric field is an integer, and the same source
produces byte-identical output. It is therefore git-diffable.

## Top-level document

```json
{
  "schemaVersion": 1,
  "origin": { "wxy": {"x": 5, "y": 8}, "qr": {"q": 0, "r": 0} },
  "nextEntityId": 6001,
  "regions":   [ /* … */ ],
  "provinces": [ /* … */ ]
}
```

| Field           | Type    | Notes                                                                       |
| --------------- | ------- | --------------------------------------------------------------------------- |
| `schemaVersion` | int     | Document shape version.                                                     |
| `origin`        | object  | Provenance: the Worldographer hex `wxy` that `woly` pinned to axial `qr`. The engine does not need it. |
| `nextEntityId`  | int     | First entity number the runtime minter may use — one past the last id `woly` minted. |
| `regions`       | array   | Sorted by `id` ascending.                                                   |
| `provinces`     | array   | Sorted by `q` ascending, then `r` ascending.                                |

## Region

```json
{ "id": "tollus", "name": "Tollus", "kind": "normal" }
```

| Field  | Type   | Notes                                                       |
| ------ | ------ | ----------------------------------------------------------- |
| `id`   | string | Stable slug; referenced by `province.region`.               |
| `name` | string | Display name.                                               |
| `kind` | string | `normal` · `hades` · `faery` · `cloudlands`.                |

## Province

```json
{
  "q": 8, "r": -5,
  "terrain": "plains",
  "region": "tollus",
  "civSeed": null,
  "routes": [ /* … */ ],
  "sublocations": [ /* … */ ]
}
```

| Field          | Type        | Notes                                                                          |
| -------------- | ----------- | ------------------------------------------------------------------------------ |
| `q`, `r`       | int         | Axial coordinate; the province's identity.                                     |
| `terrain`      | string      | `plains` · `forest` · `swamp` · `mountain` · `desert` · `ocean`.               |
| `region`       | string      | Region `id`; exactly one.                                                      |
| `civSeed`      | int or null | Authored turn-zero civ level. `null` means derive from buildings.              |
| `routes`       | array       | Outgoing edges, sorted in direction order N, NE, SE, S, SW, NW.                |
| `sublocations` | array       | Authored static sub-locations, sorted by `id`.                                 |

A direction with **no `routes` entry is a hole** — there is no route that way.

## Route

```json
{ "dir": "SE", "to": {"q": 9, "r": -5}, "days": 1, "impassable": false, "hidden": false, "waterName": "Tymaerian Sea" }
```

| Field        | Type   | Notes                                                                           |
| ------------ | ------ | ------------------------------------------------------------------------------- |
| `dir`        | string | `N` · `NE` · `SE` · `S` · `SW` · `NW`.                                           |
| `to`         | object | Destination province `{q, r}`.                                                  |
| `days`       | int    | Nominal traversal cost in days.                                                 |
| `impassable` | bool   | The route exists (and is shown) but cannot be traversed, e.g. ocean↔mountain.   |
| `hidden`     | bool   | Discoverable via `EXPLORE`; usable only by the owning faction once found.       |
| `waterName`  | string | Sea name shown for ocean routes; `""` when none.                                |

Routes are stored per origin province and need not be symmetric. `woly` emits
the matching reverse edge on the neighbouring province.

## Sub-location

Sub-locations nest recursively: an inn inside a city inside a province.

```json
{
  "id": 2845, "srcUuid": "f1c2…", "type": "city", "name": "Carim",
  "entryDays": 1, "safeHaven": false, "initialSettlement": false,
  "routes": [],
  "sublocations": [
    { "id": 3102, "srcUuid": "9ab0…", "type": "inn", "name": "Hooting Owl Inn",
      "entryDays": 0, "safeHaven": false, "initialSettlement": false,
      "routes": [], "sublocations": [] }
  ]
}
```

| Field               | Type   | Notes                                                                          |
| ------------------- | ------ | ------------------------------------------------------------------------------ |
| `id`                | int    | Entity number minted by `woly`.                                                |
| `srcUuid`           | string | Worldographer source UUID, kept as provenance; the engine ignores it.          |
| `type`              | string | `city` · `town` · `inn` · `port-city` · `island` · `tower` · `temple` · `mine` · `castle`. |
| `name`              | string | Display name.                                                                  |
| `entryDays`         | int    | Days to enter from the surrounding location; `0` when entry is free.           |
| `safeHaven`         | bool   | Safe-haven designation (feeds the civ contribution and havens rules).          |
| `initialSettlement` | bool   | A location where a new faction may begin.                                      |
| `routes`            | array  | Extra routes such as a port-city's sea access. The `OUT` edge to the parent is implicit and not stored. |
| `sublocations`      | array  | Nested sub-locations, sorted by `id`.                                          |

The engine flattens this static tree on load into the same containment relation
it uses for runtime positions: every locatable entity is `in` either a province
`(q, r)` or a container entity `id`. Mobile entities (ships, nobles) are **not**
in the artifact — their location is per-turn snapshot state.

## Coordinate conversion (Worldographer → axial)

The source is a flat-top, **odd-q vertical** layout (odd columns shifted down).
Offset `(col, row)` converts to axial `(q, r)`:

```
q = col
r = row − (col − (col & 1)) / 2
```

The inverse is `col = q`, `row = r + (q − (q & 1)) / 2`.

`cmd/woly --x-y X,Y --q-r Q,R` pins Worldographer hex `(X, Y)` to axial
`(Q, R)`. Re-centering is computed in axial space: convert the pinned hex and
each hex to axial, then `final = hexAxial − pinAxial + (Q, R)`. Raw offset
subtraction is incorrect because `r` depends nonlinearly on `col`.

## Production

`cmd/woly` always emits a **complete** artifact from the source — it does not
read, update, or extend a prior artifact. Entity numbers are stable within one
import; editing the source and re-importing may renumber them.

## See also

- [Map]({{< relref "/reference/model/map" >}}) — the map model these facts
  serialize.
- [Map coordinates]({{< relref "/explanation/map-coordinates" >}}) — why axial
  hex coordinates and a central origin.
- `docs/adr/` ADR 0004 — the binding decision and its rationale.
