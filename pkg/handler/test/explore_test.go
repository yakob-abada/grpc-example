package test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/handler"
	"github.com/yakob-abada/backend-match/pkg/mapper"
	"github.com/yakob-abada/backend-match/pkg/model"
	"github.com/yakob-abada/backend-match/pkg/repo"
	"testing"
	"time"
)

func TestListLikedYou(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		now := time.Now()
		matched := true
		likeList := []*model.Likes{{
			RecipientUserId: "1",
			ActorUserId:     "2",
			LikedBack:       false,
			Matched:         &matched,
			CreatedAt:       now,
		}}
		repoMock := &repo.LikeMock{}
		repoMock.On("ListLikedYou", "1", true).Once().Return(likeList, nil)

		listLikedYouRes := &pb.ListLikedYouResponse{Likers: []*pb.ListLikedYouResponse_Liker{
			{
				ActorId:       "2",
				UnixTimestamp: uint64(now.Unix()),
			},
		}}
		mapperMock := &mapper.LikedResponseMock{}
		mapperMock.On("List", likeList).Once().Return(listLikedYouRes, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock)
		response, err := sut.ListLikedYou(context.Background(), &pb.ListLikedYouRequest{RecipientUserId: "1"})

		assert.Equal(t, listLikedYouRes, response)
		assert.Nil(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		repoMock := &repo.LikeMock{}
		repoMock.On("ListLikedYou", "1", true).Once().Return([]*model.Likes{}, errors.New("connection failed"))

		mapperMock := &mapper.LikedResponseMock{}

		sut := handler.NewExploreServer(repoMock, mapperMock)
		response, err := sut.ListLikedYou(context.Background(), &pb.ListLikedYouRequest{RecipientUserId: "1"})

		assert.Nil(t, response)
		assert.EqualError(t, err, "connection failed")
	})
}

func TestListNewLikedYou(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		now := time.Now()
		matched := true
		likeList := []*model.Likes{{
			RecipientUserId: "1",
			ActorUserId:     "2",
			LikedBack:       false,
			Matched:         &matched,
			CreatedAt:       now,
		}}
		repoMock := &repo.LikeMock{}
		repoMock.On("ListLikedYou", "1", false).Once().Return(likeList, nil)

		listLikedYouRes := &pb.ListLikedYouResponse{Likers: []*pb.ListLikedYouResponse_Liker{
			{
				ActorId:       "2",
				UnixTimestamp: uint64(now.Unix()),
			},
		}}
		mapperMock := &mapper.LikedResponseMock{}
		mapperMock.On("List", likeList).Once().Return(listLikedYouRes, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock)
		response, err := sut.ListNewLikedYou(context.Background(), &pb.ListLikedYouRequest{RecipientUserId: "1"})

		assert.Equal(t, listLikedYouRes, response)
		assert.Nil(t, err)
	})
	t.Run("failure", func(t *testing.T) {
		repoMock := &repo.LikeMock{}
		repoMock.On("ListLikedYou", "1", false).Once().Return([]*model.Likes{}, errors.New("connection failed"))

		mapperMock := &mapper.LikedResponseMock{}

		sut := handler.NewExploreServer(repoMock, mapperMock)
		response, err := sut.ListNewLikedYou(context.Background(), &pb.ListLikedYouRequest{RecipientUserId: "1"})

		assert.Nil(t, response)
		assert.EqualError(t, err, "connection failed")
	})
}

func TestCountLikedYou(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repoMock := &repo.LikeMock{}
		count := 2
		repoMock.On("CountLikedYou", "1", false).Once().Return(&count, nil)

		resp := &pb.CountLikedYouResponse{Count: 2}
		mapperMock := &mapper.LikedResponseMock{}
		mapperMock.On("Count", &count).Once().Return(resp, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock)
		response, err := sut.CountLikedYou(context.Background(), &pb.CountLikedYouRequest{RecipientUserId: "1"})
		assert.Equal(t, resp, response)
		assert.Nil(t, err)
	})
	t.Run("failure", func(t *testing.T) {
		repoMock := &repo.LikeMock{}
		repoMock.On("CountLikedYou", "1", false).Once().Return(nil, errors.New("connection failed"))

		mapperMock := &mapper.LikedResponseMock{}

		sut := handler.NewExploreServer(repoMock, mapperMock)
		response, err := sut.CountLikedYou(context.Background(), &pb.CountLikedYouRequest{RecipientUserId: "1"})
		assert.Nil(t, response)
		assert.EqualError(t, err, "connection failed")
	})
}

func TestPutDecision(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repoMock := &repo.LikeMock{}
		repoMock.On("Decide", "1", "2", false).Once().Return(nil)

		resp := &pb.PutDecisionResponse{MutualLikes: false}
		mapperMock := &mapper.LikedResponseMock{}
		mapperMock.On("Decision", false).Once().Return(resp, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock)
		response, err := sut.PutDecision(context.Background(), &pb.PutDecisionRequest{RecipientUserId: "1", ActorUserId: "2", LikedRecipient: false})
		assert.Equal(t, resp, response)
		assert.Nil(t, err)
	})
	t.Run("failure", func(t *testing.T) {
		repoMock := &repo.LikeMock{}
		repoMock.On("Decide", "1", "2", false).Once().Return(errors.New("connection failed"))

		mapperMock := &mapper.LikedResponseMock{}

		sut := handler.NewExploreServer(repoMock, mapperMock)
		response, err := sut.PutDecision(context.Background(), &pb.PutDecisionRequest{RecipientUserId: "1", ActorUserId: "2", LikedRecipient: false})
		assert.Nil(t, response)
		assert.EqualError(t, err, "connection failed")
	})
}
