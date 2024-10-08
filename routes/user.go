package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"

	"test-fiber/database"
	"test-fiber/models"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	var totalCount int64

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&users).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.User{}).Count(&totalCount).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Error("GetUsers 트랜잭션 실패: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "사용자 정보를 가져오는 데 실패했습니다",
		})
	}
	
	log.Info("GetUsers TotalCount: ", totalCount)

	return c.JSON(fiber.Map{
		"rows": users,
		"total": totalCount,
	})
}

func GetUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")
	database.DB.Find(&user, "id = ?", id)
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	database.DB.Create(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")
	database.DB.Find(&user, "id = ?", id)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Save(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")
	database.DB.Find(&user, "id = ?", id)
	database.DB.Delete(&user)	
	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
