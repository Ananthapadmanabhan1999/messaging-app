package persistence

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/segmentio/kafka-go"
)

type Message struct {
    ID      string `json:"id"`
    Content string `json:"content"`
}

type PersistenceService struct {
    kafkaReader *kafka.Reader
    db          *pgxpool.Pool
}

func NewPersistenceService(brokers, topic, groupID, pgDSN string) (*PersistenceService, error) {
    dbpool, err := pgxpool.New(context.Background(), pgDSN)
    if err != nil {
        return nil, err
    }

    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers:  []string{brokers},
        GroupID:  groupID,
        Topic:    topic,
        MinBytes: 10e3,
        MaxBytes: 10e6,
    })

    return &PersistenceService{
        kafkaReader: reader,
        db:          dbpool,
    }, nil
}

func (p *PersistenceService) Run(ctx context.Context) {
    for {
        msg, err := p.kafkaReader.ReadMessage(ctx)
        if err != nil {
            log.Printf("Error reading message: %v", err)
            continue
        }

        err = p.saveMessage(ctx, msg.Value)
        if err != nil {
            log.Printf("Failed to save message: %v", err)
            // Decide how to handle error: continue or break
        } else {
            log.Printf("Message saved successfully, offset: %d", msg.Offset)
        }
    }
}

func (p *PersistenceService) saveMessage(ctx context.Context, msg []byte) error {
    var m Message
    err := json.Unmarshal(msg, &m)
    if err != nil {
        return err
    }

    sql := `INSERT INTO messages (id, content) VALUES ($1, $2)`
    _, err = p.db.Exec(ctx, sql, m.ID, m.Content)
    return err
}

func (p *PersistenceService) Close() {
    p.kafkaReader.Close()
    p.db.Close()
}