package models

import (
	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/domain/enums"
	"time"
)

type WalletTransaction struct {
	ID                int32                 `json:"id"`
	Credit            float32               `json:"credit"`
	Balance           float32               `json:"balance"`
	Debit             float32               `json:"deposit"`
	TransactionType   enums.TransactionType `json:"type"`
	UserId            int32                 `json:"user_id"`
	CreatedAt         time.Time             `json:"created_at"`
	TransactionReason string                `json:"transaction_reason"`
}

type Coupon struct {
	ID        int32     `gorm:"primaryKey" json:"id"`
	CouponKey string    `gorm:"uniqueIndex;not null" json:"coupon_key"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
	Amount    float32   `json:"amount"`
	// Associations
	Usages []CouponUsage `gorm:"foreignKey:CouponID" json:"usages"`
}

type CouponUsage struct {
	ID       int32     `gorm:"primaryKey" json:"id"`
	CouponID int32     `gorm:"not null;index" json:"coupon_id"`
	WalletID int32     `gorm:"not null;index" json:"wallet_id"`
	UserId   int32     `gorm:"not null;index" json:"user_id"`
	UsedAt   time.Time `json:"used_at"`
	// Associations
	Coupon Coupon            `gorm:"foreignKey:CouponID" json:"coupon"`
	Wallet WalletTransaction `gorm:"foreignKey:WalletID" json:"wallet"`
}
