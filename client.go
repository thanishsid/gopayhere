package gopayhere

import (
	"net/http"
	"net/url"
	"sync"
	"time"
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

	accessToken       string
	accessTokenExpiry time.Time
	tokenRefreshing   bool

	client *http.Client

	mu sync.Mutex
}

func (c *Client) getAccessToken() string {
	go func() {
		if c.tokenRefreshing {
			return
		}

		if c.accessTokenExpiry.After(time.Now().Add(time.Second * 30)) {
			return
		}

		c.mu.Lock()
		defer c.mu.Unlock()

		c.tokenRefreshing = true

		res, err := getAccessToken(c.client, c.BaseUrl, c.AppID, c.AppSecret)
		if err != nil {
			return
		}

		c.accessToken = res.AccessToken
		c.tokenRefreshing = false
		c.accessTokenExpiry = time.Now().Add(time.Second * time.Duration(res.ExpiresIn))
	}()

	return c.accessToken
}
