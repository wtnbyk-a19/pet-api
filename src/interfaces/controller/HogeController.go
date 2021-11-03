package controller

import (
	"docker-go-api/src/application/usecase"
	"docker-go-api/src/domain/model"
	"github.com/labstack/echo"
	"net/http"
)

type HogeController struct {
	hogeUsecase usecase.IHogeUsecase
}

func NewHogeController(hogeUsecase usecase.IHogeUsecase) HogeController {
	hogeHandler := HogeController{hogeUsecase: hogeUsecase}
	return hogeHandler
}

func (hogrController *HogeController) CreateHoge() echo.HandlerFunc {
	return func(context echo.Context) error {
		// TODO: リクエスト構造体の定義
		var request struct {
			Hogehoge string `json:"hogehoge"`
		}
		context.Bind(&request)

		hoge := model.NewHoge(request.Hogehoge)

		error := hogrController.hogeUsecase.Execute(hoge)
		if error != nil {
			return context.JSON(http.StatusBadRequest, error)
		}
		return context.JSON(http.StatusOK, error)
	}
}
