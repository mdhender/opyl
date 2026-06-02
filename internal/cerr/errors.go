// Package cerr defines canonical sentinel errors shared across opyl.
//
// Errors here express business meaning, not transport or rendering
// formatting. Outer layers translate them to exit codes, report text,
// or operator-visible diagnostics.
//
// SOUSA: cerr is an inner-layer package. It must not import app, infra,
// delivery, or runtime.
package cerr

const (
	ErrGameNotFound         = Error("game not found")
	ErrTurnNotFound         = Error("turn not found")
	ErrTurnAlreadyProcessed = Error("turn already processed")
	ErrInvalidOrders        = Error("invalid orders")
	ErrPlayerNotFound       = Error("player not found")
	ErrReportNotReady       = Error("report not ready")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
