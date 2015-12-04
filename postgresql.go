package eventsource

import (
	"encoding/json"
	"log"
	"time"
)

import "database/sql"
import "github.com/gogits/gogs/modules/uuid"
import _ "github.com/lib/pq"

type postgresqlEventStore struct {
	DB *sql.DB
}

func (e postgresqlEventStore) Send(stream string, eventData interface{}) (string, error) {
	uuid := uuid.NewV4()
	id := uuid.String()

	eventDataJson, err := json.Marshal(eventData)
	if err != nil {
		return "", err
	}

	result, err := e.DB.Exec("INSERT INTO event_store (id, stream, created_on, body) VALUES ($1, $2, $3, $4)",
		id, stream, time.Now(), eventDataJson)
	if err != nil {
		return "", err
	}

	log.Print("Result is ", result)

	return id, nil
}

func NewPostgresqlEventStore(dataSourceName string) (*postgresqlEventStore, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec("CREATE TABLE IF NOT EXISTS event_store (id TEXT, stream TEXT, created_on TIMESTAMP WITH TIME ZONE, body JSON)")
	if err != nil {
		return nil, err
	}

	log.Print("Result is ", result)

	return &postgresqlEventStore{
		DB: db,
	}, nil
}
