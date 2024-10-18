package models

import (
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/enums"
	"time"
)

type PaymentRequest struct {
	UserId         string                  `json:"user_id"`
	PaymentType    enums.PaymentType       `json:"payment_type"`
	PaymentMethod  enums.PaymentMethodType `json:"payment_method"`
	PaymentStatus  enums.PaymentStatus     `json:"payment_status"`
	Amount         float64                 `json:"amount"`
	PaidFromWallet float64                 `json:"paid_from_wallet"`
	IdempotencyKey string                  `json:"idempotency_key"`
	RequestedAt    time.Time               `json:"requested_at"`
}
