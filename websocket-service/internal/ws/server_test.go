package ws

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
)

func TestStartServer(t *testing.T) {
    // Mock Hub
    hub := NewHub()

    // Create a test server
    mux := http.NewServeMux()
    mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(hub, w, r)
    })
    server := httptest.NewServer(mux)
    defer server.Close()

    // Convert the test server URL to a WebSocket URL
    wsURL := "ws" + server.URL[len("http"):]

    // Connect to the WebSocket server
    ws, _, err := websocket.DefaultDialer.Dial(wsURL+"/ws", nil)
    if err != nil {
        t.Fatalf("failed to connect to WebSocket server: %v", err)
    }
    defer ws.Close()

    // If the connection is successful, the test passes
    t.Log("WebSocket connection established successfully")
}