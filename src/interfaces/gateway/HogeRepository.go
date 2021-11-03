package gateway

import (
	"../../domain/model"
	"../../domain/repository"
	"../../infrastructure/mysql"
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
