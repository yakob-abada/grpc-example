package repo

import "github.com/yakob-abada/backend-match/pkg/model"

type Paginator interface {
	Results() []model.Match
	HasNextPage() bool
}

func NewPaginatedResult(results []model.Match, hasNextPage bool) *PaginatedResults {
	return &PaginatedResults{
		defaultSize: 50,
		results:     results,
		hasNextPage: hasNextPage,
	}
}

type PaginatedResults struct {
	defaultSize int
	results     []model.Match
	hasNextPage bool
}

func (p *PaginatedResults) Results() []model.Match {
	return p.results
}

func (p *PaginatedResults) HasNextPage() bool {
	return p.hasNextPage
}
