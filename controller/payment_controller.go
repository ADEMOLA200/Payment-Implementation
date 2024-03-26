package controller

import (
	"errors"
	"fmt"
	"log"

	"github.com/ADEMOLA200/Payment-Implementation/database"
	"github.com/ADEMOLA200/Payment-Implementation/model"
	"github.com/ADEMOLA200/Payment-Implementation/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var paystackService *services.PaystackService

const paystackTestKey = ""

// Initialize the Paystack service with the provided secret key
func setupPaystackService(secretKey string) error {
    var err error
    // Initialize Paystack service with the secret key
    paystackService, err = InitializePaystackService(secretKey)
    if err != nil {
        return fmt.Errorf("failed to initialize Paystack service: %w", err)
    }
    return nil
}

// init function to set up Paystack service when the package is imported
func init() {
    if err := setupPaystackService(paystackTestKey); err != nil {
        log.Fatalf("Failed to set up Paystack service: %v", err)
    }
}

// MakePayment handles payment requests
func MakePayment(c *fiber.Ctx) error {
    // Parse the request body into PaymentRequest struct
    var request model.PaymentRequest
    if err := c.BodyParser(&request); err != nil {
        return err
    }

    // Generate a new UUID as the reference
    reference := uuid.New().String()

    // Make payment using Paystack service
    paymentUrl, err := paystackService.MakePayment(request.Amount, request.Email, nil, nil)
    if err != nil {
        return err
    }

    // Ensure that the payment URL is not empty
    if paymentUrl == "" {
        return errors.New("payment URL not provided by Paystack service")
    }

    // Create a new payment record
    payment := model.Payment{
        Amount:    request.Amount,
        Email:     request.Email,
        Reference: reference, // Use the generated UUID as the reference
    }

    // Save the payment record to the database
    if err := database.DB.Create(&payment).Error; err != nil {
        return err
    }

    // Redirect the user to the payment page
    return c.Redirect(paymentUrl)
}