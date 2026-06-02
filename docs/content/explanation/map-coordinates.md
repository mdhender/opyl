---
title: Map coordinates
weight: 4
prev: /explanation/use-cases
---

A province has two things that look like names. One is its **identity** — the
axial `(q, r)` pair the engine stores and reasons about. The other is its
**code** — the `[q,r]` a player reads in a report, like `[8,-5]`. This page
explains why opyl uses hexagonal provinces with axial coordinates, why the
origin sits at the centre of the world, and how the two layers relate.

## Why hexagons

opyl's ancestor laid provinces out on a square grid: four neighbours, and a
diagonal step that had to be faked with two moves. Squares force an awkward
choice — either diagonals are forbidden (movement feels blocky) or they are
allowed but cost differently from orthogonal steps (distance gets fiddly).

Hexagons remove the dilemma. A flat-topped hex has **six neighbours, each across
one shared edge, each the same kind of step**. There is no diagonal to special
-case: North, Northeast, Southeast, South, Southwest, and Northwest are all
first-class directions. Movement cost still varies — terrain and authored route
days differ — but no direction is structurally cheaper or more expensive than
another. (There is no due east or west: those point at a vertex, not an edge.)

## Axial `(q, r)`, flat-top

We use the [Red Blob Games](https://www.redblobgames.com/grids/hexagons/) **axial
flat-top** convention. A province is a pair of signed integers `(q, r)`; a third
cube axis `s = −q − r` is implied and never stored. The six neighbours are
reached by adding a fixed vector — North is `(0, −1)`, Northeast `(+1, −1)`, and
so on around the hex. Neighbour, distance, and line-of-travel queries are all
plain integer arithmetic on the pair, exactly as the square scheme's were.

The direction vectors are **screen-true**: North is straight up on the map the
game master draws, so "north of" in a report matches "above" on the GM's map.

## A central origin, not a corner

The square scheme numbered from a corner — `(1, 1)` at the top-left, counting
only positive numbers, with a fixed `(n, n)` far corner. That bakes the map's
size into its coordinates: grow the world and everything renumbers.

opyl puts the **origin `(0, 0)` near the centre** of the authored world,
wherever the GM chooses. Coordinates run **negative and positive in every
direction**, so the world can be extended north, south, or outward in any
direction *without touching a single existing coordinate*. A province minted on
a new northern frontier just has a more-negative `r`; nothing already on the map
moves. There is no off-map edge baked into the numbering — where the authored
world ends, a province simply has no route that way (a hole).

## Two layers, thinner than before

The engine still never depends on the *spelling* of a code: identity is the
`(q, r)` pair, and printing it is a rendering concern. But unlike the old
bijective base-22 scheme — where `(27, 48)` rendered as the opaque `[ae48]` —
the printed form now **is** the pair, comma-separated in brackets: `(8, -5)`
shows as `[8,-5]`. The two layers haven't merged, but the display layer has
almost nothing left to do, which is the point: there is no encoding for a player
to learn, mis-key, or mis-read aloud.

## Display grammar: emit strict, accept lenient

Codes are typed into order files by hand and pasted out of emailed reports, so
the grammar is forgiving on the way in and rigid on the way out.

- **Emit strict.** The engine always prints the canonical form: square brackets,
  a comma, no spaces, no leading `+`, and `-0` normalised to `0`. `[8,-5]`,
  `[0,0]`, `[-3,12]`.
- **Accept lenient.** The order parser tolerates interior whitespace — `[ 8 , -5 ]`
  reads the same as `[8,-5]` — and normalises whatever it accepts back to the
  canonical form before anything else sees it. Order files are untrusted input;
  anything that is not a well-formed `[q,r]` is rejected, never guessed at.

Players reference a province in an order with the same bracketed code they see in
the report's route list, so the usual workflow is copy-and-paste.

## See also

- [Geography & Movement]({{< relref "/rules/geography" >}}) — the player-facing
  rules where these codes appear.
- [Map]({{< relref "/reference/model/map" >}}) — the engine reference, including
  the full direction-vector table.
- `GAME-DESIGN.md` §2.2 — the design decision this page expands on.
