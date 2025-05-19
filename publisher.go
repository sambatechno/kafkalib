package kafkalib

import (
	"context"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type publisher struct {
	wr *kafka.Writer
}

type Publisher interface {
	Publish(ctx context.Context, messages ...proto.Message) error
	Close() error
}

func NewPublisher(brokers []string, dialer *kafka.Dialer) Publisher {
	return &publisher{
		wr: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  brokers,
			Dialer:   dialer,
			Balancer: &kafka.LeastBytes{},
		}),
	}
}

func (p *publisher) Close() error {
	return p.wr.Close()
}

func (p *publisher) Publish(ctx context.Context, messages ...proto.Message) error {
	kfkMsgs := []kafka.Message{}

	for _, m := range messages {
		topic, found := getTopicName(m)
		if !found {
			return NewErrMissingTopic(m)
		}
		mv, _ := proto.Marshal(m)
		km := kafka.Message{
			Topic: topic,
			Value: mv,
		}
		kfkMsgs = append(kfkMsgs, km)
	}

	return p.wr.WriteMessages(ctx, kfkMsgs...)
}
