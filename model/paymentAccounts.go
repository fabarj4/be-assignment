package model

import (
	"github.com/jinzhu/gorm"
)

type PaymentAccounts struct {
	gorm.Model
	ID            uint   `json:"id"`
	UsersID       int    `json:"users_id"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	Active        bool   `json:"active"`
}
