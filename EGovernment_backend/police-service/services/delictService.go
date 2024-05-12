package services

import (
	"context"
	"police-service/domain"
)

type DelictService interface {
	InsertDelict(delict *domain.DelictCreate, policemanID string) (*domain.Delict, string, error)
	GetAllDelicts() ([]*domain.Delict, error)
	GetDelictById(delictId string, ctx context.Context) (*domain.Delict, error)
	GetAllDelictsByDelictType(delictType domain.DelictType) ([]*domain.Delict, error)
}
