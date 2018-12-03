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

func TestSendSMS(t *testing.T) {
	param := map[string]string{
		"code": "123456",
	}

	cs.NewClient(ACCESS_ID, ACCESS_SECRET)
	bizId, err := cs.GetClient().SendSMS(
		"13800000000",
		"sign_name",
		"SMS_123456789",
		param,
	)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("biz_id: %s\n", bizId)
}
