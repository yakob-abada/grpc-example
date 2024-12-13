package main

import (
	"fmt"
	"github.com/yakob-abada/backend-match/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Apply drop table.
	if db.Migrator().HasTable(&model.Match{}) {
		err = db.Migrator().DropTable(&model.Match{})
		if err != nil {
			panic(err)
		}
	}
}
