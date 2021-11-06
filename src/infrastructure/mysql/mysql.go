package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type DbConnection struct {
	Connection *gorm.DB
}

func NewDbConnection() *DbConnection {
	connection := connect()

	dbConnection := new(DbConnection)
	dbConnection.Connection = connection

	return dbConnection
}

func connect() (connection *gorm.DB) {
	err := godotenv.Load("go/api/env/dev.env")
	if err != nil {
		logrus.Fatal(err)
	}

	connection, err = gorm.Open("mysql",
		os.Getenv("DB_USERNAME")+":"+
			os.Getenv("DB_PASSWORD")+"@tcp("+
			os.Getenv("DB_HOST")+")/"+
			os.Getenv("DB_DATABASE")+
			"?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		logrus.Fatal(err)
	}

	return connection
}
