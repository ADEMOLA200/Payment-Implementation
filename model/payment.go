package model

type Payment struct {
    Amount    int64  `json:"amount" gorm:"null"`
    Email     string `json:"email" gorm:"null"`
    Metadata  string `json:"metadata"` // Change type to string
    Reference string `json:"reference" gorm:"null"`
}

func NewPayment() *Payment {
    return &Payment{
        Metadata: "", // Initialize Metadata as an empty string
    }
}


type CardDetails struct {
    CardNumber   string `json:"card_number"`
    ExpiryMonth  string `json:"expiry_month"`
    ExpiryYear   string `json:"expiry_year"`
    CVV          string `json:"cvv"`
}

// // ToPaystackCard converts model.CardDetails to paystack.CardDetails format
// func (cd *CardDetails) ToPaystackCard() *paystack.CardDetails {
//     return &paystack.CardDetails{
//         Number:    cd.CardNumber,
//         ExpMonth:  cd.ExpiryMonth,
//         ExpYear:   cd.ExpiryYear,
//         CVC:       cd.CVV,
//     }
// }
