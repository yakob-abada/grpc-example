package pagination

import (
	"fmt"
)

// pageTokenChecksumMask is a random bitmask applied to offset-based page token checksums.
//
// Change the bitmask to force checksum failures when changing the page token implementation.
const pageTokenChecksumMask uint32 = 0x9acb0442

type PageToken struct {
	PageSize int
}

func NewPageToken(pageSize int) *PageToken {
	return &PageToken{PageSize: pageSize}
}

// Parse parses an offset-based page token from the provided Request.
//
// If the request does not have a page token, a page token with offset 0 will be returned.
func (p *PageToken) Parse(request Request) (Token, error) {
	requestChecksum, err := calculateRequestChecksum(request)
	if err != nil {
		return Token{}, err
	}
	requestChecksum ^= pageTokenChecksumMask // apply checksum mask for Token
	if request.GetPaginationToken() == "" {
		offset := 0

		return Token{
			Offset:          offset,
			RequestChecksum: requestChecksum,
			PageSize:        p.PageSize,
		}, nil
	}
	var pageToken Token
	if err := DecodePageTokenStruct(request.GetPaginationToken(), &pageToken); err != nil {
		return Token{}, err
	}
	if pageToken.RequestChecksum != requestChecksum {
		return Token{}, fmt.Errorf(
			"checksum mismatch (got 0x%x but expected 0x%x)", pageToken.RequestChecksum, requestChecksum,
		)
	}

	return pageToken, nil
}
