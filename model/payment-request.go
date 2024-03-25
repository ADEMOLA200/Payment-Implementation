package model

type PaymentRequest struct {
    Amount   int64                  `json:"amount"`
    Email    string                 `json:"email"`
    Metadata map[string]interface{} `json:"metadata"`
    // Add other fields related to subscription payment request
}