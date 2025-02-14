package mapper

import (
	"github.com/stretchr/testify/mock"
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/model"
)

type LikedResponseMock struct {
	mock.Mock
}

func (m *LikedResponseMock) List(likes []model.Match, nextPaginationToken string) *pb.ListLikedYouResponse {
	args := m.Called(likes, nextPaginationToken)

	return args.Get(0).(*pb.ListLikedYouResponse)
}

func (m *LikedResponseMock) Count(i int64) *pb.CountLikedYouResponse {
	args := m.Called(i)
	return args.Get(0).(*pb.CountLikedYouResponse)
}

func (m *LikedResponseMock) Decision(b bool) *pb.PutDecisionResponse {
	args := m.Called(b)
	return args.Get(0).(*pb.PutDecisionResponse)
}
