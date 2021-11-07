package model

import (
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

func NewPet(userId int, petName string, gender string, breed string, birthday time.Time, adoptaversary time.Time, memo string) *Pet {
	pet := new(Pet)

	pet.UserId = userId
	pet.PetName = petName
	pet.Gender = gender
	pet.Breed = breed
	pet.Birthday = birthday
	pet.Adoptaversary = adoptaversary
	pet.Memo = memo

	return pet
}
