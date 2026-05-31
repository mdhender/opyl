---
title: Tutorials
weight: 1
prev: /
next: /tutorials/first-turn
---

Tutorials are **learning-oriented**. Each one walks a newcomer through a
self-contained experience that ends in something visibly working. They are
not how-tos: they teach by guided practice, not by addressing a real-world
goal.

## What belongs here

- "Your first turn" — fixture game, fixture orders, render to stdout
- "Add a player to a running game" — guided practice with state files
- "Generate your first PDF report" — fixed inputs, predictable output

## What does **not** belong here

- "How do I switch from text to PDF rendering?" — that is a [how-to](../how-to)
- Lists of CLI flags — that is [reference](../reference)
- Discussion of why turns are resolved serially — that is [explanation](../explanation)

## Writing rules for this section

- Use "we will…" language; the reader is a learner, not an operator.
- Every command must work exactly as written; no "you might need to…".
- Show the goal upfront and the visible result early.
- Minimise explanation. Link to it; do not inline it.
- If a tutorial ever requires the learner to make a choice, fix the choice
  for them and move the choice to a how-to.

{{< cards >}}
  {{< card link="first-turn" title="Your first turn" subtitle="Resolve a fixture game from orders to rendered report." >}}
{{< /cards >}}
