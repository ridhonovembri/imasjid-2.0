package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"imasjid-2.0/database"
	"imasjid-2.0/models"
)

func GetAllHadist(c *fiber.Ctx) error {

	var data []models.MasjidHadist

	result := database.Database.Find(&data)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.Status(200).JSON(data)
}

func GetHadistByRandom(c *fiber.Ctx) error {

	var data []models.MasjidHadist

	result := database.Database.Raw("SELECT * FROM masjid_hadist ORDER BY RANDOM() LIMIT 1").Scan(&data)
	// result := database.Database.First(&data)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&data)
}

func GetHadistById(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.MasjidHadist

	result := database.Database.Find(&data, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&data)
}

func CreateHadist(c *fiber.Ctx) error {

	data := new(models.MasjidHadist)

	if err := c.BodyParser(data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Database.Create(&data)

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "success",
	})
}

func UpdateHadist(c *fiber.Ctx) error {

	data := new(models.MasjidHadist)
	id := c.Params("id")

	if err := c.BodyParser(data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Database.Where("id = ?", id).Updates(&data)

	return c.Status(200).JSON(fiber.Map{
		"message": "Record has been UPDATED",
	})
}

func DeleteHadist(c *fiber.Ctx) error {

	id := c.Params("id")
	var data models.MasjidHadist

	result := database.Database.Delete(&data, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Record has been DELETED",
	})
}
