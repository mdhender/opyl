---
title: Sentinel errors
weight: 3
prev: /reference/ports
---

Canonical errors declared in `internal/cerr`. Outer layers translate these
into exit codes, operator-visible messages, or log entries.

| Error                       | Meaning                                                   |
| --------------------------- | --------------------------------------------------------- |
| `ErrGameNotFound`           | No game with that id exists in the configured store      |
| `ErrTurnNotFound`           | The requested turn number has no recorded state          |
| `ErrTurnAlreadyProcessed`   | The turn was already processed with the same input hash  |
| `ErrInvalidOrders`          | Order file failed shape validation at the parser boundary|
| `ErrPlayerNotFound`         | No player with that id exists in the game                |
| `ErrReportNotReady`         | A report was requested for a turn that has not rendered  |
