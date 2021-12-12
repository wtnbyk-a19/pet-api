package Pet

import (
	"github.com/sirupsen/logrus"
	"pet-api/src/domain/Model"
	"strconv"
	"time"
)

type Pet struct {
	PetId         int       `json:"pet_id" gorm:"primary_key"`
	UserId        int       `json:"user_id"`
	PetName       string    `json:"pet_name"`
	Gender        string    `json:"gender"`
	Breed         string    `json:"breed"`
	Birthday      time.Time `json:"birthday"`
	Adoptaversary time.Time `json:"adoptaversary"`
	Memo          string    `json:"memo"`

	Model.Model
}

func NewPet(userIdStr string, petName string, gender string, breed string, birthdayStr string, adoptaversaryStr string, memo string) (pet *Pet, err error) {
	pet = &Pet{}

	var userId int
	userId, err = strconv.Atoi(userIdStr)
	if err != nil {
		return pet, err
	}

	var birthday time.Time
	birthday, err = time.Parse("2006-01-02", birthdayStr)
	logrus.Println(birthdayStr)
	logrus.Println(birthday)
	logrus.Error("birthday error:", err)
	if err != nil {
		return pet, err
	}

	var adoptaversary time.Time
	adoptaversary, err = time.Parse("2006-01-02", adoptaversaryStr)
	logrus.Println(adoptaversaryStr)
	logrus.Println(adoptaversary)
	logrus.Error("adoptaversary error:", err)
	if err != nil {
		return pet, err
	}

	pet.UserId = userId
	pet.PetName = petName
	pet.Gender = gender
	pet.Breed = breed
	pet.Birthday = birthday
	pet.Adoptaversary = adoptaversary
	pet.Memo = memo

	return pet, nil
}
