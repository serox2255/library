package entity

type Book struct {
	Title      string   `json:"title"`
	Publishers []string `json:"publishers"`
	Authors    []Key    `json:"authors"`
	Isbn       []string `json:"isbn_10"`
}

type Key struct {
	AuthorKey string `json:"key"`
}
