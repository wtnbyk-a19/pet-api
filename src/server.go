package main

import (
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"pet-api/src/infrastructure/router"
)

func init() {
	err := godotenv.Load("go/api/env/dev.env")
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	app := fiber.New()

	router.Init(app)

	err := app.Listen(3000)
	if err != nil {
		return
	}
}
