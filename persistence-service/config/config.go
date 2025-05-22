package config

import (
	"log"
	"os"
)

type Config struct {
    KafkaBrokers string
    KafkaTopic   string
    KafkaGroupID string

    PostgresDSN string
}

func Load() Config {
    brokers := os.Getenv("KAFKA_BROKERS")
    if brokers == "" {
        log.Fatal("KAFKA_BROKERS not set")
    }

    topic := os.Getenv("KAFKA_TOPIC")
    if topic == "" {
        topic = "messages"
    }

    groupID := os.Getenv("KAFKA_GROUP_ID")
    if groupID == "" {
        groupID = "persistence-group"
    }

    pgDSN := os.Getenv("POSTGRES_DSN")
    if pgDSN == "" {
        log.Fatal("POSTGRES_DSN not set")
    }

    return Config{
        KafkaBrokers: brokers,
        KafkaTopic:   topic,
        KafkaGroupID: groupID,
        PostgresDSN:  pgDSN,
    }
}