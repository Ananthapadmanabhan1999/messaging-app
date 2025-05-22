package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)
 
type Consumer struct {
	reader *kafka.Reader
}
 
func NewConsumer(brokers, topic, groupID string) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
 			Brokers:  []string{brokers},
 			GroupID:  groupID,
 			Topic:    topic,
			MinBytes: 10e3, // 10KB
 			MaxBytes: 10e6, // 10MB
		}),
 	}
 }
 
func (c *Consumer) Consume(ctx context.Context, handle func([]byte)) {
 	for {
 		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
 		log.Println("Error reading message:", err)
 		continue
 		}
		handle(msg.Value)
	}
}
 
func (c *Consumer) Close() error {
	return c.reader.Close()
}
