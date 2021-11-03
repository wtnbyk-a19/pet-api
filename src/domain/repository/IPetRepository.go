package repository

import "pet-api/src/domain/model"

type IPetRepository interface {
	Persist(pet *model.Pet) (error error)
}
