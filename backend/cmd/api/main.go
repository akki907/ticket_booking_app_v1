package main

import (
	"github.com/akki907/ticket_booking_app_v1/handlers"
	"github.com/akki907/ticket_booking_app_v1/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		AppName:      "Ticket-Booking App",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(nil)

	// Routing
	server := app.Group("/api")

	handlers.NewEventRepository(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
