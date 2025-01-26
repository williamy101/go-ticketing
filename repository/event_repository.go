package repository

import (
	"go-ticketing/entity"

	"gorm.io/gorm"
)

type EventRepository interface {
	GetAll() ([]entity.Events, error)
	GetByID(eventID int) (*entity.Events, error)
	Create(event *entity.Events) error
	Update(event *entity.Events) error
	Delete(eventID int) error
	CreateTicket(ticket *entity.Tickets) error
	CountBookedTickets(eventID int) (int, error)
	DeleteExcessTickets(eventID int, excessTicketsCount int) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) GetAll() ([]entity.Events, error) {
	var events []entity.Events
	err := r.db.Find(&events).Error
	return events, err
}

func (r *eventRepository) GetByID(eventID int) (*entity.Events, error) {
	var event entity.Events
	err := r.db.First(&event, eventID).Error
	return &event, err
}

func (r *eventRepository) Create(event *entity.Events) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) Update(event *entity.Events) error {
	return r.db.Save(event).Error
}

func (r *eventRepository) Delete(eventID int) error {
	return r.db.Delete(&entity.Events{}, eventID).Error
}

func (r *eventRepository) CreateTicket(ticket *entity.Tickets) error {
	return r.db.Create(ticket).Error
}

func (r *eventRepository) CountBookedTickets(eventID int) (int, error) {
	var count int64
	err := r.db.Model(&entity.Tickets{}).Where("event_id = ? AND status = ?", eventID, "Booked").Count(&count).Error
	return int(count), err
}

func (r *eventRepository) DeleteExcessTickets(eventID int, excessTicketsCount int) error {
	if excessTicketsCount == 0 {
		return r.db.Where("event_id = ?", eventID).Delete(&entity.Tickets{}).Error
	}

	subQuery := r.db.Model(&entity.Tickets{}).Select("ticket_id").Where("event_id = ? AND status = ?", eventID, "Available").Limit(excessTicketsCount)
	return r.db.Where("ticket_id IN (?)", subQuery).Delete(&entity.Tickets{}).Error
}
