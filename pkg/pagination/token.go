package pagination

// Token is a page token that uses an offset to delineate which page to fetch.
type Token struct {
	// Offset of the page.
	Offset int
	// RequestChecksum is the checksum of the request that generated the page token.
	RequestChecksum uint32

	PageSize int
}

// Next returns the next page token for the provided Request.
func (p Token) Next() Token {
	p.Offset += p.PageSize
	return p
}

// String returns a string representation of the page token.
func (p Token) String() string {
	return EncodePageTokenStruct(&p)
}
