package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"pet-api/src/application/usecase"
	"pet-api/src/domain/model"
)

type PetController struct {
	petUsecase usecase.IPetUsecase
}

func NewPetController(petUsecase usecase.IPetUsecase) PetController {
	petHandler := PetController{petUsecase: petUsecase}
	return petHandler
}

func (petController *PetController) SavePet() echo.HandlerFunc {
	return func(context echo.Context) error {
		// TODO: リクエスト構造体の定義
		var request struct {
			name string `json:"name"`
		}
		err := context.Bind(&request)
		if err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}

		pet := model.NewPet(request.name)

		err = petController.petUsecase.SavePet(pet)
		if err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}
		return context.JSON(http.StatusOK, err)
	}
}
