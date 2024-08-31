package repositories

import (
	"context"
	"time"

	"github.com/akki907/ticket_booking_app_v1/models"
)

type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {

	events := []*models.Event{}

	events = append(events, &models.Event{
		ID:        "1",
		Name:      "Event 1",
		Date:      time.Now(),
		Location:  "Location 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId uint) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) DeleteOne(ctx context.Context, eventId uint) error {
	return nil
}

func NewEventRepository(db *any) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
