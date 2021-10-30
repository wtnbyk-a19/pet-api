package repository

import "docker-go-api/src/domain/model"

type IHogeRepository interface {
	Persist(hoge *model.Hoge) (error error)
}