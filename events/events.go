package events

import (
	"encoding/json"
)

type Event struct {
	Platform string         `json:"platform"`
	Event    string         `json:"event"`
	UserID   uint64         `json:"user_id"`
	Data     map[string]any `json:"data"`
}

func (e *Event) Json() []byte {
	jsonEvent, _ := json.Marshal(e)

	return jsonEvent
}

const (
	connected    = "connected"
	disconnected = "disconnected"
	nodeStarted  = "node_started"
	nodeStopped  = "node_stopped"
)

func Connected(platform string, userID uint64) Event {
	return Event{
		Platform: platform,
		Event:    connected,
		Data:     map[string]any{},
		UserID:   userID,
	}
}

func Disconnected(platform string, userID uint64) Event {
	return Event{
		Platform: platform,
		Event:    disconnected,
		Data:     map[string]any{},
		UserID:   userID,
	}
}

func NodeStarted(platform string, instanceUrl string) Event {
	return Event{
		Platform: platform,
		Event:    nodeStarted,
		Data: map[string]any{
			"instance_url": instanceUrl,
		},
		UserID: 0,
	}
}

func NodeStopped(platform string, instanceUrl string) Event {
	return Event{
		Platform: platform,
		Event:    nodeStopped,
		Data: map[string]any{
			"instance_url": instanceUrl,
		},
		UserID: 0,
	}
}
