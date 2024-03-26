package routes

import (
	"github.com/ADEMOLA200/Payment-Implementation/controller"
	"github.com/ADEMOLA200/Payment-Implementation/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Setup routes
    api := app.Group("/api", middlewares.CorsMiddleware(), middlewares.LoggingMiddleware)

    api.Post("/makePayment", controller.MakePayment)
}
