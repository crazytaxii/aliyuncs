package aliyuncs

import (
	"net/http"
	"time"
)

type Client struct {
	accessKeyID  string
	accessSecret string
	*http.Client
}

func NewClient(keyID, secret string, timeout time.Duration) *Client {
	return &Client{
		accessKeyID:  keyID,
		accessSecret: secret,
		Client: &http.Client{
			Timeout: timeout,
		},
	}
}
