package services

import (
	"context"
	"police-service/domain"
)

type ReportService interface {
	InsertReport(report *domain.ReportCreate, policemanID string, delictID string, carAccidentID string) (*domain.Report, string, error)
	GetAllReport() ([]*domain.Report, error)
	GetReportById(reportId string, ctx context.Context) (*domain.Report, error)
	GetAllReportsByDelictType(delictType domain.DelictType) ([]*domain.Report, error)
}
