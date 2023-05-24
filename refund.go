package gopayhere

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) RequestRefund(data RefundRequestBody) (RefundResponse, error) {
	var info RefundResponse

	accessToken, err := c.getAccessToken()
	if err != nil {
		return info, err
	}

	req, err := createRefundRequest(c.BaseUrl, accessToken, data)
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

	if info.Status == 1 {
		return info, nil
	}

	return info, errors.New(info.Message)
}

type RefundRequestBody struct {
	PaymentID          uint `json:"payment_id,omitempty"`
	Description        string
	AuthorizationToken string `json:"authorization_token,omitempty"`
}

type RefundResponse struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
	Data    any    `json:"data"`

	// Not nil when payhere returns an error
	Error            *string `json:"error"`
	ErrorDescription *string `json:"error_description"`
}

func createRefundRequest(baseUrl *url.URL, accessToken string, body RefundRequestBody) (*http.Request, error) {
	refundUrl := baseUrl.JoinPath("payment", "refund").String()

	payloadJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", refundUrl, bytes.NewReader(payloadJson))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}
