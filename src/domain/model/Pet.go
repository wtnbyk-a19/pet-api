package model

import (
	"time"
)

type Pet struct {
	petId     int       `json:"pet_id" gorm:"primary_key"`
	Petpet    string    `json:"petpet"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewPet(petpet string) *Pet {
	pet := new(Pet)
	pet.Petpet = petpet

	// TODO: タイムゾーンの修正
	pet.CreatedAt = time.Now()
	pet.UpdatedAt = time.Now()

	return pet
}
