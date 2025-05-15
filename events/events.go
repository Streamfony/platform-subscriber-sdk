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
