package pagination

// PageToken is a page token that uses an offset to delineate which page to fetch.
type PageToken struct {
	// Offset of the page.
	Offset string
}
