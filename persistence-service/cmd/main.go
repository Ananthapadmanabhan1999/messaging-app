package main

import (
	"context"
	"log"
	"messaging-app/persistence-service/config"
	"messaging-app/persistence-service/internal/persistence"
	"os"
	"os/signal"

	"syscall"
)

func main() {
    cfg := config.Load()

    service, err := persistence.NewPersistenceService(cfg.KafkaBrokers, cfg.KafkaTopic, cfg.KafkaGroupID, cfg.PostgresDSN)
    if err != nil {
        log.Fatalf("Failed to create persistence service: %v", err)
    }

    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    go service.Run(ctx)

    log.Println("Persistence service started")

    <-ctx.Done()
    log.Println("Shutting down...")
    service.Close()
}