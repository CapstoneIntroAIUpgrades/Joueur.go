package impl

import (
	"errors"
	"joueur/base"
)

// PlayerImpl is the struct that implements the container for Player
// instances in ConnectFour.
type PlayerImpl struct {
	GameObjectImpl

	clientTypeImpl    string
	colorImpl         string
	lostImpl          bool
	nameImpl          string
	reasonLostImpl    string
	reasonWonImpl     string
	timeRemainingImpl float64
	wonImpl           bool
}

// ClientType returns what type of client this is, e.g. 'Python',
// 'JavaScript', or some other language. For potential data mining
// purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.clientTypeImpl
}

// Color returns the color (side) of this player. Either 'red' or 'yellow',
// with the 'red' player having the first move.
//
// Literal Values: "r" or "y"
func (playerImpl *PlayerImpl) Color() string {
	return playerImpl.colorImpl
}

// Lost returns if the player lost the game or not.
func (playerImpl *PlayerImpl) Lost() bool {
	return playerImpl.lostImpl
}

// Name returns the name of the player.
func (playerImpl *PlayerImpl) Name() string {
	return playerImpl.nameImpl
}

// ReasonLost returns the reason why the player lost the game.
func (playerImpl *PlayerImpl) ReasonLost() string {
	return playerImpl.reasonLostImpl
}

// ReasonWon returns the reason why the player won the game.
func (playerImpl *PlayerImpl) ReasonWon() string {
	return playerImpl.reasonWonImpl
}

// TimeRemaining returns the amount of time (in ns) remaining for this AI
// to send commands.
func (playerImpl *PlayerImpl) TimeRemaining() float64 {
	return playerImpl.timeRemainingImpl
}

// Won returns if the player won the game or not.
func (playerImpl *PlayerImpl) Won() bool {
	return playerImpl.wonImpl
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.clientTypeImpl = ""
	playerImpl.colorImpl = ""
	playerImpl.lostImpl = true
	playerImpl.nameImpl = ""
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.timeRemainingImpl = 0
	playerImpl.wonImpl = true
}

// DeltaMerge merges the delta for a given attribute in Player.
func (playerImpl *PlayerImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := playerImpl.GameObjectImpl.DeltaMerge(
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
	case "clientType":
		playerImpl.clientTypeImpl = connectfourDeltaMerge.String(delta)
		return true, nil
	case "color":
		playerImpl.colorImpl = connectfourDeltaMerge.String(delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = connectfourDeltaMerge.Boolean(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = connectfourDeltaMerge.String(delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = connectfourDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = connectfourDeltaMerge.String(delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = connectfourDeltaMerge.Float(delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = connectfourDeltaMerge.Boolean(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}
