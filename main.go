package main

import (
	"github.com/ADEMOLA200/Payment-Implementation/database"
	_"github.com/ADEMOLA200/Payment-Implementation/handlers"
	"github.com/ADEMOLA200/Payment-Implementation/middlewares"
	routes "github.com/ADEMOLA200/Payment-Implementation/route"
	"github.com/gofiber/fiber/v2"
)

func main() {

    database.ConnectDB()
    
    app := fiber.New()

	app.Use(middlewares.CorsMiddleware())
    app.Use(middlewares.LoggingMiddleware)
    app.Use(middlewares.ErrorHandlingMiddleware)
    
    routes.SetupRoutes(app)

    app.Listen(":3400")
}
