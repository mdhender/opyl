---
title: Application ports
weight: 2
prev: /reference/cli
---

Interfaces declared by `internal/app`. Implementations live in
`internal/infra/`. Use cases depend only on these interfaces.

## `OrderSource`

```go
type OrderSource interface {
    ReadOrders(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber) ([]domain.OrderBundle, error)
}
```

Reads player order bundles for a given game and turn. Implementations parse
flat files and return validated `domain.OrderBundle` values.

| Returns          | Meaning                                              |
| ---------------- | ---------------------------------------------------- |
| `ErrGameNotFound`| The game id does not exist in the source            |
| `ErrInvalidOrders`| One or more order files failed shape validation     |

## `GameStateStore`

```go
type GameStateStore interface {
    Load(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber) (*domain.GameState, error)
    Save(ctx context.Context, state *domain.GameState) error
}
```

Loads and saves authoritative game state per turn.

## `ReportRenderer`

```go
type ReportRenderer interface {
    Render(ctx context.Context, report *domain.PlayerReport, w io.Writer) (mimeType string, err error)
}
```

Turns a per-player report into bytes plus a MIME type. Does not know about
delivery channels.

## `ReportDispatcher`

```go
type ReportDispatcher interface {
    Dispatch(ctx context.Context, recipient domain.Recipient, attachment domain.Attachment) error
}
```

Delivers a rendered attachment to a recipient. Does not know how the
attachment was produced.

## `TurnLedger`

```go
type TurnLedger interface {
    AlreadyProcessed(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber, inputHash string) (bool, error)
    Record(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber, inputHash string) error
}
```

Records turn-processing events to make `ProcessTurn` idempotent across
operator reruns.

## `Clock`

```go
type Clock interface {
    NowUnix() int64
}
```

Abstracts time so use cases stay deterministic in tests.
