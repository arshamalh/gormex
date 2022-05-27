// Creating Models And Tables
// Define the models before creating tables and based on the model the table will be created.
package main

type User struct {
	ID       int
	Username string
}

func CreatingModelsAndTables() {
	// After db connection is created.
	db.CreateTable(&User{})

	// Also some useful functions
	db.HasTable(&User{})          // =>;; true
	db.DropTableIfExists(&User{}) //Drops the table if already exists
}
