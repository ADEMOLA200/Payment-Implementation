package webhook

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type paystackEvent struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func HandlePaystackWebhook(c *fiber.Ctx) error {
	// Verify webhook signature (implementation depends on Paystack's documentation)

	// Parse webhook event
	event := &paystackEvent{}
	if err := c.BodyParser(event); err != nil {
		zap.L().Error("Failed to parse webhook event", zap.Error(err))
		return c.SendStatus(http.StatusInternalServerError)
	}

	// Process webhook event
	switch event.Event {
	case "charge.success":
		// Payment was successful
		zap.L().Info("Payment successful", zap.Any("data", event.Data))
		// Update your database, send email confirmation, etc.
		return c.SendStatus(http.StatusCreated)
	case "charge.failed":
		// Payment failed
		zap.L().Info("Payment failed", zap.Any("data", event.Data))
		// Handle failed payment scenario
		return c.SendStatus(http.StatusBadRequest)
	case "charge.canceled":
		// Payment was canceled by the user
		zap.L().Info("Payment canceled", zap.Any("data", event.Data))
		// Update your database, send email notification, etc.
		return c.SendStatus(http.StatusGone)
	default:
		// Unknown event type
		zap.L().Warn("Unknown event type", zap.String("event", event.Event))
		return c.SendStatus(http.StatusNotImplemented)
	}
}

func New(logger *zap.Logger) *fiber.App {
	app := fiber.New()

	// Set up route
	app.Post("/webhook", func(c *fiber.Ctx) error {
		// Log the request body
		logger.Info("Received webhook event", zap.Any("requestBody", string(c.Request().Body())))

		// Process webhook event
		err := HandlePaystackWebhook(c)
		if err != nil {
			return err
		}

		return nil
	})

	return app
}