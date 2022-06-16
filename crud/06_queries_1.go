// Queries
// In order to fetch the records from the database and do some SQL stuffs gorm has given some query functions. We’ll now do a quick discussion on it.
package crud

import "gorm.io/gorm"

func Queries_1(db *gorm.DB) {
	user := &UserModel{}
	users := []UserModel{}
	// Get first record, order by primary key
	db.First(&user)

	// Get last record, order by primary key
	db.Last(&user)

	// Get all records
	db.Find(&users)

	// Get record with primary key (only works for integer primary key)
	db.First(&user, 10)

	// Query with Where() [some SQL functions]

	db.Where("address = ?", "Los Angeles").First(&user)
	//SELECT * FROM user_models WHERE address=’Los Angeles’ limit 1;

	db.Where("address = ?", "Los Angeles").Find(&user)
	//SELECT * FROM user_models WHERE address=’Los Angeles’;

	db.Where("address <> ?", "New York").Find(&user)
	//SELECT * FROM user_models WHERE address<>’Los Angeles’;

	// IN
	db.Where("name in (?)", []string{"John", "Martin"}).Find(&user)

	// LIKE
	db.Where("name LIKE ?", "%ti%").Find(&user)

	// AND
	db.Where("name = ? AND address >= ?", "Martin", "Los Angeles").Find(&user)

}
