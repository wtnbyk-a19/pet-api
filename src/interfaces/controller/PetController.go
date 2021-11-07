package controller

import (
	"github.com/gofiber/fiber"
	"net/http"
	"pet-api/src/application/usecase"
)

type PetController struct {
	petUsecase usecase.IPetUsecase
}

func NewPetController(petUsecase usecase.IPetUsecase) PetController {
	petController := PetController{petUsecase: petUsecase}
	return petController
}

func (petController *PetController) CreatePet(c *fiber.Ctx) {
	params := new(usecase.PetCreateParameter)

	var err error
	err = params.ParamsSetup(c)
	if err != nil {
		c.Status(http.StatusBadRequest).Send(err)
		return
	}

	err = petController.petUsecase.CreatePet(params)
	if err != nil {
		c.Status(http.StatusBadRequest).Send(err)
		return
	}

	c.Status(http.StatusOK)
}
