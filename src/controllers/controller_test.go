package controllers_test

import (
	"Dice/src/controllers"
	"Dice/src/models"
	"Dice/src/repository/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_FundWallet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepo := mocks.NewMockRepository(controller)
	ctrl := controllers.New(mockRepo, 1)

	mockRepo.EXPECT().GetCurrentUser(1).Return(&models.User{WalletBalance: 0, ID: 1}, nil)
	mockRepo.EXPECT().UpdateUserBalance(1, controllers.FixedWalletUpdateAmount, models.Credit).Return(155, nil)

	router.POST("/fund_wallet", ctrl.FundWallet)
	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/fund_wallet", nil)
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, request)
	assert.Equal(t, 200, rr.Code)
	assert.NotNil(t, rr.Body)
}

func TestController_GetWalletBallance(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepo := mocks.NewMockRepository(controller)
	ctrl := controllers.New(mockRepo, 1)

	mockRepo.EXPECT().GetCurrentUser(1).Return(&models.User{WalletBalance: 155, ID: 1}, nil)

	router.GET("/get_wallet_balance", ctrl.GetWalletBallance)
	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/get_wallet_balance", nil)
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, request)
	assert.Equal(t, 200, rr.Code)
	assert.NotNil(t, rr.Body)
}

func TestController_RollDice(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepo := mocks.NewMockRepository(controller)
	ctrl := controllers.New(mockRepo, 1)

	mockRepo.EXPECT().GetActiveGame().Return(&models.Game{GeneratedNumber: 7, NextRollPosition: models.Roll1}, nil)

	router.GET("/roll_die", ctrl.RollDice)
	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/roll_die", nil)
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, request)
	assert.Equal(t, 200, rr.Code)
	assert.NotNil(t, rr.Body)
}

func TestController_RollDice_2ndRoll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockRepo := mocks.NewMockRepository(controller)
	ctrl := controllers.New(mockRepo, 1)

	mockRepo.EXPECT().GetActiveGame().Return(&models.Game{GeneratedNumber: 7, NextRollPosition: models.Roll2}, nil)
	mockRepo.EXPECT().UpdateUserBalance(1, controllers.FixedPlayedDuoAmount, models.Debit).Return(150, nil)

	router.GET("/roll_die", ctrl.RollDice)
	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/roll_die", nil)
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, request)
	assert.Equal(t, 200, rr.Code)
	assert.NotNil(t, rr.Body)
}
