package aliyuncs

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	BaseURL = "http://dysmsapi.aliyuncs.com"

	Format           = "JSON"
	SignatureMethod  = "HMAC-SHA1"
	SignatureVersion = "1.0"
	APIVersion       = "2017-05-25"
	RegionID         = "cn-hangzhou"

	ActionSMS      = "SendSms"
	ActionBatchSMS = "SendBatchSms"
	ActionCall     = "SingleCallByTts"
)

func (c *Client) genReqParams(ctx CSContext, action string) (map[string]string, error) {
	p := map[string]string{
		"AccessKeyId":      c.accessKeyID,
		"Timestamp":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"Format":           Format,
		"SignatureMethod":  SignatureMethod,
		"SignatureVersion": SignatureVersion,
		"SignatureNonce":   randStr(32),
		"Action":           action,
		"Version":          APIVersion,
		"RegionId":         APIVersion,
	}
	extras, err := ctx.genReqParams()
	if err != nil {
		return nil, err
	}
	// merge 2 map
	for k, v := range extras {
		p[k] = v
	}

	p["Signature"] = c.doSign("GET", p)
	return p, nil
}

type (
	ResponseMessage interface {
		error(CSContext) error
	}
	CSContext interface {
		getPhone() string
		genReqParams() (map[string]string, error)
	}
)

type CSTemplate map[string]string

func (t CSTemplate) Marshal() (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type commonResponse struct {
	Message   string `json:"Message"`
	Code      string `json:"Code"`
	RequestID string `json:"RequestID"`
}

const (
	StatusOK = "OK"
)

func (cr commonResponse) error(ctx CSContext) error {
	if cr.Code != StatusOK {
		return fmt.Errorf("Send SMS to phone: %s failed, error code: %s, error message: %s",
			ctx.getPhone(), cr.Code, cr.Message)
	}
	return nil
}
