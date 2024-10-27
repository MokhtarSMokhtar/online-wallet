package models

import (
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/enums"
	"time"
)

type PaymentRequest struct {
	Id                   string                  `bson:"id"`
	UserId               string                  `bson:"user_id"`
	ChargeId             string                  `bson:"charge_id"`
	CustomerId           int                     `bson:"customer_id"`
	PaymentType          enums.PaymentType       `bson:"payment_type"`
	PaymentMethod        enums.PaymentMethodType `bson:"payment_method"`
	PaymentStatus        enums.PaymentStatus     `bson:"payment_status"`
	Amount               float64                 `bson:"amount"`
	PaidFromWallet       float64                 `bson:"paid_from_wallet"`
	IdempotencyKey       string                  `bson:"idempotency_key"`
	IdempotencyKeyExpiry time.Time               `bson:"idempotency_key_expiry"`
	RequestedAt          time.Time               `bson:"requested_at"`
	CreatedDate          time.Time               `bson:"created_date"`
	CompletionDate       time.Time               `bson:"completion_date"`
	IsThreeDSecure       bool                    `bson:"is_three_d_secure"`
	ThreeDSecureURL      string                  `bson:"three_d_secure_url"`
	OrderId              int                     `bson:"order_id"`
	ErrorMessage         string                  `bson:"error_message"`
	PaymentRequestLogs   []*PaymentRequestLog    `bson:"payment_request_logs"`
}
type PaymentRequestLog struct {
	CreatedDate   time.Time           `bson:"created_date"`
	RequestJSON   string              `bson:"request_json"`
	ResponseJSON  string              `bson:"response_json"`
	ResponseCode  string              `bson:"response_code"`
	PaymentStatus enums.PaymentStatus `bson:"payment_status"`
}

type URLHolder struct {
	URL string `json:"url"`
}
type PaymentRequestDto struct {
	ChargeID        string  `json:"charge_id"`
	PaymentStatus   string  `json:"payment_status"`
	ThreeDSecureURL string  `json:"three_d_secure_url,omitempty"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	CustomerID      int     `json:"customer_id"`
	OrderID         *int    `json:"order_id,omitempty"`
}
type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     Phone  `json:"phone"`
}

type Phone struct {
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
}
