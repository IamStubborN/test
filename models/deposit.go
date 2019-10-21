package models

import "time"

type Deposit struct {
	ID            uint64    `json:"depositId" db:"id"`
	UserID        uint64    `json:"userId" db:"user_id"`
	Amount        float64   `json:"amount" db:"amount"`
	BalanceBefore float64   `json:"-" db:"balance_before"`
	BalanceAfter  float64   `json:"-" db:"balance_after"`
	Date          time.Time `json:"-" db:"date"`
}
