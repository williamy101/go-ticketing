package service

import (
	"errors"
	"go-ticketing/entity"
	"go-ticketing/repository"
	"time"
)

type EventService interface {
	GetAll() ([]entity.Events, error)
	GetByID(eventID int) (*entity.Events, error)
	Create(event *entity.Events) error
	Update(event *entity.Events) error
	Delete(eventID int) error
}

type eventService struct {
	eventRepo repository.EventRepository
}

func NewEventService(eventRepo repository.EventRepository) EventService {
	return &eventService{eventRepo: eventRepo}
}

func (s *eventService) GetAll() ([]entity.Events, error) {
	return s.eventRepo.GetAll()
}

func (s *eventService) GetByID(eventID int) (*entity.Events, error) {
	return s.eventRepo.GetByID(eventID)
}

func (s *eventService) Create(event *entity.Events) error {
	if event.Capacity < 0 {
		return errors.New("capacity cannot be negative")
	}

	if event.Price < 0 {
		return errors.New("price cannot be negative")
	}
	parsedDate, err := time.Parse("2006-01-02", event.Date)
	if err != nil {
		return errors.New("invalid date format, expected YYYY-MM-DD")
	}
	event.ParsedDate = parsedDate
	existingEvents, err := s.eventRepo.GetAll()
	if err != nil {
		return err
	}

	for _, e := range existingEvents {
		if e.EventName == event.EventName {
			return errors.New("event name must be unique")
		}
	}

	if event.Status != "" && event.Status != "Active" {
		return errors.New("event status must be 'Active' upon creation")
	}
	event.Status = "Active"

	err = s.eventRepo.Create(event)
	if err != nil {
		return err
	}

	for i := 0; i < event.Capacity; i++ {
		ticket := entity.Tickets{
			EventID: event.EventID,
			UserID:  nil,
			Status:  "Available",
		}

		err := s.eventRepo.CreateTicket(&ticket)
		if err != nil {
			return errors.New("failed to generate tickets for the event")
		}
	}

	return nil
}

func (s *eventService) Update(event *entity.Events) error {
	// Fetch the existing event
	existingEvent, err := s.eventRepo.GetByID(event.EventID)
	if err != nil {
		return errors.New("event not found")
	}

	if existingEvent.Status == "Ongoing" || existingEvent.Status == "Completed" {
		return errors.New("cannot update an event that is ongoing or completed")
	}

	existingEventDate, err := time.Parse(time.RFC3339, existingEvent.Date)
	if err != nil {
		return errors.New("invalid date format in existing event")
	}

	if existingEventDate.Before(time.Now()) {
		return errors.New("cannot update an event that has already occurred")
	}

	newDate, err := time.Parse("2006-01-02", event.Date)
	if err != nil {
		return errors.New("invalid date format, expected YYYY-MM-DD")
	}
	event.Date = newDate.Format("2006-01-02")

	bookedTicketsCount, err := s.eventRepo.CountBookedTickets(event.EventID)
	if err != nil {
		return errors.New("failed to fetch booked tickets count")
	}

	if event.Capacity < bookedTicketsCount {
		return errors.New("capacity cannot be less than the number of booked tickets")
	}

	if event.Capacity > existingEvent.Capacity {
		for i := 0; i < event.Capacity-existingEvent.Capacity; i++ {
			ticket := entity.Tickets{
				EventID: event.EventID,
				Status:  "Available",
			}
			if err := s.eventRepo.CreateTicket(&ticket); err != nil {
				return errors.New("failed to generate additional tickets")
			}
		}
	} else if event.Capacity < existingEvent.Capacity {
		excessTickets := existingEvent.Capacity - event.Capacity
		if err := s.eventRepo.DeleteExcessTickets(event.EventID, excessTickets); err != nil {
			return errors.New("failed to delete excess tickets")
		}
	}

	event.CreatedAt = existingEvent.CreatedAt

	return s.eventRepo.Update(event)
}

func (s *eventService) Delete(eventID int) error {
	bookedTicketsCount, err := s.eventRepo.CountBookedTickets(eventID)
	if err != nil {
		return errors.New("failed to check booked tickets")
	}

	if bookedTicketsCount > 0 {
		return errors.New("cannot delete an event with booked tickets")
	}

	if err := s.eventRepo.DeleteExcessTickets(eventID, 0); err != nil {
		return errors.New("failed to delete associated tickets")
	}

	return s.eventRepo.Delete(eventID)
}
