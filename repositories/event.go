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
	if err := r.db.Model(&models.Event{}).Order("created_at DESC").Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) FindOne(ctx context.Context, id uint) (*models.Event, error) {
	event := &models.Event{}
	if err := r.db.Model(event).Where("id = ?", id).First(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventRepository) Create(ctx context.Context, event *models.Event) (*models.Event, error) {
	if err := r.db.Model(event).Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventRepository) Update(ctx context.Context, id uint, updateData map[string]any) (*models.Event, error) {
	event := &models.Event{}
	if err := r.db.Model(event).Where("id = ?", id).Updates(updateData).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.Model(&models.Event{}).Where("id = ?", id).Delete(&models.Event{}).Error; err != nil {
		return err
	}
	return nil
}
