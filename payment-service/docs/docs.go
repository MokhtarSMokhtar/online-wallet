// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/payments": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Processes a payment and returns the transaction details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payments"
                ],
                "summary": "Process a user payment",
                "parameters": [
                    {
                        "description": "Payment information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateChargeRequestPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Payment processed successfully",
                        "schema": {
                            "$ref": "#/definitions/models.ChargeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/payments/capture": {
            "post": {
                "description": "Captures a payment after authorization",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payments"
                ],
                "summary": "Capture a payment",
                "parameters": [
                    {
                        "description": "Payment capture information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ChargeResponse"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Payment captured successfully"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "enums.PaymentMethodType": {
            "type": "string",
            "enum": [
                "Credit",
                "Wallet"
            ],
            "x-enum-varnames": [
                "Credit",
                "Wallet"
            ]
        },
        "enums.PaymentType": {
            "type": "string",
            "enum": [
                "Order",
                "ChargeWallet"
            ],
            "x-enum-varnames": [
                "Order",
                "ChargeWallet"
            ]
        },
        "handler.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.ChargeResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "currency": {
                    "type": "string"
                },
                "customer": {
                    "$ref": "#/definitions/models.Customer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "receipt": {
                    "$ref": "#/definitions/models.Receipt"
                },
                "redirect": {
                    "$ref": "#/definitions/models.Redirect"
                },
                "reference": {
                    "$ref": "#/definitions/models.Reference"
                },
                "response": {
                    "$ref": "#/definitions/models.Response"
                },
                "status": {
                    "$ref": "#/definitions/models.TapChargeStatus"
                },
                "threeDSecure": {
                    "type": "boolean"
                },
                "transaction": {
                    "$ref": "#/definitions/models.Transaction"
                }
            }
        },
        "models.CreateChargeRequestPayload": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "orderId": {
                    "type": "string"
                },
                "payment_method": {
                    "$ref": "#/definitions/enums.PaymentMethodType"
                },
                "payment_type": {
                    "$ref": "#/definitions/enums.PaymentType"
                }
            }
        },
        "models.Customer": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone": {
                    "$ref": "#/definitions/models.Phone"
                }
            }
        },
        "models.Expiry": {
            "type": "object",
            "properties": {
                "period": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.Phone": {
            "type": "object",
            "properties": {
                "country_code": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                }
            }
        },
        "models.Receipt": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "boolean"
                },
                "sms": {
                    "type": "boolean"
                }
            }
        },
        "models.Redirect": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.Reference": {
            "type": "object",
            "properties": {
                "idempotent": {
                    "type": "string"
                },
                "order": {
                    "type": "string"
                },
                "transaction": {
                    "type": "string"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.TapChargeStatus": {
            "type": "string",
            "enum": [
                "INSERTED",
                "INITIATED",
                "ABANDONED",
                "CANCELLED",
                "FAILED",
                "DECLINED",
                "RESTRICTED",
                "CAPTURED",
                "VOID",
                "TIMEDOUT",
                "UNKNOWN",
                "APPROVED",
                "PENDING",
                "AUTHORIZED",
                "FAILED_SUCCESS",
                "FAILED_VALIDATED",
                "INVALID_RESPONSE",
                "INSUFFICIENT_FUNDS"
            ],
            "x-enum-varnames": [
                "TapChargeStatusInserted",
                "TapChargeStatusInitiated",
                "TapChargeStatusAbandoned",
                "TapChargeStatusCancelled",
                "TapChargeStatusFailed",
                "TapChargeStatusDeclined",
                "TapChargeStatusRestricted",
                "TapChargeStatusCaptured",
                "TapChargeStatusVoid",
                "TapChargeStatusTimedout",
                "TapChargeStatusUnknown",
                "TapChargeStatusApproved",
                "TapChargeStatusPending",
                "TapChargeStatusAuthorized",
                "TapChargeStatusFailedSuccess",
                "TapChargeStatusFailedValidated",
                "TapChargeStatusInvalidResponse",
                "TapChargeStatusInsufficientFunds"
            ]
        },
        "models.Transaction": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "asynchronous": {
                    "type": "boolean"
                },
                "created": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "expiry": {
                    "$ref": "#/definitions/models.Expiry"
                },
                "timezone": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}