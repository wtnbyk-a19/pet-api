package usecase

import (
	"pet-api/src/domain/model"
	"pet-api/src/domain/repository"
)

type IPetUsecase interface {
	Execute(*model.Pet) (error error)
}

type petUsecase struct {
	petRepository repository.IPetRepository
}

func NewPetUsecase(petRepository repository.IPetRepository) IPetUsecase {
	petUsecase := petUsecase{petRepository: petRepository}
	return &petUsecase
}

func (petUsecase *petUsecase) Execute(pet *model.Pet) (error error) {
	error = petUsecase.petRepository.Persist(pet)
	return error
}
