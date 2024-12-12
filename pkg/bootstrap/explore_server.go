package bootstrap

import (
	"github.com/yakob-abada/backend-match/pkg/handler"
	"github.com/yakob-abada/backend-match/pkg/mapper"
	"github.com/yakob-abada/backend-match/pkg/pagination"
	"github.com/yakob-abada/backend-match/pkg/repo"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func NewExploreServer(db *gorm.DB) *handler.ExploreServer {
	pageSize, err := strconv.ParseInt(os.Getenv("PAGE_SIZE"), 10, 64)
	if err != nil {
		pageSize = 2
	}

	return handler.NewExploreServer(
		repo.NewMatch(db),
		mapper.NewLikedResponseMap(),
		pagination.NewPageToken(pageSize),
	)
}
