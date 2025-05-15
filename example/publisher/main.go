package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/sambatechno/kafkalib"
	"github.com/sambatechno/kafkalib/example"
	"github.com/sambatechno/kafkalib/gen/kafkalib/msg"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"google.golang.org/protobuf/proto"
)

func main() {
	cfg, err := example.LoadConfigFromFile("config.json")
	if err != nil {
		panic(err)
	}

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		SASLMechanism: plain.Mechanism{
			Username: cfg.SASLUsername,
			Password: cfg.SASLPassword,
		},
		ClientID: cfg.PublisherClientID,
		TLS: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	publisher := kafkalib.NewPublisher(
		cfg.KafkaBrokers,
		dialer,
	)

	err = publisher.Publish(context.Background(), []proto.Message{
		&msg.UserEvent{
			Body: &msg.UserEvent_RegistrationSuccess_{
				RegistrationSuccess: &msg.UserEvent_RegistrationSuccess{
					Email: "test@test.com",
				},
			},
			CreateTimestamp: fmt.Sprint(time.Now().Format(time.RFC3339)),
		},
	})
	if err != nil {
		panic(err)
	}
}
