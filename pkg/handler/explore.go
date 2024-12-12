package handler

import (
	"context"
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/mapper"
	"github.com/yakob-abada/backend-match/pkg/pagination"
	"github.com/yakob-abada/backend-match/pkg/repo"
)

func NewExploreServer(
	repo repo.LikerRepo, mapper mapper.LikedResponseMapper, pagination pagination.Pagination,
) *ExploreServer {
	return &ExploreServer{
		repo:           repo,
		responseMapper: mapper,
		pagination:     pagination,
	}
}

type ExploreServer struct {
	pb.UnimplementedExploreServiceServer
	repo           repo.LikerRepo
	responseMapper mapper.LikedResponseMapper
	pagination     pagination.Pagination
}

// ListLikedYou returns all users who liked the recipient.
func (s *ExploreServer) ListLikedYou(_ context.Context, req *pb.ListLikedYouRequest) (*pb.ListLikedYouResponse, error) {
	pageToken, err := s.pagination.Parse(req)
	if err != nil {
		return nil, err
	}
	result, err := s.repo.ListLikedYou(
		req.GetRecipientUserId(), repo.MatchStatusMatched, repo.NewPaginatedRequest(pageToken.Offset, pageToken.PageSize),
	)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	nextPageToken := ""

	if result.HasNextPage() {
		nextPageToken = pageToken.Next().String()
	}

	return s.responseMapper.List(result.Results(), nextPageToken), nil
}

// ListNewLikedYou returns all users who liked the recipient excluding those who have been liked in return.
func (s *ExploreServer) ListNewLikedYou(_ context.Context, req *pb.ListLikedYouRequest) (*pb.ListLikedYouResponse, error) {
	pageToken, err := s.pagination.Parse(req)
	if err != nil {
		return nil, err
	}
	result, err := s.repo.ListLikedYou(
		req.GetRecipientUserId(), repo.MatchStatusPending, repo.NewPaginatedRequest(pageToken.Offset, pageToken.PageSize),
	)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	nextPageToken := ""

	if result.HasNextPage() {
		nextPageToken = pageToken.Next().String()
	}

	return s.responseMapper.List(result.Results(), nextPageToken), nil
}

// CountLikedYou counts the number of users who liked the recipient.
func (s *ExploreServer) CountLikedYou(_ context.Context, req *pb.CountLikedYouRequest) (*pb.CountLikedYouResponse, error) {
	result, err := s.repo.CountLikedYou(req.GetRecipientUserId(), repo.MatchStatusMatched)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	return s.responseMapper.Count(result), nil
}

// PutDecision record the decision of the actor to like or pass the recipient.
func (s *ExploreServer) PutDecision(_ context.Context, req *pb.PutDecisionRequest) (*pb.PutDecisionResponse, error) {
	err := s.repo.Decide(req.GetRecipientUserId(), req.GetActorUserId(), req.GetLikedRecipient())
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	return s.responseMapper.Decision(req.LikedRecipient), nil
}
