package models

import (
	"context"
	"time"
)

type Event struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Date      string    `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EventRepository interface {
	FindMany(ctx context.Context) ([]*Event, error)
	FindOne(ctx context.Context, id uint) (*Event, error)
	Create(ctx context.Context, event *Event) (*Event, error)
	Update(ctx context.Context, id uint, updateData map[string]any) (*Event, error)
	Delete(ctx context.Context, id uint) error
}
