package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "root:abc123@tcp(192.168.131.189:3306)/study_go?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB.Callback()
}

func ErrHandler(db *gorm.DB, err error) error {
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return err
	}
	return nil
}
