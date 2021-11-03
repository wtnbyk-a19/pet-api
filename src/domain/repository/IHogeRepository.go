package repository

import "../../domain/model"

type IHogeRepository interface {
	Persist(hoge *model.Hoge) (error error)
}
