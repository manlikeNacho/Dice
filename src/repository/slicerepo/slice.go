package slicerepo

import (
	"Dice/src/models"
	"Dice/src/repository"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type Db struct {
	Users           []*models.User
	Games           []*models.Game
	TransactionsLog []*models.Transaction
}

var _ repository.Repository = &Db{}

func New() *Db {
	users := make([]*models.User, 0)
	games := make([]*models.Game, 0)
	transactions := make([]*models.Transaction, 0)

	return &Db{
		Users:           users,
		Games:           games,
		TransactionsLog: transactions,
	}
}

func (s *Db) CreateUser(user *models.User) error {
	s.Users = append(s.Users, user)
	log.Printf("new user created, total_users: %v", len(s.Users))

	return nil
}

func (s *Db) GetCurrentUser(userID int) (*models.User, error) {
	for _, user := range s.Users {
		if user.ID == userID {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found, userID: %v", userID)
}

func (s *Db) UpdateUserBalance(userID int, amount int, transactionType models.TransactionType) (int, error) {
	currentUser, err := s.GetCurrentUser(userID)
	if err != nil {
		log.Printf("Invalid User")
		return 0, errors.New("Invalid User")
	}
	switch transactionType {
	case models.Credit:
		currentUser.WalletBalance += amount
	case models.Debit:
		currentUser.WalletBalance -= amount
	}
	log.Printf("amount: %v, current_balance: %v", amount, currentUser.WalletBalance)
	//update transactions auth
	transaction := &models.Transaction{
		ID:        uuid.New().String(),
		Type:      transactionType,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
	s.TransactionsLog = append(s.TransactionsLog, transaction)

	return currentUser.WalletBalance, nil
}

func (s *Db) CreateGame(game *models.Game) error {
	s.Games = append(s.Games, game)
	return nil
}

func (s *Db) GetActiveGame() (*models.Game, error) {
	for _, game := range s.Games {
		if game.Status == models.GameRunning {
			return game, nil
		}
	}

	return nil, errors.New("no active game found")
}

func (s *Db) GetGamesByStatus(status models.GameStatus) ([]*models.Game, error) {
	var games []*models.Game
	//loop
	for _, game := range s.Games {
		if game.Status == status {
			games = append(games, game)
		}
	}

	return games, nil
}

func (s *Db) UpdateGame(game *models.Game) error {
	fmt.Println(s.Games[0])
	for _, g := range s.Games {
		if game.ID == g.ID {
			g.Roll1 = game.Roll1
			g.Roll2 = game.Roll2
			g.NextRollPosition = game.NextRollPosition
			g.GeneratedNumber = game.GeneratedNumber
			g.Status = game.Status
		}

		return nil
	}

	return fmt.Errorf("game not found, gameID: %v", game.ID)
	return nil
}

func (s *Db) CreateTransaction(newTransaction *models.Transaction) error {
	s.TransactionsLog = append(s.TransactionsLog, newTransaction)
	return nil
}

func (s *Db) GetTransactions() ([]*models.Transaction, error) {
	return s.TransactionsLog, nil
}
