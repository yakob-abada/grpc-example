package pagination

import (
	"github.com/stretchr/testify/mock"
)

type PageTokenMock struct {
	mock.Mock
}

func (p *PageTokenMock) Parse(request Request) (Token, error) {
	args := p.Called(request)

	if args.Error(1) != nil {
		return Token{}, args.Error(1)
	}

	return args.Get(0).(Token), args.Error(1)
}
