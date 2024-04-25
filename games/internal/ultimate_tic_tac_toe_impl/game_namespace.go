// Package impl contains all the UltimateTicTacToe game implimentation logic.
package impl

// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/ultimate_tic_tac_toe"
)

// UltimateTicTacToeNamespace is the collection of implimentation logic for the UltimateTicTacToe game.
type UltimateTicTacToeNamespace struct{}

// Name returns the name of the UltimateTicTacToe game.
func (*UltimateTicTacToeNamespace) Name() string {
	return "UltimateTicTacToe"
}

// Version returns the current version hash as last generated for the UltimateTicTacToe game.
func (*UltimateTicTacToeNamespace) Version() string {
	return "50e7a74ecd23f8e98bbe235fc2aa7db662a607ebdeb59ad3e5a4213cff4f8a43"
}

// PlayerName returns the desired name of the AI in the UltimateTicTacToe game.
func (*UltimateTicTacToeNamespace) PlayerName() string {
	return ultimatetictactoe.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the UltimateTicTacToe game.
func (*UltimateTicTacToeNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
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
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'UltimateTicTacToe' can be created")
}

// CreateGame is the factory method for Game the UltimateTicTacToe game.
func (*UltimateTicTacToeNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the UltimateTicTacToe game.
func (*UltimateTicTacToeNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := ultimatetictactoe.AI{}
	return &ai, &ai.AIImpl
}

func (*UltimateTicTacToeNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the UltimateTicTacToe game.
func (*UltimateTicTacToeNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*ultimatetictactoe.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid ultimatetictactoe.AI to order!")
	}
	switch functionName {
	case "makeMove":
		return (*ai).MakeMove(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}
