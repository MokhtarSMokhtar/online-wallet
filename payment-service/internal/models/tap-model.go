package models

// TapChargeStatus represents the status of a charge
type TapChargeStatus string

const (
	TapChargeStatusInserted          TapChargeStatus = "INSERTED"
	TapChargeStatusInitiated         TapChargeStatus = "INITIATED"
	TapChargeStatusAbandoned         TapChargeStatus = "ABANDONED"
	TapChargeStatusCancelled         TapChargeStatus = "CANCELLED"
	TapChargeStatusFailed            TapChargeStatus = "FAILED"
	TapChargeStatusDeclined          TapChargeStatus = "DECLINED"
	TapChargeStatusRestricted        TapChargeStatus = "RESTRICTED"
	TapChargeStatusCaptured          TapChargeStatus = "CAPTURED"
	TapChargeStatusVoid              TapChargeStatus = "VOID"
	TapChargeStatusTimedout          TapChargeStatus = "TIMEDOUT"
	TapChargeStatusUnknown           TapChargeStatus = "UNKNOWN"
	TapChargeStatusApproved          TapChargeStatus = "APPROVED"
	TapChargeStatusPending           TapChargeStatus = "PENDING"
	TapChargeStatusAuthorized        TapChargeStatus = "AUTHORIZED"
	TapChargeStatusFailedSuccess     TapChargeStatus = "FAILED_SUCCESS"
	TapChargeStatusFailedValidated   TapChargeStatus = "FAILED_VALIDATED"
	TapChargeStatusInvalidResponse   TapChargeStatus = "INVALID_RESPONSE"
	TapChargeStatusInsufficientFunds TapChargeStatus = "INSUFFICIENT_FUNDS"
)

type PaymentRequestPayload struct {
	Amount            float64   `json:"amount"`
	Currency          string    `json:"currency"`
	CustomerInitiated bool      `json:"customer_initiated"`
	ThreeDSecure      bool      `json:"threeDSecure"`
	SaveCard          bool      `json:"save_card"`
	Description       string    `json:"description"`
	Metadata          Metadata  `json:"metadata"`
	Reference         Reference `json:"reference"`
	Receipt           Receipt   `json:"receipt"`
	Customer          Customer  `json:"customer"`
	Merchant          Merchant  `json:"merchant"`
	Source            Source    `json:"source"`
	Post              URLHolder `json:"post"`
	Redirect          URLHolder `json:"redirect"`
}

type ChargeResponse struct {
	Id           string          `json:"id"`
	Status       TapChargeStatus `json:"status"`
	Amount       float64         `json:"amount"`
	Currency     string          `json:"currency"`
	ThreeDSecure bool            `json:"threeDSecure"`
	Description  string          `json:"description"`
	Transaction  Transaction     `json:"transaction"`
	Customer     Customer        `json:"customer"`
	Response     Response        `json:"response"`
	Redirect     Redirect        `json:"redirect"`
	Reference    Reference       `json:"reference"`
	Receipt      Receipt         `json:"receipt"`
}

type Transaction struct {
	Timezone     string  `json:"timezone"`
	Created      string  `json:"created"`
	URL          string  `json:"url"`
	Expiry       Expiry  `json:"expiry"`
	Asynchronous bool    `json:"asynchronous"`
	Amount       float64 `json:"amount"`
	Currency     string  `json:"currency"`
}

type Expiry struct {
	Period int    `json:"period"`
	Type   string `json:"type"`
}

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type Redirect struct {
	Status string `json:"status"`
	URL    string `json:"url"`
}
type Receipt struct {
	Email bool `json:"email"`
	SMS   bool `json:"sms"`
}

type Reference struct {
	Transaction string `json:"transaction"`
	Order       string `json:"order"`
	Idempotent  string `bson:"idempotent"`
}

type Source struct {
	Object string `json:"object"`
	ID     string `json:"id"`
	OnFile bool   `json:"on_file"`
}

type Metadata struct {
	Udf1 string `json:"udf1"`
}
type Merchant struct {
	ID string `json:"id"`
}
