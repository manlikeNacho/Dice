package models

import "time"

type TransactionType string

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)

type Transaction struct {
	ID        string
	Type      TransactionType
	Amount    int
	CreatedAt time.Time
}
