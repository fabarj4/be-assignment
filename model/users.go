package model

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	ID       uint   `json:"id"`
	Email    string `json:"email" gorm:"index:idx_email,unique"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Address  string `json:"address"`
}

type UsersTransactionsWithdraws struct {
	ID                  uint    `json:"id"`
	PaymentAccountsID   int     `json:"payment_accounts_id"`
	PaymentAccountsName string  `json:"payment_accounts_name"`
	UsersID             int     `json:"users_id"`
	UserName            string  `json:"users_name"`
	Currency            string  `json:"currency"`
	Amount              float64 `json:"amount"`
	DiscountCode        string  `json:"discount_code"`
	AmountAfter         float64 `json:"amount_after"`
	Status              string  `json:"status"`
	ToAddress           string  `json:"to_address"`
}

type UsersTransactionsSends struct {
	ID          uint    `json:"id"`
	ToUsersID   int     `json:"to_users_id"`
	ToUsersName string  `json:"to_users_name"`
	UsersID     int     `json:"users_id"`
	UsersName   string  `json:"users_name"`
	Currency    string  `json:"currency"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
	ToAddress   string  `json:"to_address"`
}

type UsersPaymentAccounts struct {
	ID            uint   `json:"id"`
	UsersID       int    `json:"users_id"`
	UsersName     string `json:"users_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	Active        bool   `json:"active"`
}
