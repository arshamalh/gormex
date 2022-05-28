package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserR struct { // R is reduntant, it's here just to prevent naming conflicts
	gorm.Model
	Birthday time.Time
	Age      int64
	Name     string `sql:"size:255"`

	Emails            []Email       // One-To-Many relationship (has many)
	BillingAddress    Address       // One-To-One relationship (has one)
	BillingAddressId  sql.NullInt64 // Foreign key of BillingAddress
	ShippingAddress   Address       // One-To-One relationship (has one)
	ShippingAddressId int64         // Foreign key of ShippingAddress
	IgnoreMe          int64         `sql:"-"`                          // Ignore this field
	Languages         []LanguageR   `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
}

type Email struct {
	Id         int64
	UserId     int64  // Foreign key for UserR (belongs to)
	Email      string `sql:"type:varchar(100);"` // Set field's type
	Subscribed bool
}

type Address struct {
	ID       int64
	Address1 string         `sql:"not null;unique"` // Set field as not nullable and unique
	Address2 string         `sql:"type:varchar(100)"`
	Post     sql.NullString `sql:"not null"`
}

type LanguageR struct {
	ID   int64
	Name string
}

func Relations() {
	err := os.Remove("/tmp/sqlite.db")
	if err != nil {
		fmt.Println(err)
	}
	dsn := "host=localhost user=postgres password=pg2local dbname=simple_blog port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println(err)
	}

	db.Exec("PRAGMA foreign_keys = ON;")

	db.AutoMigrate(UserR{}, Email{}, Address{}, LanguageR{})

	user := UserR{
		Name:            "jinzhu",
		BillingAddress:  Address{Address1: "Billing Address - Address 1"},
		ShippingAddress: Address{Address1: "Shipping Address - Address 1"},
		Emails:          []Email{{Email: "jinzhu@example.com"}, {Email: "jinzhu-2@example@example.com"}},
		Languages:       []LanguageR{{Name: "ZH"}, {Name: "EN"}},
	}

	db.Create(&user)

	var u UserR
	// db.Debug().First(&u).Related(&u.Emails)
	fmt.Println(u.Emails)
}
