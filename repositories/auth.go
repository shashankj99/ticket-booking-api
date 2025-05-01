package repositories

import (
	"context"

	"github.com/shashankj99/ticket-booking-api/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) models.AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) RegisterUser(ctx context.Context, registerData *models.AuthCredentials) (*models.User, error) {
	user := &models.User{
		Email:    registerData.Email,
		Password: registerData.Password,
	}

	if err := r.db.Model(&models.User{}).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AuthRepository) GetUser(ctx context.Context, query any, args ...any) (*models.User, error) {
	user := &models.User{}
	if err := r.db.Model(user).Where(query, args...).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
