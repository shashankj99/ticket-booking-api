package models

import (
	"context"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type AuthCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type AuthRepository interface {
	RegisterUser(ctx context.Context, registerData *AuthCredentials) (*User, error)
	GetUser(ctx context.Context, query any, args ...any) (*User, error)
}

type AuthService interface {
	Register(ctx context.Context, registerData *AuthCredentials) (string, *User, error)
	Login(ctx context.Context, loginData *AuthCredentials) (string, *User, error)
}

func MatchPassword(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
