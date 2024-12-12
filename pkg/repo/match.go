package repo

import (
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

const (
	MatchStatusPending = iota
	MatchStatusMatched
	MatchStatusUnMatched
)

// ListLikedYou return list of matches based on recipientUserId and status, result returns with certain limit.
func (m *Match) ListLikedYou(recipientUserId string, status int, paginatedReq *PaginatedRequest) (Paginator, error) {
	if paginatedReq == nil {
		paginatedReq = DefaultPaginatedRequest()
	}

	var matches []model.Match
	err := m.db.Offset(paginatedReq.Offset()).Limit(paginatedReq.Limit()+1).Find(&matches).
		Where("recipient_user_id = ? AND status = ?", recipientUserId, status).
		Error

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
func (m *Match) CountLikedYou(recipientUserId string, status int) (*int64, error) {
	var matches []*model.Match
	var count int64
	err := m.db.Find(&matches).Where("recipient_user_id = ? AND status = ?", recipientUserId, status).Count(&count).Error

	if err != nil {
		return nil, err
	}

	return &count, nil
}

// Decide to updated status to match or unmatch.
func (m *Match) Decide(recipientUserId string, actorUserId string, match bool) error {
	status := MatchStatusUnMatched

	if match {
		status = MatchStatusMatched
	}

	return m.db.Model(&model.Match{}).
		Where("recipient_user_id = ? AND actor_user_id = ?", recipientUserId, actorUserId).
		Update("status", status).Error
}
