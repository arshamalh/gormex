// Auto Migration
// This Auto Migration feature will automatically migrate your schema. It will automatically create the table based on your model. We don’t need to create the table manually.

package main

import (
	"log"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserModel struct {
	ID      int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)"`
}

func AutoMigration() {
	dsn := "host=localhost user=postgres password=pg2local dbname=simple_blog port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	db.Debug().AutoMigrate(&UserModel{})
	//Auto create table based on Model
}
