package payment

import "time"

// Transaction - структура для хранения записей о переводе средств из БД
type Transaction struct {
	CreatedAt time.Time `json:"created_at"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Amount    float64   `json:"amount"`
}

// MakeTransactionDTO - модель данных, которые должны быть получены в теле запроса /send
type MakeTransactionDTO struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

// GetLastTransactionsOut - модель данных ответа на запрос /transactions
type GetLastTransactionsOut struct {
	Transactions []Transaction `json:"transactions"`
}

// GetBalanceOut - модель данных ответа на запрос /wallet/:address/balance
type GetBalanceOut struct {
	Balance float64 `json:"balance"`
}
