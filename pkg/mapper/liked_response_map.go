package mapper

import (
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/model"
)

type LikedResponseMap struct{}

func (l *LikedResponseMap) List(likes []*model.Likes) *pb.ListLikedYouResponse {
	var likers []*pb.ListLikedYouResponse_Liker
	for _, l := range likes {
		likers = append(likers, &pb.ListLikedYouResponse_Liker{
			ActorId:       l.ActorUserId,
			UnixTimestamp: uint64(l.CreatedAt.Unix()),
		})
	}

	return &pb.ListLikedYouResponse{Likers: likers}
}

func (l *LikedResponseMap) Count(i *int) *pb.CountLikedYouResponse {
	return &pb.CountLikedYouResponse{Count: uint64(*i)}
}

func (l *LikedResponseMap) Decision(b bool) *pb.PutDecisionResponse {
	return &pb.PutDecisionResponse{
		MutualLikes: b,
	}
}
