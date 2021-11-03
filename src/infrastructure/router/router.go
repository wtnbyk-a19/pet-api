package router

import (
	"github.com/labstack/echo"
	"pet-api/src/injector"
)

func RouterInit(echo *echo.Echo) {
	hogeController := injector.InjectHogeController()

	g := echo.Group("/hoge")
	{
		g.POST("/hogehoge", hogeController.CreateHoge())
	}

}
