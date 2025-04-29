package repositories

import (
	"context"

	"github.com/shashankj99/ticket-booking-api/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) FindMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}
	if err := r.db.Model(&models.Event{}).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) FindOne(ctx context.Context, id string) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) Create(ctx context.Context, event *models.Event) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) Update(ctx context.Context, event *models.Event) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) Delete(ctx context.Context, id string) error {
	return nil
}
