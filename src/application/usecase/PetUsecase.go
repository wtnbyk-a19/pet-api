package usecase

import (
	"pet-api/src/domain/model"
	"pet-api/src/domain/repository"
)

type IPetUsecase interface {
	SavePet(*model.Pet) (error error)
}

type petUsecase struct {
	petRepository repository.IPetRepository
}

func NewPetUsecase(petRepository repository.IPetRepository) IPetUsecase {
	petUsecase := petUsecase{petRepository: petRepository}
	return &petUsecase
}

func (petUsecase *petUsecase) SavePet(pet *model.Pet) (error error) {
	error = petUsecase.petRepository.Save(pet)
	return error
}
