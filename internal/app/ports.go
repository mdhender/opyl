// Package app holds opyl's use cases and the ports (interfaces) it
// declares for outer layers to implement.
//
// SOUSA: app may import domain and cerr and the standard library. It
// must not import infra, delivery, or runtime, and must not reference
// concrete drivers, file formats, PDF libraries, or SMTP libraries.
package app

import (
	"context"
	"io"

	"github.com/mdhender/opyl/internal/domain"
)

// OrderSource reads player order bundles for a given game and turn.
//
// Implementations live in infra/orderfile (flat files) or potentially
// future adapters (e.g. IMAP poller). Parsing and shape validation
// happen inside the adapter; only typed OrderBundles reach app.
type OrderSource interface {
	ReadOrders(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber) ([]domain.OrderBundle, error)
}

// GameStateStore loads and saves authoritative game state per turn.
//
// Implementations live in infra/store. The storage choice (SQLite, a
// directory of versioned files, etc.) is an infra concern.
type GameStateStore interface {
	Load(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber) (*domain.GameState, error)
	Save(ctx context.Context, state *domain.GameState) error
}

// ReportRenderer turns a per-player report into bytes plus a MIME type.
//
// Implementations live in infra/render/text and infra/render/pdf. The
// renderer does not know about delivery channels.
type ReportRenderer interface {
	Render(ctx context.Context, report *domain.PlayerReport, w io.Writer) (mimeType string, err error)
}

// ReportDispatcher delivers a rendered attachment to a recipient.
//
// Implementations live in infra/mail. The dispatcher does not know how
// the attachment was produced.
type ReportDispatcher interface {
	Dispatch(ctx context.Context, recipient domain.Recipient, attachment domain.Attachment) error
}

// TurnLedger records turn-processing events to make ProcessTurn
// idempotent across operator reruns.
type TurnLedger interface {
	AlreadyProcessed(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber, inputHash string) (bool, error)
	Record(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber, inputHash string) error
}

// Clock abstracts time so use cases stay deterministic in tests.
type Clock interface {
	NowUnix() int64
}

// RNG is a deterministic random source used by use cases that involve
// dice or stochastic decisions.
//
// Implementations live in infra/prng. The choice of generator (PCG,
// Xoshiro, etc.) and the mechanism for deriving independent substreams
// are infra concerns. Persistence of RNG state across turns is also an
// infra concern: GameStateStore round-trips RNG state alongside game
// state, and the port here intentionally does not expose marshal/scan.
//
// Splitting a master generator into per-game, per-stage, or per-player
// substreams is currently a wiring-time operation performed by Runtime
// on the concrete adapter; promote a Split method to this port only
// when a use case demonstrably needs mid-execution fan-out.
type RNG interface {
	// Roll returns a uniformly distributed integer in [low, high].
	Roll(low, high int) int
	// RollDice returns the sum of rolling n dice each with the given
	// number of sides.
	RollDice(n, sides int) int
}
