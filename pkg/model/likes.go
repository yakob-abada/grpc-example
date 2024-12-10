package model

import "time"

type Likes struct {
	RecipientUserId string
	ActorUserId     string
	LikedBack       bool
	Matched         *bool
	CreatedAt       time.Time
}
