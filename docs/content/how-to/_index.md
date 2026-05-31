---
title: How-to guides
weight: 2
prev: /tutorials
next: /reference
---

How-to guides are **task-oriented**. Each one assumes a competent operator
or developer who already knows what they want to accomplish and just needs
the steps.

## What belongs here

- "Switch the PDF renderer from `gofpdf` to `typst`"
- "Add a new flat-file order format"
- "Run the pipeline under cron"
- "Recover from a failed dispatch step"
- "Implement a new `app` port and wire its adapter"

## What does **not** belong here

- "Your first turn" — that is a [tutorial](../tutorials)
- "Available CLI flags for `opyl process`" — that is [reference](../reference)
- "Why opyl uses two ports for render and dispatch" — that is [explanation](../explanation)

## Writing rules for this section

- The title states exactly the goal: "How to <do the thing>".
- Assume the reader already knows the surrounding context.
- Handle real-world edge cases ("if SMTP authentication fails…").
- Omit anything the reader does not need to complete the task.
- Do not teach concepts; link to explanation if background is needed.

{{< cards >}}
  {{< card link="add-order-format" title="Add a new order file format" subtitle="Implement a new OrderSource adapter under internal/infra/orderfile/." >}}
{{< /cards >}}
