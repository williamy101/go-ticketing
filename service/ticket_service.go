package service

import (
	"errors"
	"go-ticketing/entity"
	"go-ticketing/repository"
)

type TicketService interface {
	GetAllTicketsByUser(userID, page, limit int) (map[string]interface{}, error)
	GetTicketByID(ticketID int) (*entity.Tickets, error)
	PurchaseTicket(eventID, userID int) (*entity.Tickets, error)
	UpdateTicketStatus(ticketID int, status string) error
	GetAvailableTicket(eventID int) (*entity.Tickets, error)
}

type ticketService struct {
	ticketRepo repository.TicketRepository
	eventRepo  repository.EventRepository
}

func NewTicketService(ticketRepo repository.TicketRepository, eventRepo repository.EventRepository) TicketService {
	return &ticketService{ticketRepo: ticketRepo, eventRepo: eventRepo}
}

func (s *ticketService) GetAllTicketsByUser(userID, page, limit int) (map[string]interface{}, error) {
	tickets, totalItems, err := s.ticketRepo.GetAllTicketsByUser(userID, page, limit)
	if err != nil {
		return nil, err
	}

	totalPages := (totalItems + int64(limit) - 1) / int64(limit)

	return map[string]interface{}{
		"current_page": page,
		"total_pages":  totalPages,
		"total_items":  totalItems,
		"data":         tickets,
	}, nil
}

func (s *ticketService) GetTicketByID(ticketID int) (*entity.Tickets, error) {
	return s.ticketRepo.GetByID(ticketID)
}

func (s *ticketService) PurchaseTicket(eventID, userID int) (*entity.Tickets, error) {
	event, err := s.eventRepo.GetByID(eventID)
	if err != nil || event == nil {
		return nil, errors.New("event not found")
	}

	if event.Capacity <= 0 {
		return nil, errors.New("event is sold out")
	}

	if event.Status == "Completed" {
		return nil, errors.New("cannot purchase tickets for a completed event")
	}

	ticket, err := s.ticketRepo.FindAvailableTicket(eventID)
	if err != nil || ticket == nil {
		return nil, errors.New("no available tickets")
	}

	ticket.Status = "Booked"
	ticket.UserID = &userID
	err = s.ticketRepo.Update(ticket)
	if err != nil {
		return nil, err
	}

	event.Capacity--
	err = s.eventRepo.Update(event)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *ticketService) UpdateTicketStatus(ticketID int, status string) error {
	ticket, err := s.ticketRepo.GetByID(ticketID)
	if err != nil || ticket == nil {
		return errors.New("ticket not found")
	}

	event, err := s.eventRepo.GetByID(ticket.EventID)
	if err != nil || event == nil {
		return errors.New("event not found")
	}

	if event.Status == "Ongoing" {
		return errors.New("cannot update ticket for an ongoing event")
	}

	if status != "Cancelled" && status != "Booked" && status != "Available" {
		return errors.New("invalid ticket status")
	}

	switch status {
	case "Cancelled":
		if ticket.Status != "Booked" {
			return errors.New("only booked tickets can be cancelled")
		}

		ticket.Status = "Cancelled"
		ticket.UserID = nil
		err = s.ticketRepo.Update(ticket)
		if err != nil {
			return err
		}

		event.Capacity++
		err = s.eventRepo.Update(event)
		if err != nil {
			return err
		}

		newTicket := entity.Tickets{
			EventID: ticket.EventID,
			Status:  "Available",
		}
		err = s.ticketRepo.Create(&newTicket)
		if err != nil {
			return err
		}

	case "Available":
		return errors.New("cannot transition ticket back to available from current state")

	default:
		return errors.New("invalid status transition")
	}

	return nil
}

func (s *ticketService) GetAvailableTicket(eventID int) (*entity.Tickets, error) {
	ticket, err := s.ticketRepo.FindAvailableTicket(eventID)
	if err != nil {
		return nil, errors.New("no available tickets found")
	}
	return ticket, nil
}
