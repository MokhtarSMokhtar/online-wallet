package enums

type PaymentStatus string

type PaymentType string

type PaymentMethodType string

const (
	Initiated  PaymentStatus = "Initiated"
	Pending    PaymentStatus = "Pending"
	Authorized PaymentStatus = "Authorized"
	Timeout    PaymentStatus = "Timeout"
	Captured   PaymentStatus = "Captured"
	Failed     PaymentStatus = "Failed"
)

const (
	Credit PaymentMethodType = "Credit"
	Wallet PaymentMethodType = "Wallet"
)

const (
	Order        PaymentType = "Order"
	ChargeWallet PaymentType = "ChargeWallet"
)
