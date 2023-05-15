package models

import "time"

type User struct {
	ID            int
	WalletBalance int
	CreatedAt     time.Time
}

//func (c *User) CreateUser() *User {
//	user := &User{
//		ID:            1,
//		WalletBalance: 0,
//		CreatedAt:     time.Now(),
//	}
//
//	return user
//}

func (c *User) UserBalance() int {
	return c.WalletBalance
}
