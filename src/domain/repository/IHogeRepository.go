package repository

import "pet-api/src/domain/model"

type IHogeRepository interface {
	Persist(hoge *model.Hoge) (error error)
}
