package controller

import (
	"github.com/gofiber/fiber/v2"
	"pet-api/src/application/Pet"
)

type PetController struct {
	petUsecase Pet.IPetUsecase
}

func NewPetController(petUsecase Pet.IPetUsecase) PetController {
	petController := PetController{petUsecase: petUsecase}
	return petController
}

func (petController *PetController) CreatePet(c *fiber.Ctx) error {
	params := new(Pet.PetCreateParameter)

	var err error
	err = params.ParamsSetup(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err = petController.petUsecase.CreatePet(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(params)
}
