package ws

import (
	"log"
	"messaging-app/websocket-service/config"
	"net/http"
)
 
func StartServer(cfg config.Config, hub *Hub) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
 
	log.Println("WebSocket server listening on port", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatal("Server error:", err)
    }
}