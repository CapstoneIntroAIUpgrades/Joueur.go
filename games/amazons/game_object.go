package amazons

import "joueur/base"

// GameObject is an object in the game. The most basic class that all game
// classes should inherit from automatically.
type GameObject interface {
	// Parent interfaces
	base.GameObject

	// -- Attributes -- \\

	// Logs is any strings logged will be stored here. Intended for
	// debugging.
	Logs() []string

	// -- Methods -- \\

	// Log adds a message to this GameObject's logs. Intended for your
	// own debugging purposes, as strings stored here are saved in the
	// gamelog.
	Log(string)
}
