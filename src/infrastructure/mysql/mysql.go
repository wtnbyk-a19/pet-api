package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type DbConnection struct {
	Connection *gorm.DB
}

func NewDbConnection() *DbConnection {
	connection := dbConnect()

	dbConnection := new(DbConnection)
	dbConnection.Connection = connection

	return dbConnection
}

func dbConnect() (connection *gorm.DB) {
	error := godotenv.Load(".env")
	if error != nil {
		logrus.Fatal(error)
	}

	connection, error = gorm.Open("mysql",
		os.Getenv("DB_USERNAME") + ":" +
		os.Getenv("DB_PASSWORD") + "@tcp)" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_DATABASE") +
		"?charset=utf8mb4&parseTime=True&loc=local")

	if error != nil {
		logrus.Fatal(error)
	}

	return connection
}