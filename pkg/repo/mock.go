package repo

import (
	"github.com/stretchr/testify/mock"
	"github.com/yakob-abada/backend-match/pkg/model"
)

type LikeMock struct {
	mock.Mock
}

func (m *LikeMock) ListLikedYou(recipientUserId string, likedBack bool) ([]*model.Likes, error) {
	args := m.Called(recipientUserId, likedBack)
	return args.Get(0).([]*model.Likes), args.Error(1)
}

func (m *LikeMock) CountLikedYou(recipientUserId string, likedBack bool) (*int, error) {
	args := m.Called(recipientUserId, likedBack)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*int), args.Error(1)
}

func (m *LikeMock) Decide(recipientUserId string, actorUserId string, match bool) error {
	args := m.Called(recipientUserId, actorUserId, match)
	return args.Error(0)
}
