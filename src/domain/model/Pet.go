package model

import (
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

	Model
}

func NewPet(userIdStr string, petName string, gender string, breed string, birthdayStr string, adoptaversaryStr string, memo string) (pet *Pet, err error) {
	pet = &Pet{}

	var userId int
	userId, err = strconv.Atoi(userIdStr)
	if err != nil {
		return pet, err
	}

	var birthday time.Time
	birthday, err = time.Parse("20211107", birthdayStr)
	if err != nil {
		return pet, err
	}

	var adoptaversary time.Time
	adoptaversary, err = time.Parse("20211107", adoptaversaryStr)
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
