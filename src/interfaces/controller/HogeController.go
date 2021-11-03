package controller

import (
	"../../application/usecase"
	"../../domain/model"
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

func (hogeController *HogeController) CreateHoge() echo.HandlerFunc {
	return func(context echo.Context) error {
		// TODO: リクエスト構造体の定義
		var request struct {
			Hogehoge string `json:"hogehoge"`
		}

		err := context.Bind(&request)
		if err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}

		hoge := model.NewHoge(request.Hogehoge)

		err = hogeController.hogeUsecase.Execute(hoge)
		if err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}
		return context.JSON(http.StatusOK, err)
	}
}
