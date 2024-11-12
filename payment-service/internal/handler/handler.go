package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MokhtarSMokhtar/online-wallet/comman/jwt"
	"github.com/MokhtarSMokhtar/online-wallet/comman/middelwares"
	"github.com/MokhtarSMokhtar/online-wallet/comman/utile"
	walletpb "github.com/MokhtarSMokhtar/online-wallet/online-wallet-protos/github.com/MokhtarSMokhtar/online-wallet-protos/wallet"
	"github.com/MokhtarSMokhtar/online-wallet/payment-service/config"
	"github.com/MokhtarSMokhtar/online-wallet/payment-service/internal/enums"
	grpcclient "github.com/MokhtarSMokhtar/online-wallet/payment-service/internal/grpc"
	"github.com/MokhtarSMokhtar/online-wallet/payment-service/internal/interfaces"
	"github.com/MokhtarSMokhtar/online-wallet/payment-service/internal/models"
	"github.com/oklog/ulid/v2"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type PaymentHandler struct {
	paymentService    interfaces.PaymentService
	paymentRepository interfaces.PaymentRepository
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

func NewPaymentHandler(paymentService interfaces.PaymentService, repository interfaces.PaymentRepository) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService, paymentRepository: repository}
}

// UserPaymentHandler godoc
// @Summary      Process a user payment
// @Description  Processes a payment and returns the transaction details
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        request body models.CreateChargeRequestPayload true "Payment information"
// @Success      200  {object} models.ChargeResponse "Payment processed successfully"
// @Failure      400  {object}  ErrorResponse   "Bad Request"
// @Failure      500  {object}  ErrorResponse   "Internal Server Error"
// @Router       /payments [post]
// @Security     ApiKeyAuth
func (h *PaymentHandler) UserPaymentHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	//get user id from claims
	cli, ok := r.Context().Value(middelwares.UserContextKey).(*jwt.Claims)
	if !ok {
		http.Error(w, "Unable to retrieve user from context", http.StatusInternalServerError)
		return
	}
	//Serialize the req body
	var pymReq models.CreateChargeRequestPayload
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&pymReq); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			http.Error(w, fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset), http.StatusBadRequest)
		case errors.Is(err, io.EOF):
			http.Error(w, "Request body must not be empty", http.StatusBadRequest)
		case errors.As(err, &unmarshalTypeError):
			http.Error(w, fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset), http.StatusBadRequest)
		default:
			http.Error(w, "Unable to process request", http.StatusBadRequest)
		}
		log.Printf("Error decoding JSON: %v", err)
		return
	}
	//payment type
	paymentRequestModel := mapPaymentRequestModelToPaymentRequest(pymReq, "Charge Wallet ", cli)

	dto, err := h.paymentService.CreateChargeRequest(r.Context(), paymentRequestModel, cli.UserId, pymReq.PaymentType)
	if err != nil {
		http.Error(w, "Failed to create payment request", http.StatusInternalServerError)
		return
	}
	paymentRequest := &models.PaymentRequest{
		Id:             generateUlid(),
		RequestedAt:    time.Now().UTC(),
		PaymentMethod:  pymReq.PaymentMethod,
		PaymentType:    pymReq.PaymentType,
		PaymentStatus:  enums.Initiated,
		IsThreeDSecure: dto.ThreeDSecure,
		Amount:         dto.Amount,
		UserId:         cli.UserId,
		ChargeId:       dto.Id,
		IdempotencyKey: paymentRequestModel.Reference.Idempotent,
	}
	switch pymReq.PaymentType {
	case enums.ChargeWallet:
		h.HandelChargeWalletPaymentRequest(r.Context(), cli.UserId, pymReq.PaymentType)

	}
	createPaymentErr := h.paymentRepository.CreatePayment(r.Context(), paymentRequest)
	if createPaymentErr != nil {
		http.Error(w, "Failed to create payment request", http.StatusInternalServerError)
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto)
}

