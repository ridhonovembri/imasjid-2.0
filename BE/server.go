package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"imasjid-2.0/database"
	"imasjid-2.0/routes"
)

func main() {

	//initiate fiber
	app := fiber.New()

	//connect to DB
	database.ConnectDb()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	routes.RouteInit(app)

	log.Fatal(app.Listen(":3000"))

}
