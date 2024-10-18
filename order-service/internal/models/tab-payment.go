package models

import (
	"time"
)

type ChargeResponse struct {
	Status       TapChargeStatus      `json:"status"`
	StatusString string               `json:"status_string"`
	ThreeDSecure bool                 `json:"three_d_secure"`
	RedirectUrl  string               `json:"redirect_url"`
	ChargeId     string               `json:"charge_id"`
	Message      string               `json:"message"`
	CustomerId   string               `json:"customer_id"`
	Code         string               `json:"code"`
	Response     TapApiChargeResponse `json:"response"`
}

type TapApiChargesListResponse struct {
	ObjectType string                 `json:"object_type"`
	LiveMode   bool                   `json:"live_mode"`
	Count      int                    `json:"count"`
	HasMore    int                    `json:"has_more"`
	ApiVersion string                 `json:"api_version"`
	Charges    []TapApiChargeResponse `json:"charges"`
}

type TapApiChargeResponse struct {
	Object              string      `json:"object"`
	LiveMode            bool        `json:"live_mode"`
	ApiVersion          string      `json:"api_version"`
	Id                  string      `json:"id"`
	Status              string      `json:"status"`
	Amount              int64       `json:"amount"`
	Currency            string      `json:"currency"`
	ThreeDSecure        bool        `json:"three_d_secure"`
	SaveCard            bool        `json:"save_card"`
	Description         string      `json:"description"`
	StatementDescriptor string      `json:"statement_descriptor"`
	Metadata            Metadata    `json:"metadata"`
	Transaction         Transaction `json:"transaction"`
	Reference           Reference   `json:"reference"`
	Response            Response    `json:"response"`
	Receipt             Receipt     `json:"receipt"`
	Customer            Customer    `json:"customer"`
	Source              Source      `json:"source"`
	Post                Post        `json:"post"`
	Redirect            Post        `json:"redirect"`
}

type Customer struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      Phone  `json:"phone"`
	Id         string `json:"id"`
}

type Phone struct {
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
}

type Metadata struct {
	Udf1 string `json:"udf1"`
}

type Post struct {
	Url string `json:"url"`
}

type Receipt struct {
	Id    string `json:"id"`
	Email bool   `json:"email"`
	Sms   bool   `json:"sms"`
}

type Reference struct {
	Track       string `json:"track"`
	Payment     string `json:"payment"`
	Gateway     string `json:"gateway"`
	Acquirer    string `json:"acquirer"`
	Transaction string `json:"transaction"`
	Order       string `json:"order"`
}

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Source struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

type Transaction struct {
	AuthorizationId string    `json:"authorization_id"`
	TimeZone        string    `json:"timezone"`
	Created         time.Time `json:"created"`
	Url             string    `json:"url"`
}

type TapChargeStatus int

const (
	Inserted TapChargeStatus = iota
	Initiated
	Abandoned
	Cancelled
	Failed
	Declined
	Restricted
	Captured
	Void
	Timedout
	Unknown
	Approved
	Pending
	Authorized
	FailedSuccess
	FailedValidated
	InvalidResponse
	InsufficientFunds
)

type ChargeDateType int

const (
	TransactionDate ChargeDateType = iota + 1
	PostDate
	BalanceAvailableDate
	PayoutBasedOnPeriodType
)
