package repo

const maxSize = 50

type PaginatedRequest struct {
	offset int64
	limit  int64
}

func DefaultPaginatedRequest() *PaginatedRequest {
	return &PaginatedRequest{
		offset: 0,
		limit:  maxSize,
	}
}

func NewPaginatedRequest(offset, limit int64) *PaginatedRequest {
	return &PaginatedRequest{
		offset: offset,
		limit:  limit,
	}
}

func (p *PaginatedRequest) Limit() int64 {
	if p.limit > maxSize {
		p.limit = maxSize
	}

	return p.limit
}

func (p *PaginatedRequest) Offset() int64 {
	return p.offset
}
