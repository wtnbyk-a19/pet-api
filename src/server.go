package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"pet-api/src/infrastructure/database"
	"pet-api/src/infrastructure/router"
)

func init() {
	err := godotenv.Load("go/api/env/dev.env")
	if err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	app := fiber.New()

	database.Init()
	defer func(conn *gorm.DB) {
		err := conn.Close()
		if err != nil {
			logrus.Fatal(err)
		}
	}(database.DbConn.Conn)

	router.Init(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
