// Package impl contains all the Amazons game implimentation logic.
package impl

// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/amazons"
)

// AmazonsNamespace is the collection of implimentation logic for the Amazons game.
type AmazonsNamespace struct{}

// Name returns the name of the Amazons game.
func (*AmazonsNamespace) Name() string {
	return "Amazons"
}

// Version returns the current version hash as last generated for the Amazons game.
func (*AmazonsNamespace) Version() string {
	return "1b2ac0db4655c0c135ffbf80607a97ab5434b880e1b76f1ade9fdaf44670ae23"
}

// PlayerName returns the desired name of the AI in the Amazons game.
func (*AmazonsNamespace) PlayerName() string {
	return amazons.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Amazons game.
func (*AmazonsNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
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
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Amazons' can be created")
}

// CreateGame is the factory method for Game the Amazons game.
func (*AmazonsNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Amazons game.
func (*AmazonsNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := amazons.AI{}
	return &ai, &ai.AIImpl
}

func (*AmazonsNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Amazons game.
func (*AmazonsNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*amazons.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid amazons.AI to order!")
	}
	switch functionName {
	case "makeMove":
		return (*ai).MakeMove(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}
