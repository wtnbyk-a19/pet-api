package model

import (
	"time"
)

type Pet struct {
	PetId     int       `json:"pet_id" gorm:"primary_key"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewPet(name string) *Pet {
	pet := new(Pet)
	pet.Name = name

	// TODO: タイムゾーンの修正
	pet.CreatedAt = time.Now()
	pet.UpdatedAt = time.Now()

	return pet
}
