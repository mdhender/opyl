---
title: Explanation
weight: 4
prev: /reference
---

Explanation is **understanding-oriented**. It makes connections, gives
context, and explains why opyl is the way it is. It is the appropriate
place for opinion, history, and discussion of alternatives.

## What belongs here

- "Why opyl uses SOUSA layering"
- "Why rendering and dispatch are two ports, not one"
- "Why turns are processed serially per game"
- "How storage choice affects opyl's audit trail"

## What does **not** belong here

- Step-by-step procedures — those are [how-tos](../how-to)
- API signatures — those are [reference](../reference)
- Guided practice — that is [tutorials](../tutorials)

## Writing rules for this section

- Make connections; do not list facts.
- Admit opinion and perspective; acknowledge alternatives.
- Stay bounded to the topic; do not drift into how-to or reference content.
- It is fine to discuss history and design trade-offs.

{{< cards >}}
  {{< card link="sousa-in-opyl" title="SOUSA in opyl" subtitle="Why opyl uses strict Onion/Clean layering, and how it adapts SOUSA to a batch CLI shape." >}}
  {{< card link="idempotency" title="Idempotency by design" subtitle="Why ProcessTurn must be safe to rerun and where that responsibility lives." >}}
  {{< card link="use-cases" title="How use cases work" subtitle="Why each pipeline stage is one method on app.Services, and what that shape buys us." >}}
  {{< card link="map-coordinates" title="Map coordinates" subtitle="Why provinces are flat-top hexes with axial (q,r) coordinates, why the origin sits at the centre of the world, and how the [q,r] code works." >}}
{{< /cards >}}
