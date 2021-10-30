package model

import "time"

type Hoge struct {
	hogeId    int        `json:"hoge_id"gorm:"primary_key"`
	hogehoge  string     `json:"hogehoge"`
	createdAt time.Time  `json:"created_at"`
	updatedAt time.Time  `json:"update_at"`
	deletedAt *time.Time `sql:"index"json:"-"`
}

func NewHoge(hogehoge string) *Hoge {
	hoge := new(Hoge)
	hoge.hogehoge = hogehoge

	return hoge
}

func (hoge Hoge) getHogehoge() string {
	return hoge.hogehoge
}

func (hoge *Hoge) setHogehoge(hogehoge string) {
	hoge.hogehoge = hogehoge
}
