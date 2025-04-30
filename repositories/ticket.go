package repositories

import (
	"context"

	"github.com/shashankj99/ticket-booking-api/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) FindMany(ctx context.Context) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}
	if err := r.db.Model(&models.Ticket{}).Preload("Event").Order("created_at DESC").Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketRepository) FindOne(ctx context.Context, id uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}
	if err := r.db.Model(ticket).Where("id = ?", id).Preload("Event").First(ticket).Error; err != nil {
		return nil, err
	}
	return ticket, nil
}

func (r *TicketRepository) Create(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	if err := r.db.Model(ticket).Create(ticket).Error; err != nil {
		return nil, err
	}
	return r.FindOne(ctx, ticket.ID)
}

func (r *TicketRepository) Update(ctx context.Context, id uint, updateData map[string]any) (*models.Ticket, error) {
	ticket := &models.Ticket{}
	if err := r.db.Model(ticket).Where("id = ?", id).Updates(updateData).Error; err != nil {
		return nil, err
	}
	return r.FindOne(ctx, id)
}

func (r *TicketRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.Model(&models.Ticket{}).Where("id = ?", id).Delete(&models.Ticket{}).Error; err != nil {
		return err
	}
	return nil
}
