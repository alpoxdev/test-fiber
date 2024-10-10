package routes

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"test-fiber/database"
	"test-fiber/lib"
	"test-fiber/models"
)

type RegisterRequest struct {
	Nickname string `json:"nickname" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func Register(c *fiber.Ctx) error {
	var request RegisterRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if err := validator.New().Struct(request); err != nil {
		// 유효성 검사 에러 추출
		validationErrors := err.(validator.ValidationErrors)
		errors := make(map[string]string)
		for _, fieldErr := range validationErrors {
				errors[fieldErr.Field()] = fieldErr.Error()
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": errors,
		})
	}

	email := request.Email
	password := request.Password
	nickname := request.Nickname

	hashedPassword, err := lib.HashPassword(password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	existingUser := models.User {}

	database.DB.Where("email = ?", email).First(&existingUser)
	if existingUser.ID != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	database.DB.Where("nickname = ?", nickname).First(&existingUser)
	if existingUser.ID != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	user := models.User{
		Email: email,
		Password: hashedPassword,
		Nickname: nickname,
	}

	database.DB.Create(&user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create token",
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "TOKEN"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour) // 24 hours
	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"id": user.ID,
	})
}

func Login(c *fiber.Ctx) error {
	var request LoginRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if err := validator.New().Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	email := request.Email
	password := request.Password

	existingUser := models.User {}

	database.DB.Where("email = ?", email).First(&existingUser)
	if existingUser.ID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	hashedPassword := existingUser.Password
	isPasswordCorrect := lib.CheckPasswordHash(password, hashedPassword)
	if !isPasswordCorrect {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": existingUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create token",
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "TOKEN"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "User logged in successfully",
		"id": existingUser.ID,
	})
}