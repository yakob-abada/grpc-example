package repo

import "context"

type LikerRepo interface {
	ListAllLikedYou(ctx context.Context, paginatedReq *PaginatedRequest, recipientUserId string) (Paginator, error)
	ListLikedYou(ctx context.Context, statuses []int, paginatedReq *PaginatedRequest, recipientUserId string) (Paginator, error)
	CountLikedYou(ctx context.Context, recipientUserId string) (int64, error)
	Decide(ctx context.Context, recipientUserId string, actorUserId string, match bool) error
}
