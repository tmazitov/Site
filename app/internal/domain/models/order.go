package models

type Order struct {
	UUID    string
	Title   string
	Writer  string
	Worker  string
	Date    int64
	Status  string
	Comment string
	Price   int
}
