package main

import (
	"fmt"
	"log"

	"github.com/ADEMOLA200/Payment-Implementation/controller"
	"github.com/ADEMOLA200/Payment-Implementation/database"
	_"github.com/ADEMOLA200/Payment-Implementation/handlers"
	"github.com/ADEMOLA200/Payment-Implementation/middlewares"
	routes "github.com/ADEMOLA200/Payment-Implementation/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
    // Define test secret key
    const paystackTestKey = "sk_test_b54a660ada3b30d7d3b9791e529a043f6f198f9e"

    database.ConnectDB()
    
    app := fiber.New()

	// Register handler for /api/makePayment endpoint
	//app.All("/api/makePayment/*", handlers.ProxyHandler)

	app.Use(middlewares.CorsMiddleware())
    app.Use(middlewares.LoggingMiddleware)
    app.Use(middlewares.ErrorHandlingMiddleware)
    
    routes.SetupRoutes(app)

    app.Listen(":3400")

    // Initialize Paystack service
    paystackService, err := controller.InitializePaystackService(paystackTestKey)
    if err != nil {
        log.Fatalf("Failed to set up Paystack service: %v", err)
    }

    fmt.Println("Secret Key:", paystackService.SecretKey)
}
