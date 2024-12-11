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

func (m *Match) ListLikedYou(recipientUserId string, status int) (Paginator, error) {
	var matches []model.Match
	err := m.db.Find(&matches).Where("recipient_user_id = ? AND status = ?", recipientUserId, status).Error

	if err != nil {
		return nil, err
	}

	return NewPaginatedResult(matches, false), nil
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
