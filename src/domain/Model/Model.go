package Model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Model struct {
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at"`
}
