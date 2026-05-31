// Package store is the infra adapter that persists game state across
// turns, implementing app.GameStateStore and app.TurnLedger.
//
// SOUSA: the storage choice (SQLite, a directory of versioned
// JSON/YAML files, etc.) is an infra-only decision. It must not leak
// driver types, file paths, or transaction objects to app or domain.
// Return typed domain values and cerr sentinels upward.
package store
