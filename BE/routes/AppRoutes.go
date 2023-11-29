package routes

import (
	"github.com/gofiber/fiber/v2"
	"imasjid-2.0/controllers"
)

func RouteInit(app *fiber.App) {

	app.Get("/", controllers.TestController)

}
