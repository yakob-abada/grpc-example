package bootstrap

import "github.com/yakob-abada/backend-match/pkg/handler"

func NewExploreServer() *handler.ExploreServer {
	return handler.NewExploreServer(nil, nil)
}
