package models

import (
	"context"
	"time"
)

type Ticket struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	EventID   uint      `json:"event_id"`
	Event     Event     `json:"event" gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Entered   bool      `json:"entered" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TicketRepository interface {
	FindMany(ctx context.Context) ([]*Ticket, error)
	FindOne(ctx context.Context, id uint) (*Ticket, error)
	Create(ctx context.Context, ticket *Ticket) (*Ticket, error)
	Update(ctx context.Context, id uint, updateData map[string]any) (*Ticket, error)
	Delete(ctx context.Context, id uint) error
}
