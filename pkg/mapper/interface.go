package mapper

import (
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/model"
)

type LikedResponseMapper interface {
	List([]*model.Likes) *pb.ListLikedYouResponse
	Count(*int) *pb.CountLikedYouResponse
	Decision(bool) *pb.PutDecisionResponse
}
