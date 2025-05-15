package example

import (
	"encoding/json"
	"os"
)

type Config struct {
	KafkaBrokers      []string `json:"kafka_brokers"`
	PublisherClientID string   `json:"publisher_client_id"`
	ConsumerClientID  string   `json:"consumer_client_id"`
	SASLUsername      string   `json:"sasl_username"`
	SASLPassword      string   `json:"sasl_password"`
}

func LoadConfigFromFile(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
