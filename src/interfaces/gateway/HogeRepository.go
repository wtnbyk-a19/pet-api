package gateway

import (
	"pet-api/src/domain/model"
	"pet-api/src/domain/repository"
	"pet-api/src/infrastructure/mysql"
)

type HogeRepository struct {
	dbConnection mysql.DbConnection
}

func NewHogeRepository(dbConnection mysql.DbConnection) repository.IHogeRepository {
	hogeRepository := HogeRepository{dbConnection}
	return &hogeRepository
}

func (hogeRepository *HogeRepository) Persist(hoge *model.Hoge) (error error) {
	result := hogeRepository.dbConnection.Connection.Create(hoge)
	return result.Error
}
