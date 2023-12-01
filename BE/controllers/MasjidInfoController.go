package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"imasjid-2.0/database"
	"imasjid-2.0/models"
)

type MasjidInfoSerializer struct {
	// This is not the model, more like a serializer
	ID            uint   `json:"id"`
	MasjidName    string `json:"masjidName"`
	MasjidAddress string `json:"masjidAddress"`
	City          string `json:"city"`
	Province      string `json:"province"`
	BkmPIC        string `json:"bkmpic"`
	BkmPhone      string `json:"bkmPhone"`
}

func CreateResponse(m models.MasjidInfo) MasjidInfoSerializer {
	return MasjidInfoSerializer{
		ID:            m.ID,
		MasjidName:    m.MasjidName,
		MasjidAddress: m.MasjidAddress,
		City:          m.City,
		Province:      m.Province,
		BkmPIC:        m.BkmPIC,
		BkmPhone:      m.BkmPhone,
	}
}

func GetAll(c *fiber.Ctx) error {

	masjidInfo := []models.MasjidInfo{}
	database.Database.Db.Find(&masjidInfo)
	responses := []MasjidInfoSerializer{}
	for _, value := range masjidInfo {
		response := CreateResponse(value)
		responses = append(responses, response)
	}

	return c.Status(200).JSON(responses)

}

func FindById(id int, masjidInfo *models.MasjidInfo) error {

	database.Database.Db.Find(&masjidInfo, "id = ?", id)
	if masjidInfo.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func Create(c *fiber.Ctx) error {
	var masjidInfo models.MasjidInfo

	if err := c.BodyParser(&masjidInfo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&masjidInfo)
	response := CreateResponse(masjidInfo)

	return c.Status(200).JSON(response)
}

func Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var masjidInfo models.MasjidInfo

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = FindById(id, &masjidInfo)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateMasjidInfo struct {
		MasjidName    string `json:"masjidName"`
		MasjidAddress string `json:"masjidAddress"`
		City          string `json:"city"`
		Province      string `json:"province"`
		BkmPIC        string `json:"bkmpic"`
		BkmPhone      string `json:"bkmPhone"`
	}

	var updateData UpdateMasjidInfo

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	masjidInfo.MasjidName = updateData.MasjidName
	masjidInfo.MasjidAddress = updateData.MasjidAddress
	masjidInfo.City = updateData.City
	masjidInfo.Province = updateData.Province
	masjidInfo.BkmPIC = updateData.BkmPIC
	masjidInfo.BkmPhone = updateData.BkmPhone

	database.Database.Db.Save(&masjidInfo)

	response := CreateResponse(masjidInfo)

	return c.Status(200).JSON(response)

}

func Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var masjidInfo models.MasjidInfo

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = FindById(id, &masjidInfo)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&masjidInfo).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted Record")
}
