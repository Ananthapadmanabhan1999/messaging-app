package config

import (
	"log"
	"os"
)

type Config struct {
	Port         string
	KafkaBrokers string
	KafkaTopic   string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		log.Fatal("KAFKA_BROKERS not set")
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "messages"
	}

	return Config{
		Port:         port,
		KafkaBrokers: brokers,
		KafkaTopic:   topic,
	}
}
