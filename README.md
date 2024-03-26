# Payment Implementation

This project implements a payment system using the Paystack and Stripe APIs for processing payments. It consists of both frontend and backend components.

## Features

- Collect payment details from users.
- Process payments securely using the Paystack and Stripe APIs.
- Display payment confirmation to users.
- Integration with backend server for processing payments and handling business logic.

## Technologies Used

- Frontend:
  - HTML
  - CSS
  - JavaScript

- Backend:
  - Go (Golang)
  - Fiber framework
  - GORM for database interaction

- Payment Gateways:
  - Paystack API
  - Stripe API

![Editor](https://imgur.com/NnHqlFR)

![Postman](https://imgur.com/Cx4ZXnw)


## Getting Started


1. Clone the repository:

```bash

1. Clone the repository:

```bash
git clone <repository_url>
```

2. Set up the backend server:
   - Navigate to the `backend` directory.
   - Install dependencies: `go mod tidy`.
   - Configure the database connection in `database/connectDB.go`.
   - Start the server: `go run main.go`.

3. Set up the frontend:
   - Navigate to the `frontend` directory.
   - Open `index.html` in a web browser or host it on a web server.

4. Test the payment system:
   - Enter payment details in the frontend form.
   - Submit the form to process the payment.
   - Check the backend logs for payment processing details.

![Payment System](https://imgur.com/VXD0uda)

![Payment System](https://imgur.com/HITKEPc)

## Configuration

- Backend:
  - Configure the database connection in `database/connectDB.go`.
  - Set the Paystack secret key in the environment variable or directly in the code.

- Frontend:
  - Update the Paystack public key in the JavaScript code (`app.js`).

## License

This project is licensed under the [MIT License](LICENSE).

```