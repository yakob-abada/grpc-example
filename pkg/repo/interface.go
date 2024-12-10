package repo

import "github.com/yakob-abada/backend-match/pkg/model"

type LikerRepo interface {
	ListLikedYou(recipientUserId string, likedBack bool) ([]*model.Likes, error)
	CountLikedYou(recipientUserId string, likedBack bool) (*int, error)
	Decide(recipientUserId string, actorUserId string, match bool) error
}
