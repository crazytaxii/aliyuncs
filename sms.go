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
func SendSMS(phone string, signName string, templateParam map[string]string,
	templateCode string) (string, error) {
	jsonTemplateParam, err := json.Marshal(templateParam)
	if err != nil {
		return "", err
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
		"SignName":      signName,
		"TemplateCode":  templateCode,
		"TemplateParam": string(jsonTemplateParam),
	}
	sign := doSign("GET", data, config.AccessSecret) // 签名
	data["Signature"] = sign

	pList := make([]string, 0)
	for key, value := range data {
		pList = append(pList, fmt.Sprintf("%s=%s", key, specialUrlEncode(value)))
	}
	queryString := strings.Join(pList, "&")

	_, body, errList := gorequest.New().Get(fmt.Sprintf("http://%s/?%s", DOMAIN,
		queryString)).End()
	if errList != nil {
		return "", errList[0]
	}

	result := new(struct {
		Message   string `json:"Message"`
		Code      string `json:"Code"`
		RequestId string `json:"RequestId"`
		BizId     string `json:"BizId"`
	})
	err = json.Unmarshal([]byte(body), result)
	if err != nil {
		return "", err
	}
	if result.Code != "OK" {
		return "", fmt.Errorf("Send sms to user:%s failed, error code:%s, error message:%s",
			phone, result.Code, result.Message)
	}

	return result.BizId, nil
} // SendSMS()

/**
 * 发送批量短信接口
 */
/* func SendBatchSMS(phoneList []string, signName []string, templateCode string,
	templateParamList []map[string]string) (string, error) {
	jsonPhoneList, err := json.Marshal(phoneList)
	if err != nil {
		return "", nil
	}
	jsonSignName, err := json.Marshal(signName)
	if err != nil {
		return "", nil
	}
	jsonTemplateParamList, err := json.Marshal(templateParamList)
	if err != nil {
		return "", nil
	}

	// 参数
	data := map[string]string{
		"PhoneNumberJson":   string(jsonPhoneList),
		"SignNameJson":      string(jsonSignName),
		"TemplateCode":      templateCode,
		"TemplateParamJson": string(jsonTemplateParamList),
	}
} // SendBatchSMS() */
