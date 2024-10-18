package models

type TRecette struct {
	ID     string  `json:"id"`
	Date   string  `json:"date"`
	Label  string  `json:"label"`
	Amount float64 `json:"amount"`
}

type TAddRecette struct {
	Label  string  `json:"label"`
	Amount float64 `json:"amount"`
}
