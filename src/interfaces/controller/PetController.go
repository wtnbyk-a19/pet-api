package controller

import (
	"github.com/gofiber/fiber"
	"pet-api/src/application/usecase"
)

type PetController struct {
	petUsecase usecase.IPetUsecase
}

func NewPetController(petUsecase usecase.IPetUsecase) PetController {
	petHandler := PetController{petUsecase: petUsecase}
	return petHandler
}

func (petController *PetController) SavePet(c *fiber.Ctx) {

	//err := c.Bind(&request)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, err)
	//}
	//
	//pet := model.NewPet(request.name)
	//
	//err = petController.petUsecase.SavePet(pet)
	//if err != nil {
	//	return context.JSON(http.StatusBadRequest, err)
	//}
	//return context.JSON(http.StatusOK, err)
}
