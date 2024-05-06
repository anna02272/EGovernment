package services

import (
	"auth-service/domain"
)

type UserService interface {
	FindUserById(id string) (*domain.User, error)
	FindUserByEmail(email string) (*domain.User, error)
	FindUserByUsername(username string) (*domain.User, error)
}
