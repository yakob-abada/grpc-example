package test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/handler"
	"github.com/yakob-abada/backend-match/pkg/mapper"
	"github.com/yakob-abada/backend-match/pkg/model"
	"github.com/yakob-abada/backend-match/pkg/pagination"
	"github.com/yakob-abada/backend-match/pkg/repo"
	"testing"
	"time"
)

func TestListLikedYou(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		now := time.Now()
		likeList := []model.Match{{
			RecipientUserId: "1",
			ActorUserId:     "2",
			Status:          repo.MatchStatusMatched,
			CreatedAt:       now,
		}}
		pr := repo.NewPaginatedResult(likeList, false)
		repoMock := &repo.LikeMock{}
		repoMock.On("ListLikedYou", "1", repo.MatchStatusMatched, repo.NewPaginatedRequest(10, 10)).Once().Return(pr, nil)

		listLikedYouRes := &pb.ListLikedYouResponse{Likers: []*pb.ListLikedYouResponse_Liker{
			{
				ActorId:       "2",
				UnixTimestamp: uint64(now.Unix()),
			},
		}}
		mapperMock := &mapper.LikedResponseMock{}
		mapperMock.On("List", likeList, "").Once().Return(listLikedYouRes, nil)

		token := pagination.Token{
			Offset:          10,
			RequestChecksum: 10,
			PageSize:        10,
		}
		pageTokenMock := &pagination.PageTokenMock{}
		pageTokenMock.On("Parse", mock.Anything).Once().Return(token, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock, pageTokenMock)
		response, err := sut.ListLikedYou(context.Background(), &pb.ListLikedYouRequest{RecipientUserId: "1"})

		assert.Equal(t, listLikedYouRes, response)
		assert.Nil(t, err)
	})

	t.Run("success with next page", func(t *testing.T) {
		now := time.Now()
		likeList := []model.Match{{
			RecipientUserId: "1",
			ActorUserId:     "2",
			Status:          repo.MatchStatusMatched,
			CreatedAt:       now,
		}}
		pr := repo.NewPaginatedResult(likeList, true)
		repoMock := &repo.LikeMock{}
		repoMock.On("ListLikedYou", "1", repo.MatchStatusMatched, repo.NewPaginatedRequest(10, 10)).Once().Return(pr, nil)

		listLikedYouRes := &pb.ListLikedYouResponse{Likers: []*pb.ListLikedYouResponse_Liker{
			{
				ActorId:       "2",
				UnixTimestamp: uint64(now.Unix()),
			},
		}}
		mapperMock := &mapper.LikedResponseMock{}
		mapperMock.On("List", likeList, mock.AnythingOfType("string")).Once().Return(listLikedYouRes, nil)

		token := pagination.Token{
			Offset:          10,
			RequestChecksum: 10,
			PageSize:        10,
		}
		pageTokenMock := &pagination.PageTokenMock{}
		pageTokenMock.On("Parse", mock.Anything).Once().Return(token, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock, pageTokenMock)
		response, err := sut.ListLikedYou(context.Background(), &pb.ListLikedYouRequest{RecipientUserId: "1"})

		assert.Equal(t, listLikedYouRes, response)
		assert.Nil(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		repoMock := &repo.LikeMock{}
		repoMock.On("ListLikedYou", "1", repo.MatchStatusMatched, repo.NewPaginatedRequest(10, 10)).Once().Return(nil, errors.New("connection failed"))

		mapperMock := &mapper.LikedResponseMock{}

		token := pagination.Token{
			Offset:          10,
			RequestChecksum: 10,
			PageSize:        10,
		}
		pageTokenMock := &pagination.PageTokenMock{}
		pageTokenMock.On("Parse", mock.Anything).Once().Return(token, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock, pageTokenMock)
		response, err := sut.ListLikedYou(context.Background(), &pb.ListLikedYouRequest{RecipientUserId: "1"})

		assert.Nil(t, response)
		assert.EqualError(t, err, "rpc error: code = Internal desc = failed to get result: connection failed")
	})
}

