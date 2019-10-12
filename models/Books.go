package models

type Books struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	ISBN   string  `json:"isbn"`
	Author *Author `json:"author"`
}
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
