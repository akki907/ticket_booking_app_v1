package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID                    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name                  string    `json:"name"`
	Location              string    `json:"location"`
	TotalTicketsPurchased int64     `json:"totalTicketsPurchased" gorm:"-"`
	TotalTicketsEntered   int64     `json:"totalTicketsEntered" gorm:"-"`
	Date                  time.Time `json:"date"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId uuid.UUID) (*Event, error)
	CreateOne(ctx context.Context, event *Event) (*Event, error)
	UpdateOne(ctx context.Context, eventId uuid.UUID, updateData map[string]interface{}) (*Event, error)
	DeleteOne(ctx context.Context, eventId uuid.UUID) error
}

func (event *Event) BeforeCreate(tx *gorm.DB) (err error) {
	event.ID = uuid.New()
	return
}
