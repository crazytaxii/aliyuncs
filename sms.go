package aliyuncs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/parnurzeal/gorequest"
)

/**
 * 向指定手机号发送短信
 */
func (c *Client) SendSMS(phone, signName, templateCode string, templateParam map[string]string) (string, error) {
	jsonTemplateParam, err := json.Marshal(templateParam)
	if err != nil {
		return "", err
	}

	// 参数
	data := map[string]string{
		"AccessKeyId":      c.AccessKeyId,
		"Timestamp":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"Format":           FORMAT,
		"SignatureMethod":  SIGNATURE_METHOD,
		"SignatureVersion": SIGNARURE_VERSION,
		"SignatureNonce":   randStr(32),

		"Action":        SMS_API_ACTION,
		"Version":       API_VERSION,
		"RegionId":      REGION_ID,
		"PhoneNumbers":  phone,
		"SignName":      signName,
		"TemplateCode":  templateCode,
		"TemplateParam": string(jsonTemplateParam),
	}
	data["Signature"] = c.doSign("GET", data) // 签名

	pList := make([]string, 0)
	for key, value := range data {
		pList = append(pList, fmt.Sprintf("%s=%s", key, specialUrlEncode(value)))
	}

	req := gorequest.New().Get("http://dysmsapi.aliyuncs.com")
	for _, param := range pList {
		req.Query(param)
	}
	_, body, errList := req.End()
	if errList != nil {
		return "", errList[0]
	}

	resp := &SendSMSResponse{}
	err = json.Unmarshal([]byte(body), resp)
	if err != nil {
		return "", err
	}
	if resp.Code != "OK" {
		return "", fmt.Errorf("Send sms to user: %s failed, err_code: %s, err_msg: %s", phone, resp.Code, resp.Message)
	}

	return resp.BizId, nil
} // SendSMS()
