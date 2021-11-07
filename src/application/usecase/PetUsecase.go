package usecase

import (
	"github.com/gofiber/fiber/v2"
	"pet-api/src/domain/model"
	"pet-api/src/domain/repository"
)

type IPetUsecase interface {
	CreatePet(params *PetCreateParameter) (error error)
}

type petUsecase struct {
	petRepository repository.IPetRepository
}

func NewPetUsecase(petRepository repository.IPetRepository) IPetUsecase {
	petUsecase := petUsecase{petRepository: petRepository}
	return &petUsecase
}

type PetCreateParameter struct {
	userId        string `json:"user_id"`
	petName       string `json:"pet_name"`
	gender        string `json:"gender"`
	variety       string `json:"variety"`
	breed         string `json:"breed"`
	birthday      string `json:"birthday"`
	adoptaversary string `json:"adoptaversary"`
	memo          string `json:"memo"`
}

func (petUsecase *petUsecase) CreatePet(params *PetCreateParameter) (err error) {
	var pet *model.Pet
	pet, err = model.NewPet(
		params.userId,
		params.petName,
		params.gender,
		params.breed,
		params.birthday,
		params.adoptaversary,
		params.memo,
	)

	err = petUsecase.petRepository.Create(pet)
	return err
}

func (params PetCreateParameter) ParamsSetup(c *fiber.Ctx) (err error) {
	err = c.BodyParser(&params)
	if err != nil {
		return err
	}

	return
}
