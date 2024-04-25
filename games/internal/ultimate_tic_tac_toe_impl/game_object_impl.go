package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/ultimate_tic_tac_toe"
)

// GameObjectImpl is the struct that implements the container for
// GameObject instances in UltimateTicTacToe.
type GameObjectImpl struct {
	base.GameObjectImpl

	game *GameImpl

	gameObjectNameImpl string
	logsImpl           []string
}

// Game returns a pointer to the UltimateTicTacToe Game instance
func (gameObjectImpl *GameObjectImpl) Game() ultimatetictactoe.Game {
	return gameObjectImpl.game
}

// GameObjectName returns string representing the top level Class that this
// game object is an instance of. Used for reflection to create new
// instances on clients, but exposed for convenience should AIs want this
// data.
func (gameObjectImpl *GameObjectImpl) GameObjectName() string {
	return gameObjectImpl.gameObjectNameImpl
}

// Logs returns any strings logged will be stored here. Intended for
// debugging.
func (gameObjectImpl *GameObjectImpl) Logs() []string {
	return gameObjectImpl.logsImpl
}

// Log runs logic that adds a message to this GameObject's logs. Intended
// for your own debugging purposes, as strings stored here are saved in the
// gamelog.
func (gameObjectImpl *GameObjectImpl) Log(message string) {
	gameObjectImpl.RunOnServer("log", map[string]interface{}{
		"message": message,
	})
}

// InitImplDefaults initializes safe defaults for all fields in GameObject.
func (gameObjectImpl *GameObjectImpl) InitImplDefaults() {
	gameObjectImpl.GameObjectImpl.InitImplDefaults()

	gameObjectImpl.logsImpl = []string{}
}

// DeltaMerge merges the delta for a given attribute in GameObject.
func (gameObjectImpl *GameObjectImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := gameObjectImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	ultimatetictactoeDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'ultimatetictactoe.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "logs":
		gameObjectImpl.logsImpl = ultimatetictactoeDeltaMerge.ArrayOfString(&gameObjectImpl.logsImpl, delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}
