package repo

type LikerRepo interface {
	ListAllLikedYou(recipientUserId string, paginatedReq *PaginatedRequest) (Paginator, error)
	ListLikedYou(recipientUserId string, statuses []int, paginatedReq *PaginatedRequest) (Paginator, error)
	CountLikedYou(recipientUserId string) (int64, error)
	Decide(recipientUserId string, actorUserId string, match bool) error
}
