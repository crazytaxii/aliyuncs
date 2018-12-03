package aliyuncs

const DOMAIN = "dysmsapi.aliyuncs.com"

type Client struct {
	AccessKeyId  string
	AccessSecret string
}

var client *Client

func NewClient(accessKeyId string, accessSecret string) {
	client = &Client{
		AccessKeyId:  accessKeyId,
		AccessSecret: accessSecret,
	}
} // NewClient()

func GetClient() *Client {
	return client
} // GetClient()
