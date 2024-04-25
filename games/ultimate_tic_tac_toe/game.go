package ultimateTicTacToe

import "joueur/base"

// Game is tic Tac Toe but on nine boards.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// Cols is the number of tiles on the board along the y (vertical)
	// axis.
	Cols() int64

	// GameObjects is a mapping of every game object's ID to the
	// actual game object. Primarily used by the server and client to
	// easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// Players is array of all the players in the game.
	Players() []Player

	// RepString is a string describing all of the information
	// necessary to fully represent the game's state.
	RepString() string

	// Rows is the number of cells on the board along the x
	// (horizontal) axis.
	Rows() int64

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string
}
