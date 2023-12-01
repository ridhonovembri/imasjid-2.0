package routes

import (
	"github.com/gofiber/fiber/v2"
	"imasjid-2.0/controllers"
)

func RouteInit(app *fiber.App) {

	app.Get("/", controllers.TestController)
	app.Get("/api/masjidInfo", controllers.GetAll)
	app.Post("/api/masjidInfo/post", controllers.Create)
	app.Put("/api/masjidInfo/update/:id", controllers.Update)
	app.Delete("/api/masjidInfo/delete/:id", controllers.Delete)

}
