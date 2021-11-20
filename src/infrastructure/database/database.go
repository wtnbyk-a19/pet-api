package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"pet-api/src/domain/Pet"
)

type DbConnection struct {
	Conn *gorm.DB
}

var (
	DbConn DbConnection
)

func Init() {
	err := godotenv.Load("go/api/env/dev.env")
	if err != nil {
		logrus.Fatal(err)
	}

	DbConn.Conn, err = gorm.Open("mysql",
		os.Getenv("DB_USERNAME")+":"+
			os.Getenv("DB_PASSWORD")+"@tcp("+
			os.Getenv("DB_HOST")+")/"+
			os.Getenv("DB_DATABASE")+
			"?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		logrus.Fatal(err)
	}

	migrate()
}

func migrate() {
	DbConn.Conn.AutoMigrate(&Pet.Pet{})
}
