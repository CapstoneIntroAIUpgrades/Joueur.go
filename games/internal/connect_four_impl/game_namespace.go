// Package impl contains all the ConnectFour game implimentation logic.
package impl

// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/connect_four"
)

// ConnectFourNamespace is the collection of implimentation logic for the ConnectFour game.
type ConnectFourNamespace struct{}

// Name returns the name of the ConnectFour game.
func (*ConnectFourNamespace) Name() string {
	return "ConnectFour"
}

// Version returns the current version hash as last generated for the ConnectFour game.
func (*ConnectFourNamespace) Version() string {
	return "ff3da34375459be10b77fd8f82b24d9880f0a100e48e2202812118bfcfd42e1e"
}

// PlayerName returns the desired name of the AI in the ConnectFour game.
func (*ConnectFourNamespace) PlayerName() string {
	return connectfour.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the ConnectFour game.
func (*ConnectFourNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch gameObjectName {
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InitImplDefaults()
		return &newGameObject, nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InitImplDefaults()
		return &newPlayer, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'ConnectFour' can be created")
}

// CreateGame is the factory method for Game the ConnectFour game.
func (*ConnectFourNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the ConnectFour game.
func (*ConnectFourNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := connectfour.AI{}
	return &ai, &ai.AIImpl
}

func (*ConnectFourNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the ConnectFour game.
func (*ConnectFourNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*connectfour.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid connectfour.AI to order!")
	}
	switch functionName {
	case "makeMove":
		return (*ai).MakeMove(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}
