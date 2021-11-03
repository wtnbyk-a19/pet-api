package router

import (
	"github.com/labstack/echo"
	"pet-api/src/injector"
)

func RouterInit(echo *echo.Echo) {
	petController := injector.InjectPetController()

	g := echo.Group("/pet")
	{
		g.POST("/petpet", petController.CreatePet())
	}

}
