package usecase

import (
	"github.com/gofiber/fiber"
	"pet-api/src/domain/model"
	"pet-api/src/domain/repository"
	"strconv"
	"time"
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
	userId        int
	petName       string
	gender        string
	variety       string
	breed         string
	birthday      time.Time
	adoptaversary time.Time
	memo          string
}

func (petUsecase *petUsecase) CreatePet(params *PetCreateParameter) (error error) {

	pet := model.NewPet(
		params.userId,
		params.petName,
		params.gender,
		params.breed,
		params.birthday,
		params.adoptaversary,
		params.memo,
	)

	error = petUsecase.petRepository.Create(pet)
	return error
}

func (params PetCreateParameter) Setup(c *fiber.Ctx) (err error) {
	var userId int
	userId, err = strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return err
	}

	var birthday time.Time
	birthday, err = time.Parse("20211107", c.Params("birthday"))
	if err != nil {
		return err
	}

	var adoptaversary time.Time
	adoptaversary, err = time.Parse("20211107", c.Params("adoptaversary"))
	if err != nil {
		return err
	}

	params.userId = userId
	params.petName = c.Params("pet_name")
	params.gender = c.Params("gender")
	params.variety = c.Params("variety")
	params.breed = c.Params("breed")
	params.birthday = birthday
	params.adoptaversary = adoptaversary
	params.memo = c.Params("memo")

	return
}
