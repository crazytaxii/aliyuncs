package aliyuncs

const (
	FORMAT            = "JSON"
	SIGNATURE_METHOD  = "HMAC-SHA1"
	SIGNARURE_VERSION = "1.0"
	API_VERSION       = "2017-05-25"
	REGION_ID         = "cn-hangzhou"

	SMS_API_ACTION       = "SendSms"
	BATCH_SMS_API_ACTION = "SendBatchSms"
	CALL_API_ACTION      = "SingleCallByTts"
)

type (
	SendSMSResponse struct {
		Message   string `json:"Message"`
		Code      string `json:"Code"`
		RequestId string `json:"RequestId"`
		BizId     string `json:"BizId"`
	}

	SingleCallByTTSResponse struct {
		Message   string `json:"Message"`
		RequestId string `json:"RequestId"`
		Code      string `json:"Code"`
		CallId    string `json:"CallId"`
	}
)
