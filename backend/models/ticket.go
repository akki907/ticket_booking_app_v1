package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	EventID   uuid.UUID `json:"eventId"`
	UserID    uuid.UUID `json:"userId" gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Event     Event     `json:"event" gorm:"foreignkey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Entered   bool      `json:"entered" default:"false"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TicketRepository interface {
	GetMany(ctx context.Context) ([]*Ticket, error)
	GetOne(ctx context.Context, ticketId any) (*Ticket, error)
	CreateOne(ctx context.Context, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context, ticketId any, updateData map[string]interface{}) (*Ticket, error)
}

type ValidateTicket struct {
	TicketId uuid.UUID `json:"ticketId"`
	OwnerId  uuid.UUID `json:"ownerId"`
}

func (ticket *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	ticket.ID = uuid.New()
	return
}
