package model

import "time"

type Hoge struct {
	hogeId    int        `json:"hoge_id"gorm:"primary_key"`
	Hogehoge  string     `json:"hogehoge"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index"json:"-"`
}

func NewHoge(hogehoge string) *Hoge {
	hoge := new(Hoge)
	hoge.Hogehoge = hogehoge

	return hoge
}
