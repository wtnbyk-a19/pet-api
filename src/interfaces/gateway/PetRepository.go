package gateway

import (
	"pet-api/src/domain/model"
	"pet-api/src/domain/repository"
	"pet-api/src/infrastructure/mysql"
)

type petRepository struct {
	dbConnection mysql.DbConnection
}

func NewpetRepository(dbConnection mysql.DbConnection) repository.IPetRepository {
	petRepository := petRepository{dbConnection}
	return &petRepository
}

func (petRepository *petRepository) Persist(pet *model.Pet) (error error) {
	result := petRepository.dbConnection.Connection.Create(pet)
	return result.Error
}
