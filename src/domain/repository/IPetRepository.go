package repository

import "pet-api/src/domain/model"

type IPetRepository interface {
	Create(pet *model.Pet) (error error)
}
