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
	"google.golang.org/protobuf/proto"
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
	evt.UserUuid = "user-uuid"
	evt.Body = &msg.UserEvent_RegistrationSuccess_{
		RegistrationSuccess: &msg.UserEvent_RegistrationSuccess{
			Email: "test@test.com",
		},
	}
	evt.TenantMeta = &msg.TenantMeta{
		TenantId:   "uuid",
		TenantName: "test",
	}

	log.Println("starting to publish")
	startTime := time.Now()
	err = publisher.Publish(context.Background(),
		[]proto.Message{evt},
		kafkalib.WithStaticHeader("X-Api-Key", []byte("test")),
		kafkalib.WithHeader(func(pm proto.Message) []kafka.Header {
			m := pm.(*msg.UserEvent)
			headers := append([]kafka.Header{}, kafka.Header{"X-Tenant-Id", []byte(m.TenantMeta.TenantId)})

			return headers
		}),
	)

	if err != nil {
		panic(err)
	}

	log.Printf("Published in %s\n", time.Since(startTime))
}
