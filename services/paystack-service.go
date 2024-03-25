package services

import (
	"encoding/json"
	"fmt"

	"github.com/ADEMOLA200/Payment-Implementation/model"
	"github.com/cadanapay/paystack-go"
)


type PaystackService struct {
    SecretKey string // Change secretKey to SecretKey
    Client    *paystack.Client
}

func (ps *PaystackService) Initialize(secretKey string) {
    ps.SecretKey = secretKey
    ps.Client = paystack.NewClient(secretKey, nil)
}

func (ps *PaystackService) MakePayment(amount int64, email string, metadata map[string]interface{}, card *model.CardDetails) (*model.Payment, error) {
    // Convert amount from int64 to float32 (kobo)
    amountKobo := float32(amount) / 100 // Convert Naira to kobo

    // Convert metadata map to JSON string
    metadataJSON, err := json.Marshal(metadata)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal metadata to JSON: %v", err)
    }

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
        return nil, fmt.Errorf("failed to initialize transaction: %v", err)
    }

    // Extract the reference from the transaction response
    reference, ok := transactionResponse["data"].(map[string]interface{})["reference"].(string)
    if !ok {
        return nil, fmt.Errorf("failed to extract reference from transaction response")
    }

    // Verify the transaction, which creates a charge on the user's account
    _, err = ps.Client.Transaction.Verify(reference)
    if err != nil {
        return nil, fmt.Errorf("failed to verify transaction: %v", err)
    }

    // Create a Payment object with the retrieved details
    payment := model.NewPayment()
    payment.Amount = amount
    payment.Email = email
    payment.Metadata = string(metadataJSON) // Assign the JSON string to Metadata
    payment.Reference = reference

    return payment, nil
}

