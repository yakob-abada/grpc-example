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

func (m *Match) ListLikedYou(recipientUserId string, status int, paginatedReq *PaginatedRequest) (Paginator, error) {
	if paginatedReq == nil {
		paginatedReq = DefaultPaginatedRequest()
	}

	var matches []model.Match
	err := m.db.Offset(int(paginatedReq.Offset())).Limit(int(paginatedReq.Limit())).Find(&matches).
		Where("recipient_user_id = ? AND status = ?", recipientUserId, status).
		Error

	if err != nil {
		return nil, err
	}

	return NewPaginatedResult(matches, true), nil
}

func (m *Match) CountLikedYou(recipientUserId string, status int) (*int64, error) {
	var matches []*model.Match
	var count int64
	err := m.db.Find(&matches).Where("recipient_user_id = ? AND status = ?", recipientUserId, status).Count(&count).Error

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func (m *Match) Decide(recipientUserId string, actorUserId string, match bool) error {
	status := MatchStatusUnMatched

	if match {
		status = MatchStatusMatched
	}

	return m.db.Model(&model.Match{}).
		Where("recipient_user_id = ? AND actor_user_id = ?", recipientUserId, actorUserId).
		Update("status", status).Error
}
