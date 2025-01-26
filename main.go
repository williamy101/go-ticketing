package main

import (
	"go-ticketing/config"
	"go-ticketing/controller"
	"go-ticketing/repository"
	"go-ticketing/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.New()

	userRepo := repository.NewUserRepository(config.DB)
	eventRepo := repository.NewEventRepository(config.DB)
	ticketRepo := repository.NewTicketRepository(config.DB)
	reportRepo := repository.NewReportRepository(config.DB)

	userService := service.NewUserService(userRepo)
	eventService := service.NewEventService(eventRepo)
	ticketService := service.NewTicketService(ticketRepo, eventRepo)
	reportService := service.NewReportService(reportRepo)

	userController := controller.NewUserController(userService)
	eventController := controller.NewEventController(eventService)
	ticketController := controller.NewTicketController(ticketService)
	reportController := controller.NewReportController(reportService)

	config.Router(r, userController, eventController, ticketController, reportController)
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
