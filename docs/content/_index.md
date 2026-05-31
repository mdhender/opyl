---
title: opyl documentation
toc: false
---

opyl is a batch-style, turn-based game engine in Go. It ingests player order
files, resolves turns deterministically, renders per-player reports as text or
PDF, and dispatches them by email.

This documentation is organised by the [Diataxis](https://diataxis.fr) framework.
Pick the section that matches what you are trying to do right now.

{{< cards >}}
  {{< card link="tutorials" title="Tutorials" icon="academic-cap"
        subtitle="Learning-oriented lessons. Start here if you are new to opyl and want a guided experience." >}}
  {{< card link="how-to" title="How-to guides" icon="cog"
        subtitle="Task-oriented recipes. Use these when you already know what you want to accomplish." >}}
  {{< card link="reference" title="Reference" icon="document-text"
        subtitle="Information-oriented descriptions of commands, configuration, file formats, and ports." >}}
  {{< card link="explanation" title="Explanation" icon="light-bulb"
        subtitle="Understanding-oriented discussion. Background, design decisions, and the SOUSA layering." >}}
{{< /cards >}}

## Where to start

| If you want to…                                       | Read                                  |
| ----------------------------------------------------- | ------------------------------------- |
| Run your first turn end-to-end                        | [Tutorials](tutorials)                |
| Wire a new infra adapter, render a one-off report     | [How-to guides](how-to)               |
| Look up a CLI flag, port signature, or error code     | [Reference](reference)                |
| Understand why opyl is layered the way it is          | [Explanation](explanation)            |
