package main

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID        string  `json:"id"`
	CreatedAt int64   `json:"created_at"`
	Raw       string  `json:"raw"`
	Meta      string  `json:"meta"`
	Events    []Event `json:"-"`
}

func (t *Transaction) BeforeCreate() (err error) {
	id, _ := uuid.NewV4()
	t.ID = id.String()
	t.CreatedAt = time.Now().UnixNano()
	return
}

type Event struct {
	ID            uint   `json:"id"`
	CreatedAt     int64  `json:"created_at"`
	Name          string `json:"name"`
	TransactionID string `json:"transaction_id"`
}

func (e *Event) BeforeCreate() (err error) {
	if e.Name == "" {
		err = errors.New("Name field cannot be empty")
	}

	e.CreatedAt = time.Now().UnixNano()
	return
}