// CapturePayment godoc
// @Summary      Capture a payment
// @Description  Captures a payment after authorization
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        request body models.ChargeResponse true "Payment capture information"
// @Success      200  "Payment captured successfully"
// @Failure      400  {object}  ErrorResponse   "Bad Request"
// @Failure      500  {object}  ErrorResponse   "Internal Server Error"
// @Router       /payments/capture [post]
func (h *PaymentHandler) CapturePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	var payRes models.ChargeResponse
	err := DecodeRequest(&payRes, r)
	if err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
	}
	//update the payment request
	allowedPaymentStatuses := []models.TapChargeStatus{
		models.TapChargeStatusPending,
		models.TapChargeStatusAuthorized,
		models.TapChargeStatusInitiated,
	}

	isCon := utile.ContainsValue(allowedPaymentStatuses, payRes.Status)
	if !isCon {
		http.Error(w, "Payment is not allowed", http.StatusForbidden)
		//TODO LOG
		//"Error in capturing payment, the payment request for Invoice id: {trackId} has already been captured",
	}
	//capturedStatuses
	paymentRequest, err := h.paymentRepository.GetPaymentByIdempotencyKey(r.Context(), payRes.Reference.Idempotent)
	if err != nil {
		http.Error(w, "Failed to get payment by idempotency key", http.StatusInternalServerError)
	}
	capturedStatuses := []models.TapChargeStatus{
		models.TapChargeStatusApproved,
		models.TapChargeStatusCaptured,
		models.TapChargeStatusAuthorized,
	}
	isCapturedPay := utile.ContainsValue(capturedStatuses, payRes.Status)
	if !isCapturedPay {
		///
		http.Error(w, "Payment is not captured", http.StatusForbidden)
		//TODO log
		//Update the payment request
		err := h.handelFailedPaymentReq(r.Context(), paymentRequest, payRes)
		if err != nil {
			http.Error(w, "Error while update Payment", http.StatusBadRequest)
		}
	}
	//handel success payments
	switch paymentRequest.PaymentType {
	case enums.ChargeWallet:
		HandelSuccessWalletChargePayment(paymentRequest, payRes)
	}

}

func HandelSuccessWalletChargePayment(request *models.PaymentRequest, res models.ChargeResponse) {
	//TODO Log
	userId, err := strconv.Atoi(request.UserId)
	if err != nil {
		log.Printf("Error converting UserId to integer: %v", err)
		return
	}
	walletClient := grpcclient.GetWalletServiceClient()
	amount := float32(request.Amount)
	updateReq := &walletpb.UpdateWalletRequest{
		UserId: int32(userId), // Correctly set the UserId
		Amount: amount,
		Reason: "Charge Wallet",
	}
	// Make the gRPC call
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := walletClient.UpdateWallet(ctx, updateReq)
	if err != nil {
		log.Printf("Failed to update wallet: %v", err)
		return
	}

	if !resp.Success {
		log.Printf("Wallet update failed: %s", resp.Message)
		return
	}

	log.Printf("Wallet updated successfully for user %s", request.UserId)
}

func (h *PaymentHandler) handelFailedPaymentReq(ctx context.Context, paymentRe *models.PaymentRequest, res models.ChargeResponse) error {
	paymentRe.PaymentStatus = enums.Failed
	paymentRe.CompletionDate = time.Now().UTC()
	paymentRe.ErrorMessage = res.Response.Message

	err := h.paymentRepository.UpdatePaymentRequest(ctx, paymentRe)
	if err != nil {
		return err
	}
	return nil
}

func (s *PaymentHandler) HandelChargeWalletPaymentRequest(ctx context.Context, customerId string, paymentType enums.PaymentType) {
	var err error
	payments, err := s.paymentRepository.GetPaymentRequestByUserAndType(ctx, customerId, paymentType)
	if err != nil {
		_ = fmt.Errorf("failed to get charge request by user: %w", err)
	}
	for _, payment := range payments {
		payment.PaymentStatus = enums.Timeout
	}
	err = s.paymentRepository.UpdatePaymentRequests(ctx, payments)
	if err != nil {
		_ = fmt.Errorf("failed to update charge requests: %w", err)
	}

}

func mapPaymentRequestModelToPaymentRequest(requestModel models.CreateChargeRequestPayload, paymentDisc string, claims *jwt.Claims) models.PaymentRequestPayload {
	// Map PaymentRequestModel to the structure expected by Tap Payments API
	conf := config.NewConfig()
	return models.PaymentRequestPayload{
		Amount:            requestModel.Amount,
		Currency:          "EGP",
		CustomerInitiated: true,
		ThreeDSecure:      true,
		SaveCard:          false,
		Description:       paymentDisc,
		Metadata: models.Metadata{
			Udf1: "dss",
		},
		Customer: models.Customer{
			FirstName: claims.Name,
			LastName:  "-",
			Email:     claims.Email,
			Phone: models.Phone{
				CountryCode: "20",
				Number:      claims.Phone,
			},
		},
		Reference: models.Reference{
			Transaction: generateUlid(),
			Order:       requestModel.OrderId,
			Idempotent:  generateUlid(),
		},
		Receipt: models.Receipt{
			Email: false,
			SMS:   false,
		},
		Merchant: models.Merchant{
			ID: "1234",
		},
		Source: models.Source{
			ID: "src_all",
		},
		Post: models.URLHolder{
			URL: conf.BaseURl + conf.PostUrl,
		},
		Redirect: models.URLHolder{
			URL: "http://your_website.com/redirect_url",
		},
	}
}

func generateUlid() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

func DecodeRequest[T any](reType T, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reType); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		case errors.Is(err, io.EOF):
			return fmt.Errorf("request body must not be empty")
		case errors.As(err, &unmarshalTypeError):
			return fmt.Errorf("request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		default:
			return fmt.Errorf("Unable to process request")
		}
	}
	return nil
}
