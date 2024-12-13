package main

import (
	"fmt"
	"github.com/yakob-abada/backend-match/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
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

	// Apply fixtures
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Match{})

	users := []*model.Match{
		{RecipientUserId: "1", ActorUserId: "2", Status: 1, CreatedAt: time.Now()},
		{RecipientUserId: "1", ActorUserId: "3", Status: 2, CreatedAt: time.Now()},
		{RecipientUserId: "1", ActorUserId: "4", Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: "1", ActorUserId: "5", Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: "1", ActorUserId: "6", Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: "1", ActorUserId: "7", Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: "1", ActorUserId: "8", Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: "1", ActorUserId: "9", Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: "1", ActorUserId: "10", Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: "1", ActorUserId: "11", Status: 0, CreatedAt: time.Now()},
	}

	err = db.Create(users).Error
	if err != nil {
		panic(err)
	}
}