func TestListNewLikedYou(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		now := time.Now()
		likeList := []model.Match{{
			RecipientUserId: "1",
			ActorUserId:     "2",
			Status:          repo.MatchStatusPending,
			CreatedAt:       now,
		}}
		pr := repo.NewPaginatedResult(likeList, false)
		repoMock := &repo.LikeMock{}
		repoMock.On("ListLikedYou", "1", repo.MatchStatusPending, repo.NewPaginatedRequest(10, 10)).Once().Return(pr, nil)

		listLikedYouRes := &pb.ListLikedYouResponse{Likers: []*pb.ListLikedYouResponse_Liker{
			{
				ActorId:       "2",
				UnixTimestamp: uint64(now.Unix()),
			},
		}}
		mapperMock := &mapper.LikedResponseMock{}
		mapperMock.On("List", likeList, "").Once().Return(listLikedYouRes, nil)

		token := pagination.Token{
			Offset:          10,
			RequestChecksum: 10,
			PageSize:        10,
		}
		pageTokenMock := &pagination.PageTokenMock{}
		pageTokenMock.On("Parse", mock.Anything).Once().Return(token, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock, pageTokenMock)
		response, err := sut.ListNewLikedYou(context.Background(), &pb.ListLikedYouRequest{RecipientUserId: "1"})

		assert.Equal(t, listLikedYouRes, response)
		assert.Nil(t, err)
	})

	t.Run("success with next page", func(t *testing.T) {
		now := time.Now()
		likeList := []model.Match{{
			RecipientUserId: "1",
			ActorUserId:     "2",
			Status:          repo.MatchStatusPending,
			CreatedAt:       now,
		}}
		pr := repo.NewPaginatedResult(likeList, true)
		repoMock := &repo.LikeMock{}
		repoMock.On("ListLikedYou", "1", repo.MatchStatusPending, repo.NewPaginatedRequest(10, 10)).Once().Return(pr, nil)

		listLikedYouRes := &pb.ListLikedYouResponse{Likers: []*pb.ListLikedYouResponse_Liker{
			{
				ActorId:       "2",
				UnixTimestamp: uint64(now.Unix()),
			},
		}}
		mapperMock := &mapper.LikedResponseMock{}
		mapperMock.On("List", likeList, mock.AnythingOfType("string")).Once().Return(listLikedYouRes, nil)

		token := pagination.Token{
			Offset:          10,
			RequestChecksum: 10,
			PageSize:        10,
		}
		pageTokenMock := &pagination.PageTokenMock{}
		pageTokenMock.On("Parse", mock.Anything).Once().Return(token, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock, pageTokenMock)
		response, err := sut.ListNewLikedYou(context.Background(), &pb.ListLikedYouRequest{RecipientUserId: "1"})

		assert.Equal(t, listLikedYouRes, response)
		assert.Nil(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		repoMock := &repo.LikeMock{}
		repoMock.On("ListLikedYou", "1", repo.MatchStatusPending, repo.NewPaginatedRequest(10, 10)).Once().Return(nil, errors.New("connection failed"))

		mapperMock := &mapper.LikedResponseMock{}

		token := pagination.Token{
			Offset:          10,
			RequestChecksum: 10,
			PageSize:        10,
		}
		pageTokenMock := &pagination.PageTokenMock{}
		pageTokenMock.On("Parse", mock.Anything).Once().Return(token, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock, pageTokenMock)
		response, err := sut.ListNewLikedYou(context.Background(), &pb.ListLikedYouRequest{RecipientUserId: "1"})

		assert.Nil(t, response)
		assert.EqualError(t, err, "rpc error: code = Internal desc = failed to get result: connection failed")
	})
}

func TestCountLikedYou(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repoMock := &repo.LikeMock{}
		count := int64(2)
		repoMock.On("CountLikedYou", "1", repo.MatchStatusMatched).Once().Return(count, nil)

		resp := &pb.CountLikedYouResponse{Count: 2}
		mapperMock := &mapper.LikedResponseMock{}
		mapperMock.On("Count", count).Once().Return(resp, nil)

		token := pagination.Token{
			Offset:          10,
			RequestChecksum: 10,
			PageSize:        10,
		}
		pageTokenMock := &pagination.PageTokenMock{}
		pageTokenMock.On("Parse", mock.Anything).Once().Return(token, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock, pageTokenMock)
		response, err := sut.CountLikedYou(context.Background(), &pb.CountLikedYouRequest{RecipientUserId: "1"})
		assert.Equal(t, resp, response)
		assert.Nil(t, err)
	})
	t.Run("failure", func(t *testing.T) {
		repoMock := &repo.LikeMock{}
		repoMock.On("CountLikedYou", "1", repo.MatchStatusMatched).Once().Return(nil, errors.New("connection failed"))

		mapperMock := &mapper.LikedResponseMock{}

		token := pagination.Token{
			Offset:          10,
			RequestChecksum: 10,
			PageSize:        10,
		}
		pageTokenMock := &pagination.PageTokenMock{}
		pageTokenMock.On("Parse", mock.Anything).Once().Return(token, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock, pageTokenMock)
		response, err := sut.CountLikedYou(context.Background(), &pb.CountLikedYouRequest{RecipientUserId: "1"})
		assert.Nil(t, response)
		assert.EqualError(t, err, "rpc error: code = Internal desc = failed to get result: connection failed")
	})
}

func TestPutDecision(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repoMock := &repo.LikeMock{}
		repoMock.On("Decide", "1", "2", false).Once().Return(nil)

		resp := &pb.PutDecisionResponse{MutualLikes: false}
		mapperMock := &mapper.LikedResponseMock{}
		mapperMock.On("Decision", false).Once().Return(resp, nil)

		token := pagination.Token{
			Offset:          10,
			RequestChecksum: 10,
			PageSize:        10,
		}
		pageTokenMock := &pagination.PageTokenMock{}
		pageTokenMock.On("Parse", mock.Anything).Once().Return(token, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock, pageTokenMock)
		response, err := sut.PutDecision(context.Background(), &pb.PutDecisionRequest{RecipientUserId: "1", ActorUserId: "2", LikedRecipient: false})
		assert.Equal(t, resp, response)
		assert.Nil(t, err)
	})
	t.Run("failure", func(t *testing.T) {
		repoMock := &repo.LikeMock{}
		repoMock.On("Decide", "1", "2", false).Once().Return(errors.New("connection failed"))

		mapperMock := &mapper.LikedResponseMock{}

		token := pagination.Token{
			Offset:          10,
			RequestChecksum: 10,
			PageSize:        10,
		}
		pageTokenMock := &pagination.PageTokenMock{}
		pageTokenMock.On("Parse", mock.Anything).Once().Return(token, nil)

		sut := handler.NewExploreServer(repoMock, mapperMock, pageTokenMock)
		response, err := sut.PutDecision(context.Background(), &pb.PutDecisionRequest{RecipientUserId: "1", ActorUserId: "2", LikedRecipient: false})
		assert.Nil(t, response)
		assert.EqualError(t, err, "rpc error: code = Internal desc = failed to update: connection failed")
	})
}
