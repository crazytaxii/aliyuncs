package aliyuncs

import (
	"fmt"
	"testing"
	"time"
)

const (
	TestAccessID     = "access_key_id"
	TestAccessSecret = "access_key_secret"
)

func TestSendSMS(t *testing.T) {
	c := NewClient(TestAccessID, TestAccessSecret, 10*time.Second)
	tmpl := CSTemplate{
		"code": "123456",
	}
	bizID, err := c.SendSMS("13712341234", "sign_name", "SMS_123456789", tmpl)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("biz ID: %s\n", bizID)
}
