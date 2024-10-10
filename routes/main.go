package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {

	// Auth
	app.Post("/register", Register)
	app.Post("/login", Login)

	// User
	app.Get("/users", GetUsers)
	app.Get("/users/:id", GetUser)
	app.Post("/users", CreateUser)
	app.Put("/users/:id", UpdateUser)
	app.Delete("/users/:id", DeleteUser)
}
