package repo

type LikerRepo interface {
	ListLikedYou(recipientUserId string, status int) (Paginator, error)
	CountLikedYou(recipientUserId string, status int) (*int64, error)
	Decide(recipientUserId string, actorUserId string, match bool) error
}
