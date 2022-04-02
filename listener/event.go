package listener

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

//go:generate  easyjson nats_publisher.go

// NatsPublisher represent event publisher.
type NatsPublisher struct {
	conn stan.Conn
}

// Close NATS connection.
func (n NatsPublisher) Close() error {
	return n.conn.Close()
}

// Event event structure for publishing to the NATS server.
//easyjson:json
type Event struct {
	ID        uuid.UUID              `json:"id"`
	Schema    string                 `json:"schema"`
	Table     string                 `json:"table"`
	Action    string                 `json:"action"`
	Data      map[string]interface{} `json:"data"`
	EventTime time.Time              `json:"commitTime"`
	Topic     string                 `json:"-"`
}

// GetSubjectName creates subject name from the prefix, schema and table name.
func (e Event) GetSubjectName(prefix string) string {
	return fmt.Sprintf("%s%s_%s", prefix, e.Schema, e.Table)
}
