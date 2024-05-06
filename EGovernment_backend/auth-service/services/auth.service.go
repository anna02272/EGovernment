package services

import (
	"auth-service/domain"
	"context"
)

type AuthService interface {
	Login() (*domain.User, error)
	Registration(user *domain.User, ctx context.Context) (*domain.UserResponse, error)
}
