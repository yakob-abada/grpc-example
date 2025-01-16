package repo

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type LikeMock struct {
	mock.Mock
}

func (m *LikeMock) ListAllLikedYou(ctx context.Context, paginatedReq *PaginatedRequest, recipientUserId string) (Paginator, error) {
	args := m.Called(ctx, recipientUserId, paginatedReq)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(Paginator), args.Error(1)
}

func (m *LikeMock) ListLikedYou(ctx context.Context, statuses []int, paginatedReq *PaginatedRequest, recipientUserId string) (Paginator, error) {
	args := m.Called(ctx, recipientUserId, statuses, paginatedReq)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(Paginator), args.Error(1)
}

func (m *LikeMock) CountLikedYou(ctx context.Context, recipientUserId string) (int64, error) {
	args := m.Called(ctx, recipientUserId)
	if args.Error(1) != nil {
		return 0, args.Error(1)
	}

	return args.Get(0).(int64), args.Error(1)
}

func (m *LikeMock) Decide(ctx context.Context, recipientUserId string, actorUserId string, match bool) error {
	args := m.Called(ctx, recipientUserId, actorUserId, match)
	return args.Error(0)
}
