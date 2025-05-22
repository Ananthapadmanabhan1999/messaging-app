package ws

import (
	"messaging-app/websocket-service/kafka"
	"net/http"

	"github.com/gorilla/websocket"
)
 
type Client struct {
	hub  *Hub
	conn *websocket.Conn 
 	send chan []byte
}
 
var upgrader = websocket.Upgrader{
 	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
 
var kafkaProducer *kafka.Producer
 
func SetKafkaProducer(p *kafka.Producer) {
	kafkaProducer = p
}
 
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
 
 		// Publish to Kafka
 		if kafkaProducer != nil {
 			kafkaProducer.PublishMessage(message)
 		}
 	}
 }
 
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
 	if err != nil {
 	return
 	}
 
 	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client
 
 	go client.writePump()
 	go client.readPump()
}
 
func (c *Client) writePump() {
 	defer c.conn.Close()
 	for msg := range c.send {
	err := c.conn.WriteMessage(websocket.TextMessage, msg)
 		if err != nil {
			break
		}	
	}
}
