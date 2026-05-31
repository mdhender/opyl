// Package domain holds opyl's pure game types, value objects, and
// invariants. It is the inner-most SOUSA layer.
//
// Domain must remain framework-free and deterministic. No I/O, no
// filesystem, no clock, no randomness, no PDF or email concerns.
package domain

// GameID identifies a single game instance.
type GameID string

// TurnNumber is the sequential turn counter within a game (1-based).
type TurnNumber int

// PlayerID identifies a single player within a game.
type PlayerID string

// OrderBundle is the set of orders a single player submitted for a turn,
// after parsing from whatever flat-file format infra produced them in.
type OrderBundle struct {
	PlayerID PlayerID
	// Parsed, validated orders go here. The raw bytes never reach domain;
	// they are decoded at the infra/orderfile boundary.
}

// GameState holds the authoritative world state at a specific turn.
type GameState struct {
	GameID GameID
	Turn   TurnNumber
	// Game-specific state fields go here.
}

// PlayerReport is a per-player view of a processed turn, ready to be
// rendered into bytes by an infra/render adapter.
type PlayerReport struct {
	GameID   GameID
	Turn     TurnNumber
	PlayerID PlayerID
	// Renderable content goes here.
}

// Recipient is the destination for a dispatched report.
type Recipient struct {
	PlayerID PlayerID
	Email    string
}

// Attachment is a rendered report ready to be delivered.
type Attachment struct {
	Filename string
	MIMEType string
	Body     []byte
}
