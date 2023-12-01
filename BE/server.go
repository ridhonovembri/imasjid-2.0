package main

import (
	"github.com/gofiber/fiber/v2"
	"imasjid-2.0/database"
	"imasjid-2.0/routes"
)

func main() {

	database.ConnectDb()

	// database.MigrationInit()

	app := fiber.New()

	routes.RouteInit(app)

	app.Listen(":3000")
}
