package main

import (
	"Dice/src/controllers"
	"Dice/src/models"
	"Dice/src/repository/slicerepo"
	"github.com/gin-gonic/gin"
	"log"
	"time"

	_ "github.com/gin-gonic/gin"
)

func main() {
	// initialize db
	repo := slicerepo.New()

	// initialize the current user
	currentUser := &models.User{
		ID:            1,
		WalletBalance: 0,
		CreatedAt:     time.Now(),
	}

	if err := repo.CreateUser(currentUser); err != nil {
		log.Printf("could not initialize current user: %v \n", err)
		return
	}

	log.Printf("current user initialized, userID %v \n", currentUser.ID)

	// initialize the controller
	ctrl := controllers.New(repo, currentUser.ID)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome amigo!",
		})
	})
	r.POST("/fund_wallet", ctrl.FundWallet)
	r.GET("/get_wallet_balance", ctrl.GetWalletBallance)
	r.POST("/start_game", ctrl.StartGame)
	r.POST("/roll_die", ctrl.RollDice)
	r.POST("/end_game", ctrl.EndGame)
	r.GET("/transactions", ctrl.GetTransactions)
	r.Run(":5000")
}
