package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	loggerInit(app)
	database.Init()
	defer func(conn *gorm.DB) {
		err := conn.Close()
		if err != nil {
			logrus.Fatal(err)
		}
	}(database.DbConn.Conn)

	router.Init(app)

	app.Listen(":3000")
}

func loggerInit(app *fiber.App) {
	var loggerConfig = logger.Config{
		Next:       nil,
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",
	}

	app.Use(loggerConfig)
}
