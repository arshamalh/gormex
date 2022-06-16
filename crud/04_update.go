package crud

import "gorm.io/gorm"

// In order to update the records in the table using gorm, look into the below sample example.
func Update(db *gorm.DB) {
	user := &UserModel{Name: "John", Address: "New York"}
	// Select, edit, and save
	db.Find(&user)
	user.Address = "Brisbane"
	db.Save(&user)

	// Update with column names, not attribute names
	db.Model(&user).Update("Name", "Jack")

	db.Model(&user).Updates(
		map[string]interface{}{
			"Name":    "Amy",
			"Address": "Boston",
		})

	// UpdateColumn()
	db.Model(&user).UpdateColumn("Address", "Phoenix")
	db.Model(&user).UpdateColumns(
		map[string]interface{}{
			"Name":    "Taylor",
			"Address": "Houston",
		})

	// Using Find()
	db.Find(&user).Update("Address", "San Diego")

	// Batch Update
	db.Table("user_models").Where("address = ?", "california").Update("name", "Walker")
}
