package platform_events

import (
	"context"
	"encoding/json"

	"github.com/Streamfony/lib-logger/logger"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

type Publisher struct {
	factory   *EventFactory
	publisher message.Publisher
	logger    logger.Logger
}

func NewPublisher(platform, platformType, pubsubAddress string, logger logger.Logger) (*Publisher, error) {
	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{pubsubAddress},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		factory:   NewFactory(platform, platformType),
		publisher: publisher,
		logger:    logger.Named("publisher.platform_events"),
	}, nil
}

func (p *Publisher) Publish(ctx context.Context, event Event) error {
	json, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := message.NewMessage(uuid.New().String(), json)

	return p.publisher.Publish(topic, msg)
}

func (p *Publisher) PublishSubscribeEvent(ctx context.Context, userID uint64, additionalInfo *string) error {
	p.logger.Debug("publishing subscribe event", logger.F("user_id", userID), logger.F("additional_info", *additionalInfo))
	event := p.factory.SubscribeEvent(userID, additionalInfo)
	return p.Publish(ctx, event)
}

func (p *Publisher) PublishUnsubscribeEvent(ctx context.Context, userID uint64, additionalInfo *string, err error) error {
	p.logger.Debug("publishing unsubscribe event", logger.F("user_id", userID), logger.F("additional_info", *additionalInfo), logger.Err(err))
	event := p.factory.UnsubscribeEvent(userID, additionalInfo, err)
	return p.Publish(ctx, event)
}

func (p *Publisher) Close() error {
	return p.publisher.Close()
}
