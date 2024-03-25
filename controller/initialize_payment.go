package controller

import (
    "fmt"
    "os"
    "github.com/ADEMOLA200/Payment-Implementation/services"
)

func InitializePaystackService() (*services.PaystackService, error) {
	// Get the secret key from environment variable
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	if secretKey == "" {
		return nil, fmt.Errorf("secret key not found in environment variable")
	}

	// Create an instance of PaystackService
	paystackService := &services.PaystackService{}

	// Initialize PaystackService with the secret key
	paystackService.Initialize(secretKey)

	return paystackService, nil
}
