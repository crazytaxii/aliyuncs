package aliyuncs

type Client struct {
	AccessKeyId  string
	AccessSecret string
}

var client *Client

func NewClient(accessKeyId string, accessSecret string) *Client {
	client = &Client{
		AccessKeyId:  accessKeyId,
		AccessSecret: accessSecret,
	}
	return client
} // NewClient()

func GetClient() *Client {
	return client
} // GetClient()
