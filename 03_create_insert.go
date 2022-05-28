// Create/Insert
// In order to create or insert a record,
// you need to use the Create() function.
// The save() is also there that will return the primary key of the inserted record.

package main

func Create() {
	user := &UserModel{Name: "John", Address: "New York"}
	db.Create(user)

	// Internally it will create the query like
	// INSERT INTO `user_models` (`name`,`address`) VALUES ('John','New York')

	//You can insert multiple records too
	var users []UserModel = []UserModel{
		{Name: "Ricky", Address: "Sydney"},
		{Name: "Adam", Address: "Brisbane"},
		{Name: "Justin", Address: "California"},
	}

	for _, user := range users {
		db.Create(&user)
	}
}
