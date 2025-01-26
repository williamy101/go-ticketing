package repository

import (
	"go-ticketing/entity"

	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ticket *entity.Tickets) error
	GetByID(ticketID int) (*entity.Tickets, error)
	Update(ticket *entity.Tickets) error
	GetAllTicketsByUser(userID, page, limit int) ([]entity.Tickets, int64, error)
	FindAvailableTicket(eventID int) (*entity.Tickets, error)
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}

func (r *ticketRepository) Create(ticket *entity.Tickets) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) GetByID(ticketID int) (*entity.Tickets, error) {
	var ticket entity.Tickets
	err := r.db.Preload("Event").First(&ticket, ticketID).Error
	return &ticket, err
}

func (r *ticketRepository) Update(ticket *entity.Tickets) error {
	return r.db.Save(ticket).Error
}

func (r *ticketRepository) GetAllTicketsByUser(userID, page, limit int) ([]entity.Tickets, int64, error) {
	var tickets []entity.Tickets
	var totalItems int64

	r.db.Model(&entity.Tickets{}).Where("user_id = ?", userID).Count(&totalItems)

	offset := (page - 1) * limit
	err := r.db.Preload("Event").Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&tickets).Error

	return tickets, totalItems, err
}

func (r *ticketRepository) FindAvailableTicket(eventID int) (*entity.Tickets, error) {
	var ticket entity.Tickets
	err := r.db.Preload("Event").Where("event_id = ? AND status = ?", eventID, "Available").
		First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}
