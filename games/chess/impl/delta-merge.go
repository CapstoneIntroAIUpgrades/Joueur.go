package impl

import (
	"joueur/base"
	"joueur/games/chess"
)

type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Chess Game Object Deltas -- \\

func (deltaMergeImpl DeltaMergeImpl) GameObject(delta interface{}) chess.GameObject {
	baseGameObject := deltaMergeImpl.BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(chess.GameObject)
	if !isGameObject {
		deltaMergeImpl.CannotConvertDeltaTo("chess.GameObject", delta)
	}

	return asGameObject
}

func (deltaMergeImpl DeltaMergeImpl) Player(delta interface{}) chess.Player {
	baseGameObject := deltaMergeImpl.BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(chess.Player)
	if !isPlayer {
		deltaMergeImpl.CannotConvertDeltaTo("chess.Player", delta)
	}

	return asPlayer
}

// -- Deep Deltas -- \\

func (deltaMergeImpl DeltaMergeImpl) ArrayOfPlayer(state *[]chess.Player, delta interface{}) *[]chess.Player {
	deltaList, listLength := deltaMergeImpl.ToDeltaArray(delta)
	newArray := make([]chess.Player, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Player(deltaValue)
	}
	return &newArray
}

func (deltaMergeImpl DeltaMergeImpl) ArrayOfString(state *[]string, delta interface{}) *[]string {
	deltaList, listLength := deltaMergeImpl.ToDeltaArray(delta)
	newArray := make([]string, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.String(deltaValue)
	}
	return &newArray
}

func (deltaMergeImpl DeltaMergeImpl) MapOfStringToGameObject(state *map[string]chess.GameObject, delta interface{}) *map[string]chess.GameObject {
	deltaMap := deltaMergeImpl.ToDeltaMap(delta)
	for deltaKey, deltaValue := range deltaMap {
		if deltaMergeImpl.IsDeltaRemoved(deltaValue) {
			delete(*state, deltaKey)
		} else {
			(*state)[deltaKey] = deltaMergeImpl.GameObject(deltaValue)
		}
	}
	return state
}