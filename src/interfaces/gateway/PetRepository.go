package gateway

import (
	"pet-api/src/domain/model"
	"pet-api/src/domain/repository"
	"pet-api/src/infrastructure/database"
)

type petRepository struct {
	dbConnection database.DbConnection
}

func NewPetRepository(dbConnection database.DbConnection) repository.IPetRepository {
	petRepository := petRepository{dbConnection}
	return &petRepository
}

func (petRepository *petRepository) Save(pet *model.Pet) (error error) {

	result := petRepository.dbConnection.Connection.Create(pet)
	return result.Error
}
