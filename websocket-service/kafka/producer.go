package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)
 
type Producer struct {
	writer *kafka.Writer
	Topic  string
}
 
func NewProducer(brokers, topic string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
 			Addr:     kafka.TCP(brokers),
 			Topic:    topic,
 			Balancer: &kafka.LeastBytes{},
 		},
 		Topic: topic,
	}
}
 
func (p *Producer) PublishMessage(msg []byte) error {
	err := p.writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(time.Now().Format(time.RFC3339Nano)),
		Value: msg,
 	})
 	if err != nil {
 	log.Printf("Failed to publish message: %v", err)
 	}
 	return err
}
 
func (p *Producer) Close() error {
 	return p.writer.Close()
}
