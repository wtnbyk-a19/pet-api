package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	e := echo.New()

	//// Routes
	//router.RouterInit(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo World!!")
	})

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
