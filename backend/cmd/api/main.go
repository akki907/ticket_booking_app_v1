package main

import (
	"fmt"

	"github.com/akki907/ticket_booking_app_v1/config"
	"github.com/akki907/ticket_booking_app_v1/db"
	"github.com/akki907/ticket_booking_app_v1/handlers"
	middlewares "github.com/akki907/ticket_booking_app_v1/middleware"
	"github.com/akki907/ticket_booking_app_v1/repositories"
	"github.com/akki907/ticket_booking_app_v1/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Ticket-Booking App",
		ServerHeader: "Fiber",
	})
	app.Use(logger.New())
	app.Use(cors.New())
	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	handlers.NewEventRepository(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
