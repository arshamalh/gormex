// One to One Relationship specifies how the fields of one models are related to others by specifying one to one mapping. For now, I’ve considered and done one to one mapping between Place and Town struct/model. Here one Town belongs to one Place relational mapping I’ve created.

package main

import (
	"fmt"
)

type Place struct {
	ID     int `gorm:"primary_key"`
	Name   string
	Town   Town
	TownId int `gorm:"ForeignKey:id"` // it seems this foreignKey tag doesn't work!
}

type Town struct {
	ID   int `gorm:"primary_key"`
	Name string
}

func OneToOne() {
	db.AutoMigrate(&Place{}, &Town{})

	// We need to add foreign keys manually.
	// db.Model(&Place{}).AddForeignKey("town_id", "towns(id)", "CASCADE", "CASCADE")

	t1 := Town{
		Name: "Pune",
	}
	t2 := Town{
		Name: "Mumbai",
	}
	t3 := Town{
		Name: "Hyderabad",
	}

	p1 := Place{
		Name: "Katraj",
		Town: t1,
	}
	p2 := Place{
		Name: "Thane",
		Town: t2,
	}
	p3 := Place{
		Name: "Secundarabad",
		Town: t3,
	}

	db.Save(&p1) //Saving one to one relationship
	db.Save(&p2)
	db.Save(&p3)

	fmt.Println("t1==>", t1, "p1==>", p1)
	fmt.Println("t2==>", t2, "p2s==>", p2)
	fmt.Println("t2==>", t3, "p2s==>", p3)

	//Delete
	db.Where("name=?", "Hyderabad").Delete(&Town{})

	//Update
	db.Model(&Place{}).Where("id=?", 1).Update("name", "Shivaji Nagar")

	//Select
	places := Place{}
	towns := Town{}
	fmt.Println("Before Association", places)
	db.Where("name=?", "Shivaji Nagar").Find(&places)
	fmt.Println("After Association", places)
	err := db.Model(&places).Association("town").Find(&places.Town).Error()
	fmt.Println("After Association", towns, places)
	fmt.Println("After Association", towns, places, err)
}
