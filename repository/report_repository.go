package repository

import (
	"go-ticketing/entity"

	"gorm.io/gorm"
)

type ReportRepository interface {
	GetSummaryReport() (entity.ReportSummary, error)
	GetEventReport(eventID int) ([]entity.EventReport, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) GetSummaryReport() (entity.ReportSummary, error) {
	var summary entity.ReportSummary

	query := `
		SELECT 
			COUNT(t.ticket_id) AS totalTicketsSold, 
			COALESCE(SUM(e.price), 0) AS totalRevenue
		FROM tickets t
		INNER JOIN events e ON t.event_id = e.event_id
		WHERE t.status = 'Booked'
	`

	err := r.db.Raw(query).Scan(&summary).Error

	return summary, err
}

func (r *reportRepository) GetEventReport(eventID int) ([]entity.EventReport, error) {
	var reports []entity.EventReport

	query := `
		SELECT 
			e.event_name AS eventName, 
			COUNT(t.ticket_id) AS ticketsSold, 
			COALESCE(SUM(e.price), 0) AS revenue
		FROM tickets t
		INNER JOIN events e ON t.event_id = e.event_id
		WHERE t.event_id = ? AND t.status = 'Booked'
		GROUP BY e.event_name
	`

	err := r.db.Raw(query, eventID).Scan(&reports).Error
	return reports, err
}
