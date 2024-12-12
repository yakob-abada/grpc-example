package mapper

import (
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/model"
)

type LikedResponseMapper interface {
	List([]model.Match, string) *pb.ListLikedYouResponse
	Count(int64) *pb.CountLikedYouResponse
	Decision(bool) *pb.PutDecisionResponse
}
