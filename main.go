package main

import (
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ormdemo?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	defer db.Close()
}
