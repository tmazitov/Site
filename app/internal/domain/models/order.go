package models

type Order struct {
	Title       string
	Date        int64
	HourCount   int
	Status      string
	FromAddress string
	ToAddress   []string
	Comment     string
}
