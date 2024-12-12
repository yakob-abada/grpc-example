package repo

const maxSize = 50

type PaginatedRequest struct {
	offset int
	limit  int
}

func DefaultPaginatedRequest() *PaginatedRequest {
	return &PaginatedRequest{
		offset: 0,
		limit:  maxSize,
	}
}

func NewPaginatedRequest(offset, limit int) *PaginatedRequest {
	return &PaginatedRequest{
		offset: offset,
		limit:  limit,
	}
}

func (p *PaginatedRequest) Limit() int {
	if p.limit > maxSize {
		p.limit = maxSize
	}

	return p.limit
}

func (p *PaginatedRequest) Offset() int {
	return p.offset
}
