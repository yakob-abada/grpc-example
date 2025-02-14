package repo

import (
	"context"
	"github.com/yakob-abada/backend-match/pkg/model"
	"gorm.io/gorm"
)

type Match struct {
	db *gorm.DB
}

func NewMatch(db *gorm.DB) *Match {
	return &Match{
		db: db,
	}
}

// ListAllLikedYou return list of matches based on recipientUserId and, result returns with certain limit.
func (m *Match) ListAllLikedYou(ctx context.Context, paginatedReq *PaginatedRequest, recipientUserId string) (Paginator, error) {
	if paginatedReq == nil {
		paginatedReq = DefaultPaginatedRequest()
	}

	var matches []model.Match
	err := m.db.WithContext(ctx).Offset(paginatedReq.Offset()).Limit(paginatedReq.Limit()+1).
		Where("recipient_user_id = ?", recipientUserId).
		Find(&matches).Error

	if err != nil {
		return nil, err
	}

	// A tricky way to answer has next page question by adding limit by own and see if result is more than limit then it
	// means there is more to be added in following page. the slice go cleaned up from the extra element.
	hasNextPage := false
	if len(matches) > paginatedReq.Limit() {
		hasNextPage = true
		matches = matches[:len(matches)-1]
	}

	return NewPaginatedResult(matches, hasNextPage), nil
}

// ListLikedYou return list of matches based on recipientUserId and statuses, result returns with certain limit.
func (m *Match) ListLikedYou(ctx context.Context, statuses []int, paginatedReq *PaginatedRequest, recipientUserId string) (Paginator, error) {
	if paginatedReq == nil {
		paginatedReq = DefaultPaginatedRequest()
	}

	var matches []model.Match
	err := m.db.WithContext(ctx).Offset(paginatedReq.Offset()).Limit(paginatedReq.Limit()+1).
		Where("recipient_user_id = ? AND status in (?)", recipientUserId, statuses).
		Find(&matches).Error

	if err != nil {
		return nil, err
	}

	// A tricky way to answer has next page question by adding limit by own and see if result is more than limit then it
	// means there is more to be added in following page. the slice go cleaned up from the extra element.
	hasNextPage := false
	if len(matches) > paginatedReq.Limit() {
		hasNextPage = true
		matches = matches[:len(matches)-1]
	}

	return NewPaginatedResult(matches, hasNextPage), nil
}

// CountLikedYou returns count of pending matches.
func (m *Match) CountLikedYou(ctx context.Context, recipientUserId string) (int64, error) {
	var matches []*model.Match
	var count int64
	err := m.db.WithContext(ctx).Find(&matches).Where("recipient_user_id = ?", recipientUserId).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

// Decide to updated status to match or unmatch.
func (m *Match) Decide(ctx context.Context, recipientUserId string, actorUserId string, match bool) error {
	status := model.MatchStatusUnMatched

	if match {
		status = model.MatchStatusMatched
	}

	return m.db.WithContext(ctx).Model(&model.Match{}).
		Where("recipient_user_id = ? AND actor_user_id = ?", recipientUserId, actorUserId).
		Update("status", status).Error
}
