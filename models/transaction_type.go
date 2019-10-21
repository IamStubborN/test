package models

type TransactionType struct {
	ID   uint8  `json:"-" db:"type_id"`
	Name string `json:"type" db:"name"`
}
