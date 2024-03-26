package controller

import (
    "github.com/ADEMOLA200/Payment-Implementation/services"
    "errors"
)

// InitializePaystackService initializes the Paystack service with the provided secret key
func InitializePaystackService(secretKey string) (*services.PaystackService, error) {
    // Validate secret key
    if secretKey == "" {
        return nil, errors.New("secret key cannot be empty")
    }

    // Create an instance of PaystackService
    paystackService := &services.PaystackService{}

    // Initialize PaystackService with the secret key
    paystackService.Initialize(secretKey)

    return paystackService, nil
}
