package platform_events

type Action string

const (
	ActionSubscribe   Action = "subscribe"
	ActionUnsubscribe Action = "unsubscribe"
)

type Event struct {
	Platform     string `json:"platform"`
	PlatformType string `json:"platform_type"`

	Action         Action  `json:"action"`
	UserID         uint64  `json:"user_id"`
	AdditionalInfo *string `json:"additional_info"`

	Error error `json:"error"`
}

type EventFactory struct {
	platform     string
	platformType string
}

func NewFactory(platform, platformType string) *EventFactory {
	return &EventFactory{
		platform:     platform,
		platformType: platformType,
	}
}

func (c *EventFactory) SubscribeEvent(userID uint64, additionalInfo *string) Event {
	return Event{
		Platform:       c.platform,
		PlatformType:   c.platformType,
		Action:         ActionSubscribe,
		UserID:         userID,
		AdditionalInfo: additionalInfo,
	}
}

func (c *EventFactory) UnsubscribeEvent(userID uint64, additionalInfo *string, err error) Event {
	return Event{
		Platform:       c.platform,
		PlatformType:   c.platformType,
		Action:         ActionUnsubscribe,
		UserID:         userID,
		AdditionalInfo: additionalInfo,
		Error:          err,
	}
}
