package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"imasjid-2.0/database"
	"imasjid-2.0/models"
)

func GetAllConfig(c *fiber.Ctx) error {

	var data models.MasjidConfig

	result := database.Database.Find(&data)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.Status(200).JSON(data)
}

func GetConfig(c *fiber.Ctx) error {
	var data models.MasjidConfig

	result := database.Database.First(&data)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.Status(200).JSON(&data)
}

func GetConfigById(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.MasjidConfig

	result := database.Database.Find(&data, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&data)
}

func CreateConfig(c *fiber.Ctx) error {

	data := new(models.MasjidConfig)

	if err := c.BodyParser(data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Database.Create(&data)

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "success",
	})
}

func UpdateConfig(c *fiber.Ctx) error {

	data := new(models.MasjidConfig)
	id := c.Params("id")

	// var jsonData, err = json.Marshal(data)
	// if err != nil {
	// 	return c.Status(503).SendString(err.Error())
	// }

	// if err := json.Marshal(data); err != nil {
	// 	return c.Status(503).SendString(err.Error())
	// }

	if err := c.BodyParser(data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Database.Where("id = ?", id).Updates(&data)

	return c.Status(200).JSON(fiber.Map{
		"message": "Record has been UPDATED",
	})
}

func DeleteConfig(c *fiber.Ctx) error {

	id := c.Params("id")
	var data models.MasjidConfig

	result := database.Database.Delete(&data, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Record has been DELETED",
	})
}
