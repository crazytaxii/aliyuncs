package aliyuncs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/parnurzeal/gorequest"
)

const (
	FORMAT            = "JSON"
	SIGNATURE_METHOD  = "HMAC-SHA1"
	SIGNARURE_VERSION = "1.0"
	API_ACTION        = "SendSms"
	API_VERSION       = "2017-05-25"
	REGION_ID         = "cn-hangzhou"
)

type Response struct {
	Message   string `json:"Message"`
	Code      string `json:"Code"`
	RequestId string `json:"RequestId"`
	BizId     string `json:"BizId"`
}

/**
 * 向指定手机号发送短信
 */
func (client *Client) SendSMS(phone string, signName string, templateCode string, templateParam map[string]string) (string, error) {
	jsonTemplateParam, err := json.Marshal(templateParam)
	if err != nil {
		return "", err
	}

	// 参数
	data := map[string]string{
		"AccessKeyId":      client.AccessKeyId,
		"Timestamp":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"Format":           FORMAT,
		"SignatureMethod":  SIGNATURE_METHOD,
		"SignatureVersion": SIGNARURE_VERSION,
		"SignatureNonce":   randStr(32),

		"Action":        API_ACTION,
		"Version":       API_VERSION,
		"RegionId":      REGION_ID,
		"PhoneNumbers":  phone,
		"SignName":      signName,
		"TemplateCode":  templateCode,
		"TemplateParam": string(jsonTemplateParam),
	}
	data["Signature"] = doSign("GET", data, client.AccessSecret) // 签名

	pList := make([]string, 0)
	for key, value := range data {
		pList = append(pList, fmt.Sprintf("%s=%s", key, specialUrlEncode(value)))
	}
	// queryString := strings.Join(pList, "&")

	// _, body, errList := gorequest.New().Get(fmt.Sprintf("http://%s/?%s", DOMAIN,
	// 	queryString)).End()
	req := gorequest.New().Get(fmt.Sprintf("http://%s", DOMAIN))
	for _, param := range pList {
		req.Query(param)
	}
	_, body, errList := req.End()
	if errList != nil {
		return "", errList[0]
	}

	resp := new(Response)
	err = json.Unmarshal([]byte(body), resp)
	if err != nil {
		return "", err
	}
	if resp.Code != "OK" {
		return "", fmt.Errorf("Send sms to user: %s failed, err_code: %s, err_msg: %s", phone, resp.Code, resp.Message)
	}

	return resp.BizId, nil
} // SendSMS()
