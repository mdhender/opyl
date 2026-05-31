// Package cli is opyl's delivery layer: thin CLI subcommand handlers
// that parse flags/env, call exactly one app use case, map errors to
// exit codes, and print operator-visible output.
//
// SOUSA: handlers stay thin (parse → call use case → format). No SQL,
// no file parsing, no rendering, no email here. Those belong in infra.
//
// Planned subcommands (one handler per file):
//
//	ingest    → app.Services.IngestOrders
//	process   → app.Services.ProcessTurn
//	render    → app.Services.RenderReports
//	dispatch  → app.Services.DispatchReports
//	pipeline  → app.Services.RunPipeline
package cli
