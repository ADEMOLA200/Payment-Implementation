package services

// import (
//     "encoding/json"
//     "fmt"

//     "github.com/stripe/stripe-go/paymentintent"
//     "github.com/stripe/stripe-go/v72"
// )

// type StripeService struct {
//     // You can initialize any necessary configuration here
//     apiKey string
// }

// func (ss *StripeService) Initialize(apiKey string) {
//     ss.apiKey = apiKey
//     stripe.Key = apiKey
// }

// func (ss *StripeService) MakePayment(amount int64, email string, metadata map[string]interface{}) error {
//     // Implement Stripe payment logic here
//     stripe.Key = ss.apiKey

//     // Construct the payment intent request body
//     params := &stripe.PaymentIntentParams{
//         Amount:      stripe.Int64(amount),
//         Currency:    stripe.String(string(stripe.CurrencyUSD)),
//         Description: stripe.String("Payment description"),
//     }

//     // Create a new payment intent
//     intent, err := paymentintent.New(params)
//     if err != nil {
//         return fmt.Errorf("failed to create payment intent: %v", err)
//     }

//     // Attach metadata to the payment intent
//     if metadata != nil {
//         metadataJSON, err := json.Marshal(metadata)
//         if err != nil {
//             return fmt.Errorf("failed to marshal metadata: %v", err)
//         }
//         intent.Metadata = map[string]string{"custom_metadata": string(metadataJSON)}
//     }

//     // Confirm the payment intent to charge the customer
//     _, err = paymentintent.Confirm(
//         intent.ID,
//         &stripe.PaymentIntentConfirmParams{
//             PaymentMethodOptions: &stripe.PaymentIntentPaymentMethodOptionsParams{
//                 Card: &stripe.PaymentIntentPaymentMethodOptionsCardParams{
//                     // Update with customer payment method ID if necessary
//                 },
//             },
//         },
//     )
//     if err != nil {
//         return fmt.Errorf("failed to confirm payment intent: %v", err)
//     }

//     // Retrieve payment intent to check status
//     pi, err := paymentintent.Get(intent.ID, nil)
//     if err != nil {
//         return fmt.Errorf("failed to retrieve payment intent: %v", err)
//     }

//     // Handle the payment intent's status and return any errors
//     switch pi.Status {
//     case stripe.PaymentIntentStatusSucceeded:
//         return nil
//     case stripe.PaymentIntentStatusProcessing:
//         return fmt.Errorf("payment is processing")
//     default:
//         return fmt.Errorf("payment failed with status: %s", pi.Status)
//     }
// }
