package middlewares

import "github.com/gofiber/fiber/v2"

// ErrorHandlingMiddleware handles errors and sends appropriate responses
func ErrorHandlingMiddleware(ctx *fiber.Ctx) error {
    // Proceed to next middleware or route handler
    if err := ctx.Next(); err != nil {
        // If there's an error, send an appropriate response
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return nil
}
