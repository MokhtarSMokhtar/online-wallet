package tapclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/interfaces"
	"github.com/MokhatrSMokhtar/online-wallet/payment-service/internal/models"
	"io"
	"net/http"
)

// HttpClientFactory handles HTTP requests to the Tap Payments API
type httpClientFactory struct {
	httpClient *http.Client
	baseURL    string
	authToken  string
}

// NewHttpClientFactory initializes a new HttpClientFactory instance
func NewHttpClientFactory(authToken string, apiVersion string) interfaces.TapClient {
	if apiVersion == "" {
		apiVersion = "v2"
	}
	baseURL := fmt.Sprintf("https://api.tap.company/%s/", apiVersion)

	client := &http.Client{}

	return &httpClientFactory{
		httpClient: client,
		baseURL:    baseURL,
		authToken:  authToken,
	}
}

func (h *httpClientFactory) PostCharge(ctx context.Context, endpoint string, jObject models.PaymentRequestPayload, lang string) (*models.ChargeResponse, error) {
	if lang == "" {
		lang = "EN"
	}
	fullURL := h.baseURL + endpoint

	jsonBytes, err := json.Marshal(jObject)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fullURL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+h.authToken)
	req.Header.Set("lang_code", lang)
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Response Status: %d\nResponse Body: %s\n", resp.StatusCode, string(bodyBytes))

	isSuccess := resp.StatusCode >= 200 && resp.StatusCode < 300
	if isSuccess {
		var successResponse models.ChargeResponse
		if err := json.Unmarshal(bodyBytes, &successResponse); err != nil {
			return nil, fmt.Errorf("failed to unmarshal success response: %w", err)
		}
		return &successResponse, nil
	} else {
		// Define a struct to capture error details from the API
		type TapApiErrorResponse struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			// Add other fields if the API provides more details
		}

		var errorResp TapApiErrorResponse
		if err := json.Unmarshal(bodyBytes, &errorResp); err != nil {
			// If unable to unmarshal, return the raw response
			return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(bodyBytes))
		}
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, errorResp.Message)
	}
}
