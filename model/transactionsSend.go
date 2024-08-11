package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TransactionsSend struct {
	gorm.Model
	ID        uint    `json:"id"`
	UsersID   int     `json:"users_id"`
	ToUsersID int     `json:"to_users_id"`
	Currency  string  `json:"currency"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	ToAddress string  `json:"to_address"`
}

// status
var (
	transactionsSendStatusAccept = "accept"
	transactionsSendStatusReject = "reject"
)

func (t *TransactionsSend) Process(timeout int) {
	time.Sleep(time.Duration(timeout) * time.Second)
	t.Status = transactionsSendStatusAccept
	if timeout > 30 {
		t.Status = transactionsSendStatusReject
	}
}
