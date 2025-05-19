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

	consumer, err := kafkalib.NewConsumer(
		cfg.KafkaBrokers,
		"example-consumer",
		dialer,
		kevt.NewUserEvent,
	)
	if err != nil {
		log.Fatalln("failed to create consumer", err)
	}

	for {
		m, err := consumer.ReadMessage(context.Background())
		if err != nil {
			log.Println("err reading", err)
			break
		}
		switch body := m.Body.(type) {
		case *msg.UserEvent_RegistrationSuccess_:
			log.Println("registration success", body, m.CreateTimestamp)
		default:
		}
	}

	if err := consumer.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
