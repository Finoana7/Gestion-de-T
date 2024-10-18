package models

type TTransaction struct {
	ID     string  `json:"id"`
	Date   string  `json:"date"`
	Label  string  `json:"label"`
	Amount float64 `json:"amount"`
}

type TAddTransaction struct {
	TAddRecette
}
