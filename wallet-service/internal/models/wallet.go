package models

import "github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/enums"

type Wallet struct {
	Id              string                `json:"id"`
	Credit          float32               `json:"credit"`
	Balance         float32               `json:"balance"`
	Debit           float32               `json:"deposit"`
	TransactionType enums.TransactionType `json:"type"`
	UserId          string                `json:"user_id"`
}
