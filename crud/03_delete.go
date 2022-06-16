// Delete
// In order to delete the record from the table, gorm has provided Delete() as given in below examples
package crud

import "gorm.io/gorm"

func Delete(db *gorm.DB) {
	// Select records and delete it
	db.Table("user_models").Where("address= ?", "San Diego").Delete(&UserModel{})

	//Find the record and delete it
	db.Where("address=?", "Los Angeles").Delete(&UserModel{})

	// Select all records from a model and delete all
	db.Model(&UserModel{}).Delete(&UserModel{})
}
