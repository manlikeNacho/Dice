package slicerepo

import (
	"Dice/src/models"
	"github.com/google/uuid"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	s := New()
	testUser := models.User{
		ID:            1,
		WalletBalance: 0,
		CreatedAt:     time.Now(),
	}
	err := s.CreateUser(&testUser)
	assert.Nil(t, err)
	assert.EqualValues(t, s.Users[0].WalletBalance, testUser.WalletBalance)
	assert.EqualValues(t, s.Users[0].ID, testUser.ID)
}

func TestDb_UpdateUserBalance(t *testing.T) {
	s := New()
	testUser := models.User{
		ID:            1,
		WalletBalance: 0,
		CreatedAt:     time.Now(),
	}
	err := s.CreateUser(&testUser)
	assert.Nil(t, err)
	currentBalance, err := s.UpdateUserBalance(testUser.ID, 100, models.Credit)
	assert.Nil(t, err)
	assert.Equal(t, testUser.WalletBalance, currentBalance)

	currentBalance2, err := s.UpdateUserBalance(testUser.ID, 100, models.Debit)
	assert.Nil(t, err)
	assert.Equal(t, testUser.WalletBalance, currentBalance2)
}

func TestDb_GetCurrentUser(t *testing.T) {
	s := New()
	testUser := models.User{
		ID:            1,
		WalletBalance: 0,
		CreatedAt:     time.Now(),
	}
	err := s.CreateUser(&testUser)
	assert.Nil(t, err)
	currentuser, err := s.GetCurrentUser(testUser.ID)
	assert.Equal(t, currentuser.ID, testUser.ID)
}

func TestDb_GetActiveGame(t *testing.T) {
	s := New()
	testGame1 := &models.Game{
		ID:     uuid.NewString(),
		Status: models.GameRunning,
	}
	err := s.CreateGame(testGame1)
	assert.Nil(t, err)
	activeGame, err := s.GetActiveGame()
	assert.Nil(t, err)
	assert.EqualValues(t, testGame1.Status, activeGame.Status)
}

func TestDb_CreateGame(t *testing.T) {
	s := New()
	testGame := &models.Game{
		Status: models.GameRunning,
	}
	err := s.CreateGame(testGame)
	assert.Nil(t, err)
}

func TestDb_UpdateGame(t *testing.T) {
	s := New()
	testGame := &models.Game{Status: models.GameRunning}
	testGame2 := &models.Game{Status: models.GameEnded}
	err := s.CreateGame(testGame)
	assert.Nil(t, err)
	s.UpdateGame(testGame2)
	assert.Equal(t, testGame2.Status, testGame.Status)
}

func TestDb_CreateTransaction(t *testing.T) {
	s := New()
	transaction := &models.Transaction{
		Type:   models.Credit,
		Amount: 50,
	}
	err := s.CreateTransaction(transaction)
	assert.Nil(t, err)
}

func TestDb_GetTransactions(t *testing.T) {
	s := New()
	transaction := &models.Transaction{
		Type:   models.Credit,
		Amount: 50,
	}
	err := s.CreateTransaction(transaction)
	assert.Nil(t, err)
	transactions, err := s.GetTransactions()
	assert.Equal(t, transaction.Amount, transactions[0].Amount)
}
