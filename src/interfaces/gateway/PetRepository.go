package gateway

import (
	"pet-api/src/domain/model"
	"pet-api/src/domain/repository"
	"pet-api/src/infrastructure/mysql"
)

type petRepository struct {
	dbConnection mysql.DbConnection
}

func NewPetRepository(dbConnection mysql.DbConnection) repository.IPetRepository {
	petRepository := petRepository{dbConnection}
	return &petRepository
}

func (petRepository *petRepository) Save(pet *model.Pet) (error error) {

	result := petRepository.dbConnection.Connection.Create(pet)
	return result.Error
}
