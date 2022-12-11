package entity

type Author struct {
	Name       string       `json:"name"`
	Key        string       `json:"key"`
	Born       string       `json:"birth_date"`
	AuthorWork []AuthorWork `json:"entries"`
}

type AuthorWork struct {
	Title string `json:"title"`
}
