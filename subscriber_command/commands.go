package subscriber_command

type Action string

const (
	ActionSubscribe   Action = "subscribe"
	ActionUnsubscribe Action = "unsubscribe"
)

type Command struct {
	Platform     string `json:"platform"`
	PlatformType string `json:"platform_type"`

	Action         Action  `json:"action"`
	UserID         uint64  `json:"user_id"`
	AdditionalInfo *string `json:"additional_info"`

	Error error `json:"error"`
}

type CommandFactory struct {
	platform     string
	platformType string
}

func NewFactory(platform, platformType string) *CommandFactory {
	return &CommandFactory{
		platform:     platform,
		platformType: platformType,
	}
}

func (c *CommandFactory) SubscribeCommand(userID uint64, additionalInfo *string) Command {
	return Command{
		Platform:       c.platform,
		PlatformType:   c.platformType,
		Action:         ActionSubscribe,
		UserID:         userID,
		AdditionalInfo: additionalInfo,
	}
}

func (c *CommandFactory) UnsubscribeCommand(userID uint64, additionalInfo *string, err error) Command {
	return Command{
		Platform:       c.platform,
		PlatformType:   c.platformType,
		Action:         ActionUnsubscribe,
		UserID:         userID,
		AdditionalInfo: additionalInfo,
		Error:          err,
	}
}
