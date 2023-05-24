package gopayhere

import (
	"log"
	"net/http"
	"net/url"
)

func NewClient(payhereUrl, appID, appSecret string) (*Client, error) {
	baseUrl, err := url.Parse(payhereUrl)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}

	return &Client{
		BaseUrl:   baseUrl,
		AppID:     appID,
		AppSecret: appSecret,
		client:    httpClient,
	}, nil
}

type Client struct {
	BaseUrl   *url.URL
	AppID     string
	AppSecret string

	client *http.Client
}

func (c *Client) getAccessToken() (string, error) {
	res, err := getAccessToken(c.client, c.BaseUrl, c.AppID, c.AppSecret)
	if err != nil {
		log.Printf("failed to get access token: %s", err.Error())
		return "", err
	}

	return res.AccessToken, nil
}
