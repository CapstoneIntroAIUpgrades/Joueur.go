package impl

import (
	"joueur/base"
	"joueur/games/ultimateTicTacToe"
)

// DeltaMerge is the set of functions that can delta merge a
// UltimateTicTacToe game.
type DeltaMerge interface {
	base.DeltaMerge

	GameObject(interface{}) ultimatetictactoe.GameObject
	Player(interface{}) ultimatetictactoe.Player

	ArrayOfPlayer(*[]ultimatetictactoe.Player, interface{}) []ultimatetictactoe.Player
	ArrayOfString(*[]string, interface{}) []string
	MapOfStringToGameObject(*map[string]ultimatetictactoe.GameObject, interface{}) map[string]ultimatetictactoe.GameObject
}

// DeltaMergeImpl is the implimentation struct for the DeltaMerge interface.
type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- UltimateTicTacToe Game Object Deltas -- \\

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) ultimateTicTacToe.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(ultimateTicTacToe.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("ultimateTicTacToe.GameObject", delta)
	}

	return asGameObject
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) ultimateTicTacToe.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(ultimateTicTacToe.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("ultimateTicTacToe.Player", delta)
	}

	return asPlayer
}

// -- Deep Deltas -- \\

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]ultimatetictactoe.Player, delta interface{}) []ultimatetictactoe.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]ultimatetictactoe.Player, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Player(deltaValue)
	}
	return newArray
}

// ArrayOfString delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfString(state *[]string, delta interface{}) []string {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]string, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.String(deltaValue)
	}
	return newArray
}

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]ultimatetictactoe.GameObject, delta interface{}) map[string]ultimatetictactoe.GameObject {
	deltaMap := (*deltaMergeImpl).ToDeltaMap(delta)
	for deltaKey, deltaValue := range deltaMap {
		if (*deltaMergeImpl).IsDeltaRemoved(deltaValue) {
			delete(*state, deltaKey)
		} else {
			(*state)[deltaKey] = deltaMergeImpl.GameObject(deltaValue)
		}
	}
	return *state
}
