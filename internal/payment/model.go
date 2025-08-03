package payment

import "time"

type Transaction struct {
	CreatedAt time.Time `json:"created_at"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Amount    float64   `json:"amount"`
}

type MakeTransactionDTO struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}
