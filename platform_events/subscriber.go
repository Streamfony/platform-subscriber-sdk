package platform_events

import (
	"context"
	"encoding/json"

	"github.com/Streamfony/lib-logger/logger"
	watermillLogger "github.com/Streamfony/lib-logger/watermill-logger"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

type Subscriber struct {
	consumer message.Subscriber
	logger   logger.Logger
}

const topic = "platform_connections_events"

func NewSubscriber(pubsubAddress string, pubsubGroup string, logger logger.Logger) (*Subscriber, error) {
	consumer, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:       []string{pubsubAddress},
			ConsumerGroup: pubsubGroup,
			Unmarshaler:   kafka.DefaultMarshaler{},
		},
		watermillLogger.New(logger.Named("subscriber.platform_events")),
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
	l := s.logger.WithFields(logger.F("topic", topic))
	if err != nil {
		l.Error("failed to subscribe to topic",
			logger.Err(err))
		return err
	}

	l.Info("successfully subscribed to topic")

	for {
		select {
		case <-ctx.Done():
			l.Info("context cancelled, stopping subscriber")
			return ctx.Err()
		case msg, ok := <-messages:
			if !ok {
				l.Info("messages channel closed")
				return nil
			}

			if msg == nil {
				l.Error("received nil message")
				continue
			}
			msg.Ack()

			l.Debug("received message",
				logger.F("uuid", msg.UUID),
				logger.F("payload", string(msg.Payload)))

			var event Event
			if err := json.Unmarshal(msg.Payload, &event); err != nil {
				l.Error("failed to unmarshal event",
					logger.F("uuid", msg.UUID),
					logger.F("payload", string(msg.Payload)),
					logger.Err(err))
				continue
			}

			if event.Platform == "" || event.PlatformType == "" || event.Action == "" || event.UserID == 0 {
				l.Error("invalid event data",
					logger.F("uuid", msg.UUID),
					logger.F("event", event))
				continue
			}

			if err := handler(ctx, event); err != nil {
				l.Error("failed to handle event",
					logger.F("uuid", msg.UUID),
					logger.F("event", event),
					logger.Err(err))
				continue
			}
		}
	}
}
