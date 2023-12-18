package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"imasjid-2.0/database"
	"imasjid-2.0/models"
)

func GetAllMarquee(c *fiber.Ctx) error {

	var data []models.MasjidMarquee

	result := database.Database.Find(&data)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.Status(200).JSON(data)
}

func GetMarqueeById(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.MasjidMarquee

	result := database.Database.Find(&data, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&data)
}

func CreateMarquee(c *fiber.Ctx) error {

	data := new(models.MasjidMarquee)

	if err := c.BodyParser(data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Database.Create(&data)

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "success",
	})
}

func UpdateMarquee(c *fiber.Ctx) error {

	data := new(models.MasjidMarquee)
	id := c.Params("id")

	if err := c.BodyParser(data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Database.Where("id = ?", id).Updates(&data)

	return c.Status(200).JSON(fiber.Map{
		"message": "Record has been UPDATED",
	})
}

func DeleteMarquee(c *fiber.Ctx) error {

	id := c.Params("id")
	var data models.MasjidMarquee

	result := database.Database.Delete(&data, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Record has been DELETED",
	})
}
