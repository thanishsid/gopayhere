package gopayhere

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) RequestCharging(data ChargeRequestBody) (ChargeResponse, error) {
	var info ChargeResponse

	accessToken, err := c.getAccessToken()
	if err != nil {
		return info, err
	}

	req, err := createChargeRequest(c.BaseUrl, accessToken, data)
	if err != nil {
		return info, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return info, err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&info); err != nil {
		return info, err
	}

	if payhereErr := info.Error; payhereErr != nil && *payhereErr == "invalid_token" {

		if info.ErrorDescription != nil && strings.HasPrefix(*info.ErrorDescription, "Invalid access token") {
			return info, ErrInvalidAccessToken
		}

		if info.ErrorDescription != nil && strings.HasPrefix(*info.ErrorDescription, "Access token expired") {
			return info, ErrAccessTokenExpired
		}

		return info, ErrInvalidAccessToken
	}

	if info.Status == -2 {
		switch info.Message {
		case "Authentication error":
			return info, ErrUnauthorized
		}

		return info, ErrFailed
	}

	if info.Status == -1 {
		switch info.Message {
		case "Invalid token":
			return info, ErrInvalidCustomerToken
		case "Invalid currency":
			return info, ErrInvalidCurrency
		case "Invalid amount":
			return info, ErrInvalidAmount
		}

		return info, ErrCancelled
	}

	if info.Data == nil {
		return info, ErrNoData
	}

	return info, nil
}

// Payhere Charge Response Data
type ChargeResponseData struct {
	OrderID            string  `json:"order_id"`
	Items              string  `json:"items"`
	Currency           string  `json:"currency"`
	Amount             float64 `json:"amount"`
	Custom1            string  `json:"custom_1"`
	Custom2            string  `json:"custom_2"`
	PaymentID          uint    `json:"payment_id"`
	StatusCode         int     `json:"status_code"`
	StatusMessage      string  `json:"status_message"`
	Md5sig             string  `json:"md5sig"`
	AuthorizationToken *string `json:"authorization_token"`
}

// Payhere Charge Response Body
type ChargeResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"msg"`
	Data    *ChargeResponseData `json:"data"`

	// Not nil when payhere returns an error
	Error            *string `json:"error"`
	ErrorDescription *string `json:"error_description"`
}

type ChargeRequestBody struct {
	Type          string              `json:"type"`
	OrderID       string              `json:"order_id"`
	Items         string              `json:"items"`
	Currency      string              `json:"currency"`
	Amount        float64             `json:"amount"`
	CustomerToken string              `json:"customer_token"`
	Custom1       string              `json:"custom_1,omitempty"`
	Custom2       string              `json:"custom_2,omitempty"`
	NotifyUrl     string              `json:"notify_url"`
	ItemList      []ChargeRequestItem `json:"itemList"`
}

type ChargeRequestItem struct {
	Name       string  `json:"name"`
	Number     string  `json:"number"`
	Quantity   int     `json:"quantity"`
	UnitAmount float32 `json:"unit_amount"`
}

func createChargeRequest(baseUrl *url.URL, accessToken string, body ChargeRequestBody) (*http.Request, error) {
	chargeUrl := baseUrl.JoinPath("payment", "charge").String()

	payloadJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", chargeUrl, bytes.NewReader(payloadJson))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}
