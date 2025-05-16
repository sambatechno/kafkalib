package kafkalib

import (
	"context"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type Publisher struct {
	wr *kafka.Writer
}

func NewPublisher(brokers []string, dialer *kafka.Dialer) *Publisher {
	return &Publisher{
		wr: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  brokers,
			Dialer:   dialer,
			Balancer: &kafka.LeastBytes{},
		}),
	}
}
func (p *Publisher) Publish(ctx context.Context, messages []proto.Message) error {
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
