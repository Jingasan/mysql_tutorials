package model

import (
	"gorm.io/gorm"
)

type BlogEntity struct {
	gorm.Model
	Title string
	Body  string
}

func CreateTable() {
	Db.Migrator().CreateTable(&BlogEntity{})
}

func DeleteTable() {
	Db.Migrator().DropTable(&BlogEntity{})
}

func GetAll() (datas []BlogEntity) {
	result := Db.Find(&datas)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOne(id int) (data BlogEntity) {
	result := Db.First(&data, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func (b *BlogEntity) Create() {
	result := Db.Create(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *BlogEntity) Update() {
	result := Db.Save(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *BlogEntity) Delete() {
	result := Db.Delete(b)
	if result.Error != nil {
		panic(result.Error)
	}
}
