package config

import (
	"golang-login-with-jwt/pkg/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/golang_login_with_jwt?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helpers.CheckError(err)
	return db
}
