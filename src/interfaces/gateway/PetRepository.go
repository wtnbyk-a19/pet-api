package gateway

import (
	"github.com/sirupsen/logrus"
	"pet-api/src/domain/model"
	"pet-api/src/domain/repository"
	"pet-api/src/infrastructure/database"
)

type petRepository struct {
	dbConn database.DbConnection
}

func NewPetRepository(dbConnection database.DbConnection) repository.IPetRepository {
	petRepository := petRepository{dbConnection}
	return &petRepository
}

func (petRepository *petRepository) Create(pet *model.Pet) (error error) {
	result := petRepository.dbConn.Conn.Create(pet)

	logrus.Println(result)

	return result.Error
}
