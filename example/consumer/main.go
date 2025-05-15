package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

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

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.KafkaBrokers,
		Dialer:  dialer,
		Topic:   "user-event",
		GroupID: "consumer-a",
	})

	log.Println("start listening")

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("err reading", err)
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(m.Key))

		fmt.Printf("%#v\n", string(m.Value))
		ev := &msg.UserEvent{}
		if err := proto.Unmarshal(m.Value, ev); err != nil {
			log.Println("failed to unmarshal ", err)
		}
		switch body := ev.Body.(type) {
		case *msg.UserEvent_RegistrationSuccess_:
			fmt.Println("registration success", body, ev.CreateTimestamp)
		default:
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
