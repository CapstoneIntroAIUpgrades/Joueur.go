package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"joueur/base/errorhandler"

	"github.com/fatih/color"
)

type BaseEvent struct {
	EventName string `json:"event"`
}

const readSize = 1024

var eventsStack = make([][]byte, 0)
var receivedBuffer = make([]byte, 0)

func waitForEvents() {
	if len(eventsStack) > 0 {
		return
	}

	for {
		sent := make([]byte, readSize)
		bytesSent, err := instance.conn.Read(sent)
		if err != nil {
			errorhandler.HandleError(
				errorhandler.CannotReadSocket,
				err,
				"Error reading socket while waiting for events",
			)
		}

		if bytesSent == 0 {
			continue
		}

		if instance.printIO {
			color.Magenta("FROM SERVER <-- " + string(sent))
		}

		split := bytes.Split(sent, []byte{eotChar})
		// the last item will either be "" if the last char was an EOT_CHAR,
		//	or a partial data we need to buffer anyways
		for i, eventBytes := range split {
			if i == len(split)-1 {
				receivedBuffer = eventBytes // left over bytes after the last EOT char for next event
				continue
			}

			if i == 0 {
				eventBytes = append(receivedBuffer, eventBytes...)
				receivedBuffer = make([]byte, 0)
			}

			eventsStack = append(eventsStack, eventBytes)
		}

		if len(eventsStack) > 0 {
			return
		}
	}
}

func WaitForEvent(eventName string, destination *interface{}) {
	for {
		waitForEvents()

		for len(eventsStack) > 0 {
			// pop first event off the front of the events stack
			eventBytes := eventsStack[0]
			eventsStack = eventsStack[1:]
			var baseEvent *BaseEvent = nil
			nameErr := json.Unmarshal(eventBytes, &baseEvent)

			if baseEvent == nil {
				nameErr = errors.New("No parsed base event")
			}

			if nameErr != nil {
				errorhandler.HandleError(
					errorhandler.MalformedJSON,
					nameErr,
					"Could not parse base JSON"+string(eventBytes),
				)
			}

			if baseEvent.EventName == eventName {
				err := json.Unmarshal(eventBytes, &destination)

				if destination == nil {
					err = errors.New("No destination data")
				}

				if err != nil {
					errorhandler.HandleError(
						errorhandler.MalformedJSON,
						err,
						"Error occurred while waiting for "+eventName,
					)
				}

				return // destination should not be populated for them to deal with
			} else { // attempt to auto handle the event
				switch baseEvent.EventName {
				case "fatal":
					autoHandleEventFatal(eventBytes)
				default:
					errorhandler.HandleError(
						errorhandler.UnknownEventFromServer,
						errors.New("No event auto handler for "+baseEvent.EventName),
						"Unknown event could not be handled",
					)
				}
			}
		}
	}
}