package test

import (
	"github.com/stretchr/testify/assert"
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/mapper"
	"github.com/yakob-abada/backend-match/pkg/model"
	"testing"
	"time"
)

func TestList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		now := time.Now()
		likeList := []model.Match{{
			RecipientUserId: "1",
			ActorUserId:     "2",
			CreatedAt:       now,
		}}

		listLikedYouRes := &pb.ListLikedYouResponse{Likers: []*pb.ListLikedYouResponse_Liker{
			{
				ActorId:       "2",
				UnixTimestamp: uint64(now.Unix()),
			},
		}}

		sut := mapper.LikedResponseMap{}
		result := sut.List(likeList, "")

		assert.Equal(t, listLikedYouRes, result)
	})

	t.Run("success with nextPaginationToken", func(t *testing.T) {
		now := time.Now()
		nextPaginationToken := "nextPaginationToken"
		likeList := []model.Match{{
			RecipientUserId: "1",
			ActorUserId:     "2",
			CreatedAt:       now,
		}}

		listLikedYouRes := &pb.ListLikedYouResponse{Likers: []*pb.ListLikedYouResponse_Liker{
			{
				ActorId:       "2",
				UnixTimestamp: uint64(now.Unix()),
			},
		}, NextPaginationToken: &nextPaginationToken}

		sut := mapper.LikedResponseMap{}
		result := sut.List(likeList, nextPaginationToken)

		assert.Equal(t, listLikedYouRes, result)
	})
}

func TestCount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		count := int64(2)
		resp := &pb.CountLikedYouResponse{Count: 2}
		sut := mapper.LikedResponseMap{}
		result := sut.Count(&count)

		assert.Equal(t, resp, result)
	})
}

func TestDecide(t *testing.T) {
	t.Run("success with match", func(t *testing.T) {
		resp := &pb.PutDecisionResponse{MutualLikes: true}

		sut := mapper.LikedResponseMap{}
		result := sut.Decision(true)
		assert.Equal(t, resp, result)
	})

	t.Run("success with unmatch", func(t *testing.T) {
		resp := &pb.PutDecisionResponse{MutualLikes: false}

		sut := mapper.LikedResponseMap{}
		result := sut.Decision(false)
		assert.Equal(t, resp, result)
	})
}
