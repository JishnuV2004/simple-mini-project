package config

import (
	models "gorm/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



var DB *gorm.DB

func InitDB(){
	dsn := "root:jishnu2004@tcp(127.0.0.1:3306)/movion?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("database connection failed")
	}
	DB.AutoMigrate(&models.User{})
}