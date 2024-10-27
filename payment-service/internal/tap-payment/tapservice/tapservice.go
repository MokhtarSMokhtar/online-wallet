package services

import (
	"context"
	"fmt"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/enums"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/interfaces"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/models"
)

type paymentService struct {
	tapClient interfaces.TapClient
}

func NewPaymentService(client interfaces.TapClient) interfaces.PaymentService {
	return &paymentService{
		tapClient: client,
	}
}

func (s *paymentService) CreateChargeRequest(ctx context.Context, payload models.PaymentRequestPayload, customerId string, paymentType enums.PaymentType) (*models.ChargeResponse, error) {

	chargeResponse, err := s.tapClient.PostCharge(ctx, "charges", payload, "EN")
	if err != nil {
		return nil, fmt.Errorf("failed to create charge: %w", err)
	}
	return chargeResponse, nil
}

func (s *paymentService) CapturePayment(ctx context.Context, payload models.ChargeResponse) {
	//TODO implement me
	panic("implement me")
}
