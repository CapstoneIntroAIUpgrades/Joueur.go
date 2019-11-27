package client

import (
	"encoding/json"
	"errors"
	"joueur/base/errorhandler"
	"time"

	"github.com/fatih/color"
)

type SendEventDataStringShape struct {
	Event    string `json:"event"`
	SentTime int64  `json:"sentTime"`
	Data     string `json:"data"`
}

func SendEventDataString(event string, data string) {
	bytes, err := json.Marshal(SendEventDataStringShape{
		Event:    event,
		Data:     data,
		SentTime: time.Now().Unix(),
	})

	if err != nil {
		errorhandler.HandleError(
			errorhandler.MalformedJSON,
			err,
			"Could not encode event to json",
		)
	}

	sendRaw(append(bytes, eotChar))
}

func sendRaw(bytes []byte) error {
	/**
	 * Sends a raw string to the game server
	 * @param str The string to send.
	 */
	if instance.conn == nil {
		return errors.New("Cannot write to socket that has not been initialized")
	}

	if instance.printIO {
		color.Magenta("TO SERVER <-- " + string(bytes))
	}

	_, err := instance.conn.Write((bytes))
	if err != nil {
		errorhandler.HandleError(
			errorhandler.DisconnectedUnexpectedly,
			err,
			"Could not send string through server.",
		)
	}

	return nil
}