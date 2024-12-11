package bootstrap

import (
	"github.com/yakob-abada/backend-match/pkg/handler"
	"github.com/yakob-abada/backend-match/pkg/mapper"
	"github.com/yakob-abada/backend-match/pkg/repo"
	"gorm.io/gorm"
)

func NewExploreServer(db *gorm.DB) *handler.ExploreServer {
	return handler.NewExploreServer(repo.NewMatch(db), mapper.NewLikedResponseMap())
}
