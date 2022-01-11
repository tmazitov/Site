package models

type Message struct {
	UUID      string
	FromUser  string
	ToUser    string
	Date      int
	Message   string
	IsChanged bool
}
