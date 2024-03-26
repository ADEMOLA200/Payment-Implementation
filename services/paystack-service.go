package services

// Inside your services package

import (
	"errors"
	"fmt"

	"github.com/ADEMOLA200/Payment-Implementation/model"
	"github.com/cadanapay/paystack-go"
)

// Inside your services package
type PaystackService struct {
    SecretKey string          // Secret key for Paystack API
    Client    *paystack.Client // Paystack client
	DebugMode bool            // Flag to indicate debug mode

}


// Modify the Initialize method to accept the secret key as an argument
func (ps *PaystackService) Initialize(secretKey string) {
    // Initialize Paystack client with the self.SecretKey
    ps.SecretKey = secretKey
    ps.Client = paystack.NewClient(secretKey, nil)
}

func (ps *PaystackService) MakePayment(amount int64, email string, metadata map[string]interface{}, card *model.CardDetails) (string, error) {
    // Convert amount from int64 to float32 (kobo)
    amountKobo := float32(amount) * 100 // Convert Naira to kobo

    // Construct the transaction request body
    transactionRequest := &paystack.TransactionRequest{
        Amount:   amountKobo,
        Email:    email,
        Metadata: metadata, // Assign the metadata map directly
        //Card:     card.ToPaystackCard(), // Convert model.CardDetails to paystack.CardDetails
    }

    // Initialize a new transaction
    transactionResponse, err := ps.Client.Transaction.Initialize(transactionRequest)
    if err != nil {
        return "", fmt.Errorf("failed to initialize transaction: %v", err)
    }

    // Log the transactionResponse for debugging
    fmt.Printf("Transaction Response: %+v\n", transactionResponse)

    // Check if the transactionResponse contains the authorization URL
    paymentUrl, ok := transactionResponse["authorization_url"].(string)
    if !ok || paymentUrl == "" {
        return "", errors.New("authorization URL not provided in transaction response")
    }

    // Return the payment URL
    return paymentUrl, nil
}
