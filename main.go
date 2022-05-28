package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	dsn := "host=localhost user=postgres password=pg2local dbname=simple_blog port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	fmt.Printf("Connection Established: %+v\n", db)
}
