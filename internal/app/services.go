package app

import (
	"context"

	"github.com/mdhender/opyl/internal/domain"
)

// Services aggregates the ports a fully-wired opyl pipeline needs.
// Runtime constructs one Services with concrete adapters injected and
// passes it to delivery/cli handlers.
type Services struct {
	Orders     OrderSource
	Store      GameStateStore
	Renderer   ReportRenderer
	Dispatcher ReportDispatcher
	Ledger     TurnLedger
	Clock      Clock
}

// IngestOrders reads and validates player orders for (gameID, turn).
// TODO: implement.
func (s *Services) IngestOrders(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber) error {
	return nil
}

// ProcessTurn resolves the turn deterministically from validated orders
// and the prior turn's state, persists the new state, and marks the
// turn processed in the ledger.
// TODO: implement.
func (s *Services) ProcessTurn(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber) error {
	return nil
}

// RenderReports produces per-player report artifacts for a processed
// turn using the configured renderer.
// TODO: implement.
func (s *Services) RenderReports(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber) error {
	return nil
}

// DispatchReports delivers rendered reports to each player's recipient
// address using the configured dispatcher.
// TODO: implement.
func (s *Services) DispatchReports(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber) error {
	return nil
}

// RunPipeline executes ingest → process → render → dispatch end-to-end.
// Each stage remains independently re-runnable so an operator can
// resume mid-pipeline after fixing a problem.
// TODO: implement.
func (s *Services) RunPipeline(ctx context.Context, gameID domain.GameID, turn domain.TurnNumber) error {
	return nil
}
