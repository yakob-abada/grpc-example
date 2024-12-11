package model

import (
	"time"
)

type Match struct {
	RecipientUserId string `gorm:"primaryKey;autoIncrement:false"`
	ActorUserId     string `gorm:"primaryKey;autoIncrement:false"`
	Status          int8
	CreatedAt       time.Time
}
