package mapper

import (
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/model"
)

type LikedResponseMap struct{}

func NewLikedResponseMap() *LikedResponseMap {
	return &LikedResponseMap{}
}

func (l *LikedResponseMap) List(likes []model.Match, nextPaginationToken string) *pb.ListLikedYouResponse {
	var likers []*pb.ListLikedYouResponse_Liker
	for _, l := range likes {
		likers = append(likers, &pb.ListLikedYouResponse_Liker{
			ActorId:       l.ActorUserId,
			UnixTimestamp: uint64(l.CreatedAt.Unix()),
		})
	}

	response := &pb.ListLikedYouResponse{Likers: likers}

	if nextPaginationToken != "" {
		response.NextPaginationToken = &nextPaginationToken
	}

	return response
}

func (l *LikedResponseMap) Count(i *int64) *pb.CountLikedYouResponse {
	return &pb.CountLikedYouResponse{Count: uint64(*i)}
}

func (l *LikedResponseMap) Decision(b bool) *pb.PutDecisionResponse {
	return &pb.PutDecisionResponse{
		MutualLikes: b,
	}
}
