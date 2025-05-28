package kafkalib

import (
	"context"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type publisher struct {
	wr *kafka.Writer
}

type Publisher interface {
	Publish(ctx context.Context, messages []proto.Message, opt ...MessageOpt) error
	Close() error
}

func NewPublisher(brokers []string, dialer *kafka.Dialer) Publisher {
	p := &publisher{
		wr: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  brokers,
			Dialer:   dialer,
			Balancer: &kafka.LeastBytes{},
		}),
	}

	return p
}

func (p *publisher) Close() error {
	return p.wr.Close()
}

func (p *publisher) Publish(ctx context.Context, messages []proto.Message, opt ...MessageOpt) error {
	kfkMsgs := []kafka.Message{}

	for _, pm := range messages {
		topic, found := getTopicName(pm)
		if !found {
			return NewErrMissingTopic(pm)
		}
		mv, _ := protojson.Marshal(pm)
		km := &kafka.Message{
			Topic: topic,
			Value: mv,
		}

		for _, o := range opt {
			km = o(km, pm)
		}

		kfkMsgs = append(kfkMsgs, *km)
	}

	return p.wr.WriteMessages(ctx, kfkMsgs...)
}

type MessageOpt func(*kafka.Message, proto.Message) *kafka.Message

func WithStaticHeader(key string, value []byte) MessageOpt {
	return func(m *kafka.Message, pm proto.Message) *kafka.Message {
		m.Headers = append(m.Headers, kafka.Header{Key: key, Value: value})
		return m
	}
}

func WithHeader(fn func(pm proto.Message) []kafka.Header) MessageOpt {
	return func(m *kafka.Message, pm proto.Message) *kafka.Message {
		m.Headers = append(m.Headers, fn(pm)...)
		return m
	}
}
