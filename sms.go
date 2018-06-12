package aliyuncs

import (
	"encoding/json"
	"fmt"
	"strings"
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

/**
 * 向指定手机号发送短信
 */
func SendSMS(phone string, param map[string]string, templateCode string) error {
	templateParam, err := json.Marshal(param)
	if err != nil {
		return err
	}

	// 参数
	data := map[string]string{
		"AccessKeyId":      config.AccessKeyId,
		"Timestamp":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"Format":           FORMAT,
		"SignatureMethod":  SIGNATURE_METHOD,
		"SignatureVersion": SIGNARURE_VERSION,
		"SignatureNonce":   randStr(32),

		"Action":        API_ACTION,
		"Version":       API_VERSION,
		"RegionId":      REGION_ID,
		"PhoneNumbers":  phone,
		"SignName":      config.SignName,
		"TemplateCode":  templateCode,
		"TemplateParam": string(templateParam),
	}
	sign := doSign("GET", data, config.AccessSecret) // 签名
	data["Signature"] = sign

	pList := make([]string, 0)
	for key, value := range data {
		pList = append(pList, fmt.Sprintf("%s=%s", key, specialUrlEncode(value)))
	}
	queryString := strings.Join(pList, "&")
	// url := fmt.Sprintf("http://%s/?%s", DOMAIN, queryString)

	_, body, errList := gorequest.New().Get(fmt.Sprintf("http://%s/?%s", DOMAIN,
		queryString)).End()
	if errList != nil {
		return errList[0]
	}

	result := new(struct {
		Message   string `json:"Message"`
		Code      string `json:"Code"`
		RequestId string `json:"RequestId"`
		BizId     string `json:"BizId"`
	})
	err = json.Unmarshal([]byte(body), result)
	if err != nil {
		return err
	}
	if result.Code != "OK" {
		return fmt.Errorf("Send sms to user:%s failed, error code:%s, error message:%s",
			phone, result.Code, result.Message)
	}

	return nil
} // SendSMS()

/* func Simulate() {
	param := &map[string]string{
		"AccessKeyId":      "testId",
		"Timestamp":        "2017-07-12T02:42:19Z",
		"Format":           "XML",
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureVersion": "1.0",
		"SignatureNonce":   "45e25e9b-0a6f-4070-8c85-2956eda1b466",

		"Action":        "SendSms",
		"Version":       "2017-05-25",
		"RegionId":      "cn-hangzhou",
		"PhoneNumbers":  "15300000001",
		"SignName":      "阿里云短信测试专用",
		"TemplateCode":  "SMS_71390007",
		"TemplateParam": "{\"customer\":\"test\"}",
		"OutId":         "123",
	}
	sign := doSign("GET", *param, "testSecret")
	(*param)["Signature"] = sign
	fmt.Println("sign:", sign)

	pList := make([]string, 0)
	for key, value := range *param {
		pList = append(pList, fmt.Sprintf("%s=%s", key, specialUrlEncode(value)))
	}
	queryString := strings.Join(pList, "&")
	url := fmt.Sprintf("http://%s/?%s", DOMAIN, queryString)
	fmt.Println("request_url:", url)
} */
