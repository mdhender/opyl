---
title: SOUSA in opyl
weight: 1
prev: /explanation
---

opyl is the fifth project to adopt SOUSA and the first to apply it outside
the web-app shape (Diacous, EC, GemGem, Hokey were all variations on Go
backend + SQL + browser frontend). This page explains why SOUSA still fits
— arguably better — when the project is a batch CLI that reads files and
sends email.

## The shape problem

opyl has unusually heavy outer-layer concerns for its size: flat-file
parsing, PDF generation, SMTP, persistent state. It has unusually little
outer surface area: a handful of CLI subcommands. That is exactly the
shape where SOUSA's port discipline pays off most — each heavy external
concern lives behind a tiny port and can be swapped without touching
game rules.

## What we kept

- **Five layers**, same as the web projects: Domain, Application,
  Infrastructure, Delivery, Runtime.
- **Application owns the ports.** Infra implements; app never imports a
  concrete adapter.
- **Sentinel errors in `internal/cerr`**, expressing business meaning.
- **Delivery stays thin** — parse args, call one use case, format output.

## What we changed

- **No auth layer.** The only user is a trusted operator. The auth
  chapter of the unified SOUSA skill simply does not apply.
- **Delivery is tiny.** No HTTP, no SPA. A handful of CLI subcommands.
- **Idempotency moved to Application.** Operators rerun things. A
  `TurnLedger` port plus an input hash makes use cases re-runnable.
  Infra adapters stay dumb.
- **Render and dispatch are two ports.** Combining them would make
  dry-runs (render without sending) hard. Keeping them split also lets
  the dispatcher not care how the bytes were produced.
- **Input files are untrusted.** This rule existed implicitly for HTTP
  handlers in the web projects; for opyl we make it explicit at the
  `infra/orderfile/` boundary.

## What is deferred

Several decisions are intentionally open because SOUSA lets us defer them
behind ports:

- State storage: SQLite or a directory of per-turn JSON snapshots
- PDF library: `gofpdf`, `typst`, `chromedp`, or none
- Mail transport: SMTP, SES, or "drop EML files in /outbox"
- Order file format: custom DSL, YAML, or structured email body

Each lives entirely inside its `internal/infra/<adapter>/` package. If a
choice starts requiring changes outside its infra package, the port
boundary is wrong — fix the port first.
