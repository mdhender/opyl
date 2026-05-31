// Command opyl is the runtime composition root for the opyl turn engine.
//
// This binary is the only place that knows about every layer. It parses
// configuration, constructs concrete adapters, injects them into the
// application services, and dispatches the requested CLI subcommand.
//
// SOUSA: this file is Runtime. Keep wiring here; do not host business
// rules. Subcommand argument parsing belongs in internal/delivery/cli;
// use case logic belongs in internal/app.
package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "opyl:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		usage(os.Stderr)
		return nil
	}
	// TODO: wire concrete adapters (orderfile, store, render, mail) into
	// app.Services and dispatch to internal/delivery/cli.
	return fmt.Errorf("subcommand %q not yet implemented", args[0])
}

func usage(w *os.File) {
	fmt.Fprintln(w, `opyl — turn-based game engine

Usage:
  opyl <subcommand> [flags]

Planned subcommands:
  ingest     Read player order files for a game/turn
  process    Resolve a turn deterministically from validated orders
  render     Produce per-player report artifacts (text or PDF)
  dispatch   Deliver rendered reports to recipients
  pipeline   Run ingest → process → render → dispatch end-to-end

See README.md and AGENTS.md for project layout and SOUSA conventions.`)
}
