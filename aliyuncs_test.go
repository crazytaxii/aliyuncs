package aliyuncs_test

import (
	"fmt"
	"testing"

	cs "github.com/crazytaxii/aliyuncs"
)

const (
	ACCESS_ID     = "access_key_id"
	ACCESS_SECRET = "access_key_secret"
	SIGN_NAME     = "sign_name"
)

func TestInit(t *testing.T) {
	cs.Init(
		ACCESS_ID,
		ACCESS_SECRET,
		SIGN_NAME,
	)
}

func TestSendSMS(t *testing.T) {
	param := map[string]string{
		"code": "123456",
	}

	err := cs.SendSMS(
		"13800000000",
		param,
		"SMS_123456789",
	)
	if err != nil {
		fmt.Println("err:", err)
	}
}
