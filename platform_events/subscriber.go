package platform_events

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"
)

type Subscriber struct {
	consumer message.Subscriber
	logger   *zap.Logger
}

const topic = "platform_connections_events"

func NewSubscriber(pubsubAddress string, pubsubGroup string, logger *zap.Logger) (*Subscriber, error) {
	consumer, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:       []string{pubsubAddress},
			ConsumerGroup: pubsubGroup,
		},
		nil,
		nil,
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return nil, err
	}

	return &Subscriber{
		consumer: consumer,
		logger:   logger.Named("subscriber.platform_events"),
	}, nil
}

func (s *Subscriber) Subscribe(ctx context.Context, handler func(ctx context.Context, event Event) error) error {
	messages, err := s.consumer.Subscribe(ctx, topic)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			s.logger.Info("context cancelled, stopping subscriber")
			return ctx.Err()
		case msg, ok := <-messages:
			if !ok {
				s.logger.Info("messages channel closed")
				return nil
			}

			var event Event
			if err := json.Unmarshal(msg.Payload, &event); err != nil {
				s.logger.Error("failed to unmarshal event", zap.Error(err))
				continue
			}

			if err := handler(ctx, event); err != nil {
				s.logger.Error("failed to handle event", zap.Error(err))
				continue
			}

			msg.Ack()
		}
	}
}
