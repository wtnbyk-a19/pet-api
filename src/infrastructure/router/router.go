package router

import (
	"github.com/labstack/echo"
	"pet-api/src/injector"
)

func Init(echo *echo.Echo) {
	petController := injector.InjectPetController()

	g := echo.Group("/pet")
	{
		g.POST("/edit", petController.SavePet())
	}

}
