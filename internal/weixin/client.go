package weixin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	AppId  string
	Secret string

	ticker      *time.Ticker
	accessToken string
}

func NewClient(c *Config) (client *Client, err error) {
	client = &Client{
		AppId:  c.AppId,
		Secret: c.Secret,
	}
	go client.refreshAccessToken()
	return client, nil
}

func (c *Client) refreshAccessToken() {
	c.updateAccessToken()
	ticker := time.NewTicker(time.Hour)
	for range ticker.C {
		c.updateAccessToken()
	}
}

func (c *Client) updateAccessToken() {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", c.AppId, c.Secret)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	c.accessToken = result.AccessToken
}

func (c *Client) Close() error {
	if c.ticker != nil {
		c.ticker.Stop()
		c.ticker = nil
	}
	return nil
}
