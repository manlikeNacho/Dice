package controllers

import (
	"Dice/src/models"
	"Dice/src/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	FixedWalletUpdateAmount = 155
	FixedStartGameCost      = 20
	FixedGameWinAmount      = 10
	FixedPlayedDuoAmount    = 5
)

type controller struct {
	repo          repository.Repository
	currentUserID int
	currentGameID int
}

func New(repo repository.Repository, currentUserID int) controller {

	return controller{
		repo:          repo,
		currentUserID: currentUserID,
	}
}

func (ct controller) FundWallet(c *gin.Context) {
	//Check User Balance
	currentUser, err := ct.repo.GetCurrentUser(ct.currentUserID)
	if err != nil {
		log.Printf("error generating user")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error generating user",
		})
		return
	}

	if currentUser.WalletBalance > 35 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wallet balance is already sufficient"})
		return
	}

	// update current users balance
	walletBalance, err := ct.repo.UpdateUserBalance(currentUser.ID, FixedWalletUpdateAmount, models.Credit)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "wallet balance is already sufficient"})
		return
	}

	log.Println("successfully handled fund wallet request")

	c.JSON(http.StatusOK, gin.H{
		"message":        "Successfully Funded wallet",
		"wallet_balance": walletBalance,
	})

}

func (ct controller) GetWalletBallance(c *gin.Context) {
	currentUser, err := ct.repo.GetCurrentUser(ct.currentUserID)
	if err != nil {
		log.Printf("An error occured while getting current user: %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "An error occurred",
		})
		return
	}

	Balance := currentUser.WalletBalance

	log.Println("Successfully generated wallet Balance")

	c.JSON(http.StatusOK, gin.H{
		"Wallet Balance": Balance,
	})
}

func (ct controller) StartGame(c *gin.Context) {
	log.Println("handling start game request")

	// check users balance
	currentUser, err := ct.repo.GetCurrentUser(ct.currentUserID)
	if err != nil {
		log.Printf("error getting current user")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting current user user",
		})
		return
	}

	if currentUser.WalletBalance < FixedStartGameCost {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient current user's balance to start game"})
		return
	}

	//Generate random number
	rand.Seed(time.Now().UnixNano())
	generatednumber := rand.Intn(10) + 2

	//Validate status of the game
	games, err := ct.repo.GetGamesByStatus(models.GameRunning)
	if err != nil {
		log.Printf("Error generating current game")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "An error occurred in the server",
		})
		return
	}
	if len(games) != 0 {
		//log
		log.Printf("An active game is still in session")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "An active game is still in session",
		})

		return
	}

	newGame := &models.Game{
		ID:               uuid.New().String(),
		GeneratedNumber:  generatednumber,
		Status:           models.GameRunning,
		NextRollPosition: models.Roll1,
		CreatedAt:        time.Now(),
	}

	if err := ct.repo.CreateGame(newGame); err != nil {
		// log error
		log.Println("Unable to create new game")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to create new game",
		})
		// return error
		return
	}

	// deduct user
	balance, err := ct.repo.UpdateUserBalance(ct.currentUserID, FixedStartGameCost, models.Debit)
	if err != nil {
		return
	}

	// return message to the user
	c.JSON(http.StatusOK, gin.H{
		"balance": balance,
		"message": "successfully created game",
	})
}

func (ct controller) RollDice(c *gin.Context) {
	// generatate a random number
	rand.Seed(time.Now().UnixNano())
	generatednumber := rand.Intn(7)

	// getActiveGame
	activeGame, err := ct.repo.GetActiveGame()
	if err != nil {
		log.Println("Unable to get active game")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No active game in session",
		})
		return
	}
	switch activeGame.NextRollPosition {
	case models.Roll1:
		activeGame.Roll1 = generatednumber
		activeGame.NextRollPosition = models.Roll2

		if err := ct.repo.UpdateGame(activeGame); err != nil {
			log.Printf("error updating game: %v \n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Unable to handle request",
			})
			return
		}
		// return message to user and the result, successfully rolled first die
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully rolled the first die",
			"roll":    activeGame.Roll1,
		})

	case models.Roll2:
		activeGame.Roll2 = generatednumber
		activeGame.NextRollPosition = models.Roll1
		ct.handleSecondRoll(activeGame, c)

	default:
	}
}

func (ct controller) handleSecondRoll(activeGame *models.Game, c *gin.Context) {
	if err := ct.repo.UpdateGame(activeGame); err != nil {
		log.Printf("error updating game: %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to handle request",
		})
		return
	}
	// user wins if the generated number is equal to the sum of the two die rolls
	usersTotal := activeGame.Roll1 + activeGame.Roll2
	userWins := usersTotal == activeGame.GeneratedNumber

	// deduct user after playing second die
	currentBalance, err := ct.repo.UpdateUserBalance(ct.currentUserID, FixedPlayedDuoAmount, models.Debit)
	if err != nil {
		log.Println("error deducting user")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "couldn't handle request",
		})
		return
	}

	//return response
	res := models.GameResponse{
		Win:            userWins,
		Roll1:          activeGame.Roll1,
		Roll2:          activeGame.Roll2,
		Total:          usersTotal,
		CurrentBalance: currentBalance,
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "successfully rolled the second die",
		"response": res,
	})

}

func (ct controller) EndGame(c *gin.Context) {
	// get active game
	activeGame, err := ct.repo.GetActiveGame()
	if err != nil {
		log.Printf("error ending game: %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "no active game currently running",
		})

		return
	}
	activeGame.Status = models.GameEnded

	if err := ct.repo.UpdateGame(activeGame); err != nil {
		log.Printf("error updating game: %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to end game at this moment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully ended the game"})
}

func (ct controller) GetTransactions(c *gin.Context) {
	transactions, err := ct.repo.GetTransactions()
	if err != nil {
		log.Printf("error getting transactions")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting transactions",
		})
		return
	}

	c.JSON(http.StatusOK, transactions)

}
