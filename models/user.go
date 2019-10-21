package models

type User struct {
	ID      uint64  `json:"id" db:"id"`
	Balance float64 `json:"balance" db:"balance"`
}
