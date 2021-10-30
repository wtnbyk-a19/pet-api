package router

import (
	"docker-go-api/src/injector"
	"github.com/labstack/echo"
)

func RouterInit(echo *echo.Echo) {
	hogeController := injector.InjectHogeController()

	g := echo.Group("/hoge")
	{
		g.POST("hogehoge", hogeController.CreateHoge())
	}

}