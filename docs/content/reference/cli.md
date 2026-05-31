---
title: CLI
weight: 1
prev: /reference
---

## Synopsis

```
opyl <subcommand> [flags]
```

## Subcommands

| Subcommand | Purpose                                                   |
| ---------- | --------------------------------------------------------- |
| `ingest`   | Read player order files for a game/turn                   |
| `process`  | Resolve a turn deterministically from validated orders    |
| `render`   | Produce per-player report artifacts (text or PDF)         |
| `dispatch` | Deliver rendered reports to recipients                    |
| `pipeline` | Run `ingest → process → render → dispatch` end-to-end     |

## Exit codes

| Code | Meaning                                          |
| ---: | ------------------------------------------------ |
| `0`  | Subcommand completed successfully                |
| `1`  | Subcommand failed (see stderr for the error)    |

_(per-subcommand flag tables to be added as each subcommand lands)_
