package gateway

import (
	"pet-api/src/domain/Pet"
	"pet-api/src/infrastructure/database"
)

type petRepository struct {
	dbConn database.DbConnection
}

func NewPetRepository(dbConnection database.DbConnection) Pet.IPetRepository {
	petRepository := petRepository{dbConnection}
	return &petRepository
}

func (petRepository *petRepository) Create(pet *Pet.Pet) (error error) {
	result := petRepository.dbConn.Conn.Create(pet)
	return result.Error
}
