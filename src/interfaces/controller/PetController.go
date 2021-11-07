package controller

import (
	"github.com/gofiber/fiber/v2"
	"pet-api/src/application/usecase"
)

type PetController struct {
	petUsecase usecase.IPetUsecase
}

func NewPetController(petUsecase usecase.IPetUsecase) PetController {
	petController := PetController{petUsecase: petUsecase}
	return petController
}

func (petController *PetController) CreatePet(c *fiber.Ctx) error {
	params := new(usecase.PetCreateParameter)

	var err error
	err = params.ParamsSetup(c)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = petController.petUsecase.CreatePet(params)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}
