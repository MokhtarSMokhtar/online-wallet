package models

import "github.com/MokhtarSMokhtar/online-wallet/payment-service/internal/enums"

type PaymentRequestModel struct {
	AmountToPay              float64                 `json:"amountToPay"`
	PaymentMerchant          string                  `json:"paymentMerchant"`
	CurrencyISO              string                  `json:"currencyIso"`
	PaymentDescription       string                  `json:"paymentDescription"`
	TransactionID            string                  `json:"transactionId"`
	OrderID                  string                  `json:"orderId"`
	SaveCardInformation      bool                    `json:"saveCardInformation"`
	SendEmail                bool                    `json:"sendEmail"`
	SendSMS                  bool                    `json:"sendSms"`
	CustomerID               string                  `json:"customerId"`
	CustomerFirstName        string                  `json:"customerFirstName"`
	CustomerMiddleName       string                  `json:"customerMiddleName"`
	CustomerLastName         string                  `json:"customerLastName"`
	CustomerEmail            string                  `json:"customerEmail"`
	CustomerPhoneCountryCode string                  `json:"customerPhoneCountryCode"`
	CustomerPhoneNumber      string                  `json:"customerPhoneNumber"`
	PaymentSource            string                  `json:"paymentSource"`
	Language                 string                  `json:"language"`
	IdempotencyKey           string                  `json:"idempotencyKey"`
	PaymentMethod            enums.PaymentMethodType `json:"paymentMethod"`
}

type CreateChargeRequestPayload struct {
	Amount        float64                 `json:"amount"`
	OrderId       string                  `json:"orderId"`
	PaymentMethod enums.PaymentMethodType `json:"payment_method"`
	PaymentType   enums.PaymentType       `json:"payment_type"`
}
