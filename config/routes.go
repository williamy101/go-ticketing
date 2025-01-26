package config

import (
	"go-ticketing/controller"
	"go-ticketing/middleware"

	"github.com/gin-gonic/gin"
)

func Router(
	router *gin.Engine,
	userController *controller.UserController,
	eventController *controller.EventController,
	ticketController *controller.TicketController,
	reportController *controller.ReportController,
) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
	}

	adminUserRoutes := router.Group("/admin/users")
	adminUserRoutes.Use(middleware.AuthMiddleware("Admin"))
	{
		adminUserRoutes.GET("/", userController.GetAllUsers)
		adminUserRoutes.GET("/:id", userController.GetUserByID)
		adminUserRoutes.PUT("/:id", userController.UpdateUserRole)
	}

	eventRoutes := router.Group("/events")
	{
		eventRoutes.GET("/", middleware.AuthMiddleware("User", "Admin"), eventController.GetAllEvents)
		eventRoutes.GET("/:id", middleware.AuthMiddleware("User", "Admin"), eventController.GetEventByID)
		eventRoutes.POST("/", middleware.AuthMiddleware("Admin"), eventController.CreateEvent)
		eventRoutes.PUT("/:id", middleware.AuthMiddleware("Admin"), eventController.UpdateEvent)
		eventRoutes.DELETE("/:id", middleware.AuthMiddleware("Admin"), eventController.DeleteEvent)
	}

	ticketRoutes := router.Group("/tickets")
	{
		ticketRoutes.GET("/", middleware.AuthMiddleware("User"), ticketController.GetAllTicketsByUser)
		ticketRoutes.GET("/:id", middleware.AuthMiddleware("User"), ticketController.GetTicketByID)
		ticketRoutes.POST("/", middleware.AuthMiddleware("User"), ticketController.PurchaseTicket)
		ticketRoutes.PATCH("/:id", middleware.AuthMiddleware("User", "Admin"), ticketController.UpdateTicketStatus)
	}

	reportRoutes := router.Group("/reports")
	reportRoutes.Use(middleware.AuthMiddleware("Admin"))
	{
		reportRoutes.GET("/summary", reportController.GetSummaryReport)
		reportRoutes.GET("/event/:id", reportController.GetEventReport)
	}
}
