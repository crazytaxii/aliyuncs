package aliyuncs

import ()

const DOMAIN = "dysmsapi.aliyuncs.com"

type Client struct {
	AccessKeyId  string
	AccessSecret string
}

func NewClient(accessKeyId string, accessSecret string) *Client {
	return &Client{
		AccessKeyId:  accessKeyId,
		AccessSecret: accessSecret,
	}
} // NewClient()
