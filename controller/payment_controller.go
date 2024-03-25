package controller

import (
	"encoding/json"

	"github.com/ADEMOLA200/Payment-Implementation/database"
	"github.com/ADEMOLA200/Payment-Implementation/model"
	_ "github.com/ADEMOLA200/Payment-Implementation/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid" // Import the UUID package
)

func MakePayment(c *fiber.Ctx) error {
    // Parse the request body into PaymentRequest struct
    var request model.PaymentRequest
    if err := c.BodyParser(&request); err != nil {
        return err
    }

    // Generate a new UUID as the reference
    reference := uuid.New().String()

    // Convert metadata map to JSON string
    metadataJSON, err := json.Marshal(make(map[string]interface{})) // Initialize an empty map and marshal it to JSON
    if err != nil {
        return err
    }

    // Create a new payment record
    payment := model.Payment{
        Amount:    request.Amount,
        Email:     request.Email,
        Metadata:  string(metadataJSON), // Assign the JSON string to Metadata
        Reference: reference, // Use the generated UUID as the reference
    }

    // Save the payment record to the database
    if err := database.DB.Create(&payment).Error; err != nil {
        return err
    }

    // Return success response
    return c.SendString("Subscription payment made successfully")
}