package routes

import (
	"test-fiber/database"
	"test-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Find(&users)

	totalCount := int64(0)
	database.DB.Model(&models.User{}).Count(&totalCount)

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
