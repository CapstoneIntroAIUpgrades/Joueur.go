package amazons

// Player is a player in this game. Every AI controls one player.
type Player interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// ClientType is what type of client this is, e.g. 'Python',
	// 'JavaScript', or some other language. For potential data mining
	// purposes.
	ClientType() string

	// Lost is if the player lost the game or not.
	Lost() bool

	// Name is the name of the player.
	Name() string

	// Piece is the color (side) of this player. Either 'Q' or 'q',
	// with the 'Q' player having the first move.
	//
	// Literal Values: "Q" or "q"
	Piece() string

	// ReasonLost is the reason why the player lost the game.
	ReasonLost() string

	// ReasonWon is the reason why the player won the game.
	ReasonWon() string

	// TimeRemaining is the amount of time (in ns) remaining for this
	// AI to send commands.
	TimeRemaining() float64

	// Won is if the player won the game or not.
	Won() bool
}
