package interfaces

import (
	"context"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/enums"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/models"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *models.PaymentRequest) error
	UpdatePaymentRequests(ctx context.Context, payment []*models.PaymentRequest) error
	UpdatePaymentRequest(ctx context.Context, payment *models.PaymentRequest) error
	GetPaymentByIdempotencyKey(ctx context.Context, key string) (*models.PaymentRequest, error)
	GetPaymentRequestByUserAndType(ctx context.Context, userId string, paymentType enums.PaymentType) ([]*models.PaymentRequest, error)
}

type PaymentService interface {
	CreateChargeRequest(ctx context.Context, payload models.PaymentRequestPayload, customerId string, paymentType enums.PaymentType) (*models.ChargeResponse, error)
	CapturePayment(ctx context.Context, payload models.ChargeResponse)
}
type TapClient interface {
	PostCharge(ctx context.Context, endpoint string, jObject models.PaymentRequestPayload, lang string) (*models.ChargeResponse, error)
}
