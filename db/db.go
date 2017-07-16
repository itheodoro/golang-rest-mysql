package db

import (
	"github.com/itheodoro/rest-mysql/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitDB - Connect on db and create table if not exists
func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/dogs")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	if !db.HasTable(&models.Dog{}) {
		db.CreateTable(&models.Dog{})
		db.Set("gorm:tabel_options", "ENGINE=InnoDb").CreateTable(&models.Dog{})
	}

	return db
}
