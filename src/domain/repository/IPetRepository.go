package repository

import "pet-api/src/domain/model"

type IPetRepository interface {
	Save(pet *model.Pet) (error error)
}
