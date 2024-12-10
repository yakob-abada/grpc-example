package handler

import (
	"context"
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/mapper"
	"github.com/yakob-abada/backend-match/pkg/repo"
)

func NewExploreServer(repo repo.LikerRepo, mapper mapper.LikedResponseMapper) *ExploreServer {
	return &ExploreServer{
		repo:           repo,
		responseMapper: mapper,
	}
}

type ExploreServer struct {
	pb.UnimplementedExploreServiceServer
	repo           repo.LikerRepo
	responseMapper mapper.LikedResponseMapper
}

// ListLikedYou returns all users who liked the recipient.
func (s *ExploreServer) ListLikedYou(_ context.Context, req *pb.ListLikedYouRequest) (*pb.ListLikedYouResponse, error) {
	result, err := s.repo.ListLikedYou(req.RecipientUserId, true)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	return s.responseMapper.List(result), nil
}

// ListNewLikedYou returns all users who liked the recipient excluding those who have been liked in return.
func (s *ExploreServer) ListNewLikedYou(_ context.Context, req *pb.ListLikedYouRequest) (*pb.ListLikedYouResponse, error) {
	result, err := s.repo.ListLikedYou(req.RecipientUserId, false)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	return s.responseMapper.List(result), nil
}

// CountLikedYou counts the number of users who liked the recipient.
func (s *ExploreServer) CountLikedYou(_ context.Context, req *pb.CountLikedYouRequest) (*pb.CountLikedYouResponse, error) {
	result, err := s.repo.CountLikedYou(req.RecipientUserId, false)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	return s.responseMapper.Count(result), nil
}

// PutDecision record the decision of the actor to like or pass the recipient.
func (s *ExploreServer) PutDecision(_ context.Context, req *pb.PutDecisionRequest) (*pb.PutDecisionResponse, error) {
	err := s.repo.Decide(req.RecipientUserId, req.ActorUserId, req.LikedRecipient)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	return s.responseMapper.Decision(req.LikedRecipient), nil
}
