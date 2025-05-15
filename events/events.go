package events

type Event struct {
	Platform string         `json:"platform"`
	Event    string         `json:"event"`
	Data     map[string]any `json:"data"`
}

const (
	connected    = "connected"
	disconnected = "disconnected"
)

func Connected(platform string) Event {
	return Event{
		Platform: platform,
		Event:    connected,
		Data:     map[string]any{},
	}
}

func Disconnected(platform string) Event {
	return Event{
		Platform: platform,
		Event:    disconnected,
		Data:     map[string]any{},
	}
}
