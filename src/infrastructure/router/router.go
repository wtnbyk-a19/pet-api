package router

import (
	"github.com/gofiber/fiber/v2"
	"pet-api/src/injector"
)

func Init(app *fiber.App) {
	petController := injector.InjectPetController()

	app.Post("/pet", petController.CreatePet)
}
