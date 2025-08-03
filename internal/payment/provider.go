package payment

import (
	"database/sql"
	"time"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (repository *PostgresRepository) MakeTransaction(from, to string, amount float64) error {

	tx, err := repository.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE wallet SET balance = ROUND(balance - $1, 2) WHERE address = $2", amount, from)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	_, err = tx.Exec("UPDATE wallet SET balance = ROUND(balance + $1, 2) WHERE address = $2", amount, to)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	timestamp := time.Now()

	_, err = tx.Exec("INSERT INTO transaction(from_address, to_address, created_at, amount) VALUES ($1, $2, $3, $4)", from, to, timestamp, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (repository *PostgresRepository) GetLastTransactions(count int) ([]Transaction, error) {

	rows, err := repository.db.Query("SELECT created_at, from_address, to_address, amount FROM transaction ORDER BY created_at DESC LIMIT $1", count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction

	for rows.Next() {
		var transaction Transaction
		err = rows.Scan(&transaction.CreatedAt, &transaction.From, &transaction.To, &transaction.Amount)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (repository *PostgresRepository) GetBalance(address string) (float64, error) {

	var balance float64

	err := repository.db.QueryRow("SELECT balance FROM wallet WHERE address = $1", address).Scan(&balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}
