package main

import (
	"github.com/yakob-abada/backend-match/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dating_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Apply migration
	err = db.AutoMigrate(&model.Match{})
	if err != nil {
		panic(err)
	}
}
