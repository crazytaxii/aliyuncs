package aliyuncs_test

import (
	"fmt"
	"testing"

	cs "github.com/crazytaxii/aliyuncs"
)

const (
	ACCESS_ID     = "access_key_id"
	ACCESS_SECRET = "access_key_secret"
)

func TestInit(t *testing.T) {
	cs.Init(
		ACCESS_ID,
		ACCESS_SECRET,
	)
}

func TestSendSMS(t *testing.T) {
	param := map[string]string{
		"code": "123456",
	}

	_, err := cs.SendSMS(
		"13800000000",
		"sign_name",
		param,
		"SMS_123456789",
	)
	if err != nil {
		fmt.Println("err:", err)
	}
}
