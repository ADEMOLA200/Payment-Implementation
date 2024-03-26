package routes

import (
	"github.com/ADEMOLA200/Payment-Implementation/controller"
	"github.com/ADEMOLA200/Payment-Implementation/middlewares"
	"github.com/ADEMOLA200/Payment-Implementation/webhook"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    // Setup API routes
    api := app.Group("/api", middlewares.CorsMiddleware(), middlewares.LoggingMiddleware)
    api.Post("/makePayment", controller.MakePayment)
	app.Post("/webhook/paystack", webhook.HandlePaystackWebhook) // Connect the webhook handler function to a route


    // Setup webhook endpoint
    // webhookGroup := app.Group("/webhook")
    // webhookGroup.Post("/paystack", webhook.HandlePaystackWebhook)
}