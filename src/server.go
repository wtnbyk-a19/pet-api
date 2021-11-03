package main

import (
	"docker-go-api/src/infrastructure/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func init() {
	error := godotenv.Load("/go/api/.env")
	if error != nil {
		logrus.Fatal(error)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	echo := echo.New()

	// Routes
	router.RouterInit(echo)

	// Start server
	echo.Logger.Fatal(echo.Start(":3000"))
}
