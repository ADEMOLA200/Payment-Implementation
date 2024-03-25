package main

import (
	"fmt"

	"github.com/ADEMOLA200/Payment-Implementation/controller"
	"github.com/ADEMOLA200/Payment-Implementation/database"
	"github.com/ADEMOLA200/Payment-Implementation/middlewares"
	routes "github.com/ADEMOLA200/Payment-Implementation/route"
	"github.com/gofiber/fiber/v2"
)
func main() {

	database.ConnectDB()
	
    app := fiber.New()

	app.Use(middlewares.LoggingMiddleware)
	app.Use(middlewares.ErrorHandlingMiddleware())
	app.Use(middlewares.CorsMiddleware())

    routes.SetupRoutes(app)

    app.Listen(":3400")

	// Initialize PaystackService
	paystackService, err := controller.InitializePaystackService()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Output the secret key
	fmt.Println("Secret Key:", paystackService.SecretKey)
}
