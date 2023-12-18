package routes

import (
	"github.com/gofiber/fiber/v2"
	"imasjid-2.0/controllers"
)

func RouteInit(app *fiber.App) {

	app.Get("/api/masjidConfigs", controllers.GetAllConfig)
	app.Get("/api/masjidConfig", controllers.GetConfig)
	app.Get("/api/masjidConfigById/:id", controllers.GetConfigById)
	app.Post("/api/masjidConfig/post", controllers.CreateConfig)
	app.Put("/api/masjidConfig/update/:id", controllers.UpdateConfig)
	app.Delete("/api/masjidConfig/delete/:id", controllers.DeleteConfig)

	app.Get("/api/masjidHadist", controllers.GetAllHadist)
	app.Get("/api/masjidHadistRandom", controllers.GetHadistByRandom)
	app.Get("/api/masjidHadistById/:id", controllers.GetHadistById)
	app.Post("/api/masjidHadist/post", controllers.CreateHadist)
	app.Put("/api/masjidHadist/update/:id", controllers.UpdateHadist)
	app.Delete("/api/masjidHadist/delete/:id", controllers.DeleteHadist)

	app.Get("/api/masjidInfos", controllers.GetAllInfo)
	app.Get("/api/masjidInfo", controllers.GetInfo)
	app.Get("/api/masjidInfoById/:id", controllers.GetInfoById)
	app.Post("/api/masjidInfo/post", controllers.CreateInfo)
	app.Put("/api/masjidInfo/update/:id", controllers.UpdateInfo)
	app.Delete("/api/masjidInfo/delete/:id", controllers.DeleteInfo)

	app.Get("/api/masjidMarquee", controllers.GetAllMarquee)
	app.Get("/api/masjidMarqueeById/:id", controllers.GetMarqueeById)
	app.Post("/api/masjidMarquee/post", controllers.CreateMarquee)
	app.Put("/api/masjidMarquee/update/:id", controllers.UpdateMarquee)
	app.Delete("/api/masjidMarquee/delete/:id", controllers.DeleteMarquee)

	app.Get("/api/slides", controllers.GetAllSlide)
	app.Get("/api/getSlides", controllers.GetSlides)
	app.Get("/api/slideById/:id", controllers.GetSlideById)
	app.Post("/api/slide/post", controllers.CreateSlide)
	app.Put("/api/slide/update/:id", controllers.UpdateSlide)
	app.Delete("/api/slide/delete/:id", controllers.DeleteSlide)

}
