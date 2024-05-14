package services

import "statistics-service/domain"

type ReportDelicTypeService interface {
	Create(report *domain.ReportDelict) (error, bool)
}
