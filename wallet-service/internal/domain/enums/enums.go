package enums

type TransactionType string

const (
	BalanceAddition TransactionType = "BalanceAddition"
	SendToUser      TransactionType = "SendToUser"
)

type IsolationLevel string

const (
	Unspecified     IsolationLevel = "unspecified"
	ReadUncommitted IsolationLevel = "read uncommitted"
	ReadCommitted   IsolationLevel = "read committed"
	RepeatableRead  IsolationLevel = "repeatable read"
	Serializable    IsolationLevel = "serializable"
)
