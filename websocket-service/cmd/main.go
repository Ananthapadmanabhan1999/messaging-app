package main

import (
	"context"
	"messaging-app/websocket-service/config"
	"messaging-app/websocket-service/internal/ws"
	"messaging-app/websocket-service/kafka"
)
   

func main() {
	cfg := config.Load()
	hub := ws.NewHub()
	go hub.Run()

	// Kafka Producer
	producer := kafka.NewProducer(cfg.KafkaBrokers, cfg.KafkaTopic)
	ws.SetKafkaProducer(producer)

	// Kafka Consumer
	consumer := kafka.NewConsumer(cfg.KafkaBrokers, cfg.KafkaTopic, "websocket-group")
	go consumer.Consume(context.Background(), func(msg []byte) {
		hub.Broadcast(msg)
	})

	ws.StartServer(cfg, hub)
}
