package enums

type TransactionType string

const (
	BalanceAddition TransactionType = "BalanceAddition"
	SendToUser      TransactionType = "SendToUser"
)
