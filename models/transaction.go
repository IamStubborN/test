package models

import "time"

type Transaction struct {
	ID     uint64  `json:"transactionId" db:"id"`
	UserID uint64  `json:"userId" db:"user_id"`
	Amount float64 `json:"amount" db:"amount"`
	TransactionType
	BalanceBefore float64   `json:"-" db:"balance_before"`
	BalanceAfter  float64   `json:"-" db:"balance_after"`
	Date          time.Time `json:"-" db:"date"`
}
