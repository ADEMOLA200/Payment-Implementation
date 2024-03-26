package model

import "encoding/json"

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

// SerializeMetadata serializes the metadata map to a JSON string
func SerializeMetadata(metadata map[string]interface{}) (string, error) {
    jsonBytes, err := json.Marshal(metadata)
    if err != nil {
        return "", err
    }
    return string(jsonBytes), nil
}

// DeserializeMetadata deserializes the JSON string to a metadata map
func DeserializeMetadata(metadataStr string) (map[string]interface{}, error) {
    var metadata map[string]interface{}
    if err := json.Unmarshal([]byte(metadataStr), &metadata); err != nil {
        return nil, err
    }
    return metadata, nil
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
