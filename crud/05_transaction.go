package crud

import "gorm.io/gorm"

func Transaction(db *gorm.DB) {
	user := &UserModel{Name: "John", Address: "New York"}
	tx := db.Begin()
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
}
