package aliyuncs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/parnurzeal/gorequest"
)

func (client *Client) CallByTts(phone string, show string, ttsCode string, ttsParam map[string]string) (string, error) {
	jsonTtsParam, err := json.Marshal(ttsParam)
	if err != nil {
		return "", err
	}

	data := map[string]string{
		"AccessKeyId":      client.AccessKeyId,
		"Timestamp":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"Format":           FORMAT,
		"SignatureMethod":  SIGNATURE_METHOD,
		"SignatureVersion": SIGNARURE_VERSION,
		"SignatureNonce":   randStr(32),

		"Action":           CALL_API_ACTION,
		"Version":          API_VERSION,
		"RegionId":         REGION_ID,
		"CalledNumber":     phone,
		"CalledShowNumber": show,
		"TtsCode":          ttsCode,
		"TtsParam":         string(jsonTtsParam),
	}
	data["Signature"] = doSign("GET", data, client.AccessSecret) // 签名

	pList := make([]string, 0)
	for key, value := range data {
		pList = append(pList, fmt.Sprintf("%s=%s", key, specialUrlEncode(value)))
	}

	req := gorequest.New().Get("http://dyvmsapi.aliyuncs.com")
	for _, param := range pList {
		req.Query(param)
	}
	_, body, errList := req.End()
	if errList != nil {
		return "", errList[0]
	}
	fmt.Println("resp_body:", body)
	resp := new(SingleCallByTtsResponse)
	err = json.Unmarshal([]byte(body), resp)
	if err != nil {
		return "", err
	}
	if resp.Code != "OK" {
		return "", fmt.Errorf("Call by tts to user: %s failed, err_code: %s, err_msg: %s", phone, resp.Code, resp.Message)
	}

	return resp.CallId, nil
} // CallByTts()
