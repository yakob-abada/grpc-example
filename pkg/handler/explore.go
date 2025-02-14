package handler

import (
	"context"
	pb "github.com/yakob-abada/backend-match/explore/proto"
	"github.com/yakob-abada/backend-match/pkg/mapper"
	"github.com/yakob-abada/backend-match/pkg/model"
	"github.com/yakob-abada/backend-match/pkg/pagination"
	"github.com/yakob-abada/backend-match/pkg/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
func (s *ExploreServer) ListLikedYou(ctx context.Context, req *pb.ListLikedYouRequest) (*pb.ListLikedYouResponse, error) {
	pageToken, err := s.pagination.Parse(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "page token parsing failed: %v", err)
	}
	result, err := s.repo.ListAllLikedYou(ctx, repo.NewPaginatedRequest(pageToken.Offset, pageToken.PageSize), req.GetRecipientUserId())
	if err != nil {
		//log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "failed to get result: %v", err)
	}

	nextPageToken := ""

	if result.HasNextPage() {
		nextPageToken = pageToken.Next().String()
	}

	return s.responseMapper.List(result.Results(), nextPageToken), nil
}

// ListNewLikedYou returns all users who liked the recipient excluding those who have been liked in return.
func (s *ExploreServer) ListNewLikedYou(ctx context.Context, req *pb.ListLikedYouRequest) (*pb.ListLikedYouResponse, error) {
	pageToken, err := s.pagination.Parse(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "page token parsing failed: %v", err)
	}
	result, err := s.repo.ListLikedYou(ctx, []int{model.MatchStatusPending, model.MatchStatusUnMatched}, repo.NewPaginatedRequest(pageToken.Offset, pageToken.PageSize), req.GetRecipientUserId())
	if err != nil {
		//log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "failed to get result: %v", err)
	}

	nextPageToken := ""

	if result.HasNextPage() {
		nextPageToken = pageToken.Next().String()
	}

	return s.responseMapper.List(result.Results(), nextPageToken), nil
}

// CountLikedYou counts the number of users who liked the recipient.
func (s *ExploreServer) CountLikedYou(ctx context.Context, req *pb.CountLikedYouRequest) (*pb.CountLikedYouResponse, error) {
	result, err := s.repo.CountLikedYou(ctx, req.GetRecipientUserId())
	if err != nil {
		//log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "failed to get result: %v", err)
	}

	return s.responseMapper.Count(result), nil
}

// PutDecision record the decision of the actor to like or pass the recipient.
func (s *ExploreServer) PutDecision(ctx context.Context, req *pb.PutDecisionRequest) (*pb.PutDecisionResponse, error) {
	err := s.repo.Decide(ctx, req.GetRecipientUserId(), req.GetActorUserId(), req.GetLikedRecipient())
	if err != nil {
		//log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "failed to update: %v", err)
	}

	return s.responseMapper.Decision(req.LikedRecipient), nil
}
