---
title: Reference
weight: 3
prev: /how-to
next: /explanation
---

Reference is **information-oriented**. It describes opyl's machinery
neutrally and completely. It does not teach, persuade, or guide. The
reader is at work and needs an accurate fact.

## What belongs here

- Every CLI subcommand: flags, arguments, exit codes
- Every application port: signature, contract, returned errors
- File formats: order files, state snapshots
- Sentinel errors and their meanings
- Configuration: flags, env vars, defaults

## What does **not** belong here

- Step-by-step procedures — those are [how-tos](../how-to)
- Opinions, rationale, comparisons — those are [explanation](../explanation)
- Worked examples designed to teach — those are [tutorials](../tutorials)

## Writing rules for this section

- Describe, do not discuss.
- Mirror the structure of the thing being documented (one page per CLI
  subcommand, one page per port, etc.).
- Use a consistent pattern across pages of the same kind.
- Examples are illustration, not explanation.
- No "you should", no "we recommend", no narrative.

{{< cards >}}
  {{< card link="cli" title="CLI" subtitle="Subcommands, flags, exit codes." >}}
  {{< card link="ports" title="Application ports" subtitle="Interfaces declared by internal/app." >}}
  {{< card link="errors" title="Sentinel errors" subtitle="Canonical errors from internal/cerr." >}}
{{< /cards >}}
