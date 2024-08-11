package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TransactionsWithdraw struct {
	gorm.Model
	ID                uint    `json:"id"`
	PaymentAccountsID int     `json:"payment_accounts_id"`
	UsersID           int     `json:"users_id"`
	Currency          string  `json:"currency"`
	Amount            float64 `json:"amount"`
	DiscountCode      string  `json:"discount_code"`
	AmountAfter       float64 `json:"amount_after"`
	Status            string  `json:"status"`
	ToAddress         string  `json:"to_address"`
}

// status
var (
	transactionsWithdrawStatusAccept = "accept"
	transactionsWithdrawStatusReject = "reject"
)

func (t *TransactionsWithdraw) Process(timeout int) {
	time.Sleep(time.Duration(timeout) * time.Second)
	t.Status = transactionsWithdrawStatusAccept
	if timeout > 30 {
		t.Status = transactionsWithdrawStatusReject
	}
}
