package models

type Books struct {
	ID     string
	Title  string
	ISBN   string
	Author Author
}
type Author struct {
	FirstName, LastName string
}
