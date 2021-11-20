package Pet

import (
	"github.com/gofiber/fiber/v2"
	"pet-api/src/domain/Pet"
)

type IPetUsecase interface {
	CreatePet(params *PetCreateParameter) (error error)
}

type petUsecase struct {
	petRepository Pet.IPetRepository
}

func NewPetUsecase(petRepository Pet.IPetRepository) IPetUsecase {
	petUsecase := petUsecase{petRepository: petRepository}
	return &petUsecase
}

type PetCreateParameter struct {
	UserId        string `json:"user_id"`
	PetName       string `json:"pet_name"`
	Gender        string `json:"gender"`
	Variety       string `json:"variety"`
	Breed         string `json:"breed"`
	Birthday      string `json:"birthday"`
	Adoptaversary string `json:"adoptaversary"`
	Memo          string `json:"memo"`
}

func (petUsecase *petUsecase) CreatePet(params *PetCreateParameter) (err error) {
	var pet *Pet.Pet
	pet, err = Pet.NewPet(
		params.UserId,
		params.PetName,
		params.Gender,
		params.Breed,
		params.Birthday,
		params.Adoptaversary,
		params.Memo,
	)

	err = petUsecase.petRepository.Create(pet)
	return err
}

func (params *PetCreateParameter) ParamsSetup(c *fiber.Ctx) (err error) {
	err = c.BodyParser(params)
	if err != nil {
		return err
	}

	return
}
