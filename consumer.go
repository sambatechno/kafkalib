package kafkalib

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type Consumer[T proto.Message] struct {
	r        *kafka.Reader
	newTfunc func() T
}

func NewConsumer[T proto.Message](brokers []string, groupId string, dialer *kafka.Dialer, newTfunc func() T) (*Consumer[T], error) {
	t := newTfunc()
	topic, found := getTopicName(t)
	if !found {
		return nil, NewErrMissingTopic(t)
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: groupId,
		Topic:   topic,
		Dialer:  dialer,
	})

	return &Consumer[T]{
		r:        r,
		newTfunc: newTfunc,
	}, nil
}

func (c *Consumer[T]) Close() error {
	return c.r.Close()
}

func (c *Consumer[T]) ReadMessage(ctx context.Context) (T, error) {
	t := c.newTfunc()

	m, err := c.r.ReadMessage(ctx)
	if err != nil {
		return t, err
	}

	if err := proto.Unmarshal(m.Value, t); err != nil {
		return t, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return t, nil
}
