package repo

import (
	"github.com/stretchr/testify/mock"
)

type LikeMock struct {
	mock.Mock
}

func (m *LikeMock) ListLikedYou(recipientUserId string, status int, paginatedReq *PaginatedRequest) (Paginator, error) {
	args := m.Called(recipientUserId, status, paginatedReq)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(Paginator), args.Error(1)
}

func (m *LikeMock) CountLikedYou(recipientUserId string, status int) (int64, error) {
	args := m.Called(recipientUserId, status)
	if args.Error(1) != nil {
		return 0, args.Error(1)
	}

	return args.Get(0).(int64), args.Error(1)
}

func (m *LikeMock) Decide(recipientUserId string, actorUserId string, match bool) error {
	args := m.Called(recipientUserId, actorUserId, match)
	return args.Error(0)
}
