// Auto Migration
// This Auto Migration feature will automatically migrate your schema. It will automatically create the table based on your model. We don’t need to create the table manually.

package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserModel struct {
	ID      int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)"`
}

func AutoMigration() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ormdemo?charset=utf8&parseTime=True")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	db.Debug().DropTableIfExists(&UserModel{})
	//Drops table if already exists
	db.Debug().AutoMigrate(&UserModel{})
	//Auto create table based on Model
}
