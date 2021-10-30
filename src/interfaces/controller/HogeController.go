package controller

import (
	"docker-go-api/src/application/usecase"
	"docker-go-api/src/domain/model"
	"fmt"
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
		// TODO: 確認したら削除
		var request interface{}
		context.Bind(&request)

		fmt.Println(request)

		hoge := model.NewHoge(context.FormValue("hogehoge"))

		error := hogrController.hogeUsecase.Execute(hoge)
		if error != nil {
			return context.JSON(http.StatusBadRequest, error)
		}
		return context.JSON(http.StatusOK, error)
	}
}