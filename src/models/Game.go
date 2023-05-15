package models

import "time"

type (
	GameStatus   string
	RollPosition string
)

const (
	GameRunning GameStatus = "running"
	GameEnded   GameStatus = "ended"

	Roll1 RollPosition = "Roll1"
	Roll2 RollPosition = "Roll2"
)

type Game struct {
	ID               string
	GeneratedNumber  int
	Roll1            int
	Roll2            int
	NextRollPosition RollPosition //
	Status           GameStatus
	CreatedAt        time.Time
}

type GameResponse struct {
	Win            bool
	Roll1          int
	Roll2          int
	Total          int
	CurrentBalance int
}
