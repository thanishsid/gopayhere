package gopayhere

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func getAccessToken(client *http.Client, payhereUrl *url.URL, appID, appSecret string) (*AccessTokenResponse, error) {
	var info AccessTokenResponse

	req, err := createAccessTokenRequest(payhereUrl, appID, appSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute access token request: %w", err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&info); err != nil {
		return nil, err
	}

	if payhereErr := info.Error; payhereErr != nil {
		if info.ErrorDescription != nil {
			return nil, fmt.Errorf("payhere error: %s; details: %s", *info.Error, *info.ErrorDescription)
		} else {
			return nil, fmt.Errorf("payhere error: %s", *info.Error)
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, ErrFailed
	}

	return &info, nil
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`

	// Not nil when payhere returns an error
	Error            *string `json:"error"`
	ErrorDescription *string `json:"error_description"`
}

func createAccessTokenRequest(payhereUrl *url.URL, appID, appSecret string) (*http.Request, error) {
	tokenUrl := payhereUrl.JoinPath("oauth", "token").String()

	payload := strings.NewReader("grant_type=client_credentials")

	req, err := http.NewRequest("POST", tokenUrl, payload)
	if err != nil {
		return nil, err
	}

	authCode := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", appID, appSecret)))

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", authCode))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}
