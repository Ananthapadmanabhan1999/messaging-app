package main

import (
	"fmt"
	"log"
	"messaging-app/auth-service/config"
	"messaging-app/auth-service/internal/auth"
	"net/http"
)

func main() {
    cfg := config.Load()
    service := auth.NewAuthService()

    mux := http.NewServeMux()
    mux.HandleFunc("/login", auth.LoginHandler(service))
    mux.HandleFunc("/secure", auth.AuthMiddleware(service, secureEndpoint))

    addr := fmt.Sprintf(":%s", cfg.Port)
    log.Println("Starting auth service on", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}

func secureEndpoint(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This is a protected route."))
}