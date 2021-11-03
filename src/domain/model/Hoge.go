package model

import (
	"time"
)

type Hoge struct {
	hogeId    int       `json:"hoge_id" gorm:"primary_key"`
	Hogehoge  string    `json:"hogehoge"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewHoge(hogehoge string) *Hoge {
	hoge := new(Hoge)
	hoge.Hogehoge = hogehoge

	// TODO: タイムゾーンの修正
	hoge.CreatedAt = time.Now()
	hoge.UpdatedAt = time.Now()

	return hoge
}
