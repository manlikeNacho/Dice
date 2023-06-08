package repository

import (
	"Dice/src/models"
)

// go:generate mockgen -source repository.go -destination ./mocks/repository.go -package mocks Repository
type Repository interface {
	CreateUser(user *models.User) error
	GetCurrentUser(userID int) (*models.User, error)
	UpdateUserBalance(userID int, amount int, transactionType models.TransactionType) (int, error)
	CreateGame(game *models.Game) error
	GetActiveGame() (*models.Game, error)
	GetGamesByStatus(status models.GameStatus) ([]*models.Game, error)
	UpdateGame(game *models.Game) error
	CreateTransaction(newTransaction *models.Transaction) error
	GetTransactions() ([]*models.Transaction, error)
}
