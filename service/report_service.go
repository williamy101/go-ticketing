package service

import (
	"go-ticketing/entity"
	"go-ticketing/repository"
)

type ReportService interface {
	GetSummaryReport() (entity.ReportSummary, error)
	GetEventReport(eventID int) ([]entity.EventReport, error)
}

type reportService struct {
	reportRepo repository.ReportRepository
}

func NewReportService(reportRepo repository.ReportRepository) ReportService {
	return &reportService{reportRepo: reportRepo}
}

func (s *reportService) GetSummaryReport() (entity.ReportSummary, error) {
	return s.reportRepo.GetSummaryReport()
}

func (s *reportService) GetEventReport(eventID int) ([]entity.EventReport, error) {
	return s.reportRepo.GetEventReport(eventID)
}
