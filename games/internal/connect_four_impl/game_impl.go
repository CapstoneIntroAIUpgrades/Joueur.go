package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/connect_four"
)

// GameImpl is the struct that implements the container for Game instances
// in ConnectFour.
type GameImpl struct {
	base.GameImpl

	colsImpl        int64
	gameObjectsImpl map[string]connectfour.GameObject
	playersImpl     []connectfour.Player
	repStringImpl   string
	rowsImpl        int64
	sessionImpl     string
}

// Cols returns the number of tiles on the board along the y (vertical)
// axis.
func (gameImpl *GameImpl) Cols() int64 {
	return gameImpl.colsImpl
}

// GameObjects returns a mapping of every game object's ID to the actual
// game object. Primarily used by the server and client to easily refer to
// the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]connectfour.GameObject {
	return gameImpl.gameObjectsImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []connectfour.Player {
	return gameImpl.playersImpl
}

// RepString returns a string describing all of the information necessary
// to fully represent the game's state.
func (gameImpl *GameImpl) RepString() string {
	return gameImpl.repStringImpl
}

// Rows returns the number of cells on the board along the x (horizontal)
// axis.
func (gameImpl *GameImpl) Rows() int64 {
	return gameImpl.rowsImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.colsImpl = 0
	gameImpl.gameObjectsImpl = make(map[string]connectfour.GameObject)
	gameImpl.playersImpl = []connectfour.Player{}
	gameImpl.repStringImpl = ""
	gameImpl.rowsImpl = 0
	gameImpl.sessionImpl = ""
}

// DeltaMerge merges the delta for a given attribute in Game.
func (gameImpl *GameImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := gameImpl.GameImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	connectfourDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'connectfour.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "cols":
		gameImpl.colsImpl = connectfourDeltaMerge.Int(delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = connectfourDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = connectfourDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "repString":
		gameImpl.repStringImpl = connectfourDeltaMerge.String(delta)
		return true, nil
	case "rows":
		gameImpl.rowsImpl = connectfourDeltaMerge.Int(delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = connectfourDeltaMerge.String(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}
