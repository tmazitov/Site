package models

type Order struct {
	UUID        string
	Title       string
	Writer      string
	Date        int64
	HourCount   int
	Status      string
	FromAddress string
	ToAddress   string
	Comment     string
}
