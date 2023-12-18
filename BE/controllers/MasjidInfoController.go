package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"imasjid-2.0/database"
	"imasjid-2.0/models"
)

// type MasjidInfoSerializer struct {
// 	// This is not the model, more like a serializer
// 	ID            uint   `json:"id"`
// 	MasjidName    string `json:"masjidName"`
// 	MasjidAddress string `json:"masjidAddress"`
// 	City          string `json:"city"`
// 	Province      string `json:"province"`
// 	BkmPIC        string `json:"bkmpic"`
// 	BkmPhone      string `json:"bkmPhone"`
// }

// func CreateResponse(m models.MasjidInfo) MasjidInfoSerializer {
// 	return MasjidInfoSerializer{
// 		ID:            m.ID,
// 		MasjidName:    m.MasjidName,
// 		MasjidAddress: m.MasjidAddress,
// 		City:          m.City,
// 		Province:      m.Province,
// 		BkmPIC:        m.BkmPIC,
// 		BkmPhone:      m.BkmPhone,
// 	}
// }

func GetAllInfo(c *fiber.Ctx) error {

	var data []models.MasjidInfo

	result := database.Database.Find(&data)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.Status(200).JSON(data)

	// masjidInfo := []models.MasjidInfo{}
	// database.Database.Db.Find(&masjidInfo)

	// responses := []MasjidInfoSerializer{}

	// for _, value := range masjidInfo {
	// 	response := CreateResponse(value)
	// 	responses = append(responses, response)
	// }

	// // fmt.Println("responses", responses)

	// return c.Status(200).JSON(responses)

}

func GetInfo(c *fiber.Ctx) error {

	var data models.MasjidInfo

	result := database.Database.First(&data)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.Status(200).JSON(data)
}

func GetInfoById(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.MasjidInfo

	result := database.Database.Find(&data, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&data)
}

// func GetInfo(c *fiber.Ctx) {

// 	id := c.Params("id")

// 	var data model.masjidInfo

// 	result := database.Database.Db.Find(&data, "id = ?", id)

// 	if result.Error |= nil {
// 		log.Println("Record does NOT exist")
// 	}

// 	// return nil

// 	// db := database.Database.Db()
// 	// newid, err := strconv.Atoi(id)
// 	// if err != nil {
// 	// 	c.JSON(400, gin.H{
// 	// 		"Error": "ID has to be integer",
// 	// 	})
// 	// 	return
// 	// }

// 	// var book models.Book
// 	// err = db.First(&book, newid).Error
// 	// if err != nil {
// 	// 	c.JSON(400, gin.H{
// 	// 		"Error": "Book not found: " + err.Error(),
// 	// 	})

// 	// 	return
// 	// }
// 	c.JSON(200, data)
// }

// func FindById(id int, masjidInfo *models.MasjidInfo) error {

// 	database.Database.Db.Find(&masjidInfo, "id = ?", id)
// 	if masjidInfo.ID == 0 {
// 		return errors.New("record does not exist")
// 	}

// 	return nil
// }

// func FindById(id int, c *fiber.Ctx) error {
// 	// id := c.Params("id")

// 	var data models.MasjidInfo
// 	err := database.Database.Db.Find(&data, "id = ?", id).Error

// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"message": "data not found",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status":  true,
// 		"message": "success",
// 		"data":    data,
// 	})
// }

// func Create(c *fiber.Ctx) error {
// 	var masjidInfo models.MasjidInfo

// 	if err := c.BodyParser(&masjidInfo); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	database.Database.Db.Create(&masjidInfo)
// 	response := CreateResponse(masjidInfo)

// 	return c.Status(200).JSON(response)
// }

func CreateInfo(c *fiber.Ctx) error {

	// var data models.MasjidInfo

	// if err := c.BodyParser(&data); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": err.Error(),
	// 	})
	// }

	// if err := database.Database.Create(&data).Error; err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": err.Error(),
	// 	})
	// }

	data := new(models.MasjidInfo)

	if err := c.BodyParser(data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Database.Create(&data)

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "success",
		// "data":    data,
	})
}

// func Update(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")

// 	var masjidInfo models.MasjidInfo

// 	if err != nil {
// 		return c.Status(400).JSON("Please ensure that :id is an integer")
// 	}

// 	err = FindById(id, &masjidInfo)
// 	// err = FindById(id, c)

// 	if err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	type UpdateMasjidInfo struct {
// 		MasjidName    string `json:"masjidName"`
// 		MasjidAddress string `json:"masjidAddress"`
// 		City          string `json:"city"`
// 		Province      string `json:"province"`
// 		BkmPIC        string `json:"bkmpic"`
// 		BkmPhone      string `json:"bkmPhone"`
// 	}

// 	var updateData UpdateMasjidInfo

// 	if err := c.BodyParser(&updateData); err != nil {
// 		return c.Status(500).JSON(err.Error())
// 	}

// 	masjidInfo.MasjidName = updateData.MasjidName
// 	masjidInfo.MasjidAddress = updateData.MasjidAddress
// 	masjidInfo.City = updateData.City
// 	masjidInfo.Province = updateData.Province
// 	masjidInfo.BkmPIC = updateData.BkmPIC
// 	masjidInfo.BkmPhone = updateData.BkmPhone

// 	database.Database.Db.Save(&masjidInfo)

// 	response := CreateResponse(masjidInfo)

// 	return c.Status(200).JSON(response)

// }

func UpdateInfo(c *fiber.Ctx) error {

	// id := c.Params("id")

	// var data models.MasjidInfo
	// if err := c.BodyParser(&data); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": err.Error(),
	// 	})
	// }

	// if database.Database.Db.Where("id = ?", id).Updates(&data).RowsAffected == 0 {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "Record is NOT Updated",
	// 	})
	// }

	data := new(models.MasjidInfo)
	id := c.Params("id")

	if err := c.BodyParser(data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Database.Where("id = ?", id).Updates(&data)

	return c.Status(200).JSON(fiber.Map{
		"message": "Record has been UPDATED",
	})
}

func DeleteInfo(c *fiber.Ctx) error {

	// id := c.Params("id")

	// var data models.MasjidInfo

	// if database.Database.Db.Delete(&data, id).RowsAffected == 0 {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"message": "Record is NOT Deleted",
	// 	})
	// }

	id := c.Params("id")
	var data models.MasjidInfo

	result := database.Database.Delete(&data, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Record has been DELETED",
	})

}

// func Delete(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")

// 	var masjidInfo models.MasjidInfo

// 	if err != nil {
// 		return c.Status(400).JSON("Please ensure that :id is an integer")
// 	}

// 	err = FindById(id, &masjidInfo)
// 	// err = FindById(id, c)

// 	if err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	if err = database.Database.Db.Delete(&masjidInfo).Error; err != nil {
// 		return c.Status(404).JSON(err.Error())
// 	}
// 	return c.Status(200).JSON("Successfully deleted record")
// }
