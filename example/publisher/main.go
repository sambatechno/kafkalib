package main

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/sambatechno/kafkalib"
	"github.com/sambatechno/kafkalib/example"
	"github.com/sambatechno/kafkalib/gen/kafkalib/msg"
	"github.com/sambatechno/kafkalib/kevt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func main() {
	log.Println("init")
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

	evt := kevt.NewUserEvent()
	evt.Body = &msg.UserEvent_RegistrationSuccess_{
		RegistrationSuccess: &msg.UserEvent_RegistrationSuccess{
			Email: "test@test.com",
		},
	}

	log.Println("starting to publish")
	startTime := time.Now()
	err = publisher.Publish(context.Background(), evt)

	if err != nil {
		panic(err)
	}

	log.Printf("Published in %s\n", time.Since(startTime))
}
