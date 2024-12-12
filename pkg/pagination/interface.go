package pagination

type Pagination interface {
	Parse(Request) (Token, error)
}
