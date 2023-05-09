package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"prakerja3/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja_3?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Init database failed")
	}
	migratDatabase()
}

func migratDatabase() {
	DB.AutoMigrate(&models.User{})
}
