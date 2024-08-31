package models

import (
	"context"
	"time"
)

type Event struct {
	ID        string
	Name      string
	Date      time.Time
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId uint) (*Event, error)
	CreateOne(ctx context.Context, event *Event) (*Event, error)
	UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*Event, error)
	DeleteOne(ctx context.Context, eventId uint) error
}
