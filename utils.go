package aliyuncs

import (
	"crypto/rand"
	"encoding/hex"
	"net/url"
	"strings"
)

/**
 * 生成指定长度的随机字符串
 */
func randStr(l int) string {
	half := l / 2
	odd := (l%2 != 0)
	if odd {
		half = half + 1
	}
	bytes := make([]byte, half)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	str := hex.EncodeToString(bytes)
	if odd {
		return str[0 : len(str)-1]
	}
	return str
} // RandStr()

/**
 * 特殊URL编码
 */
func specialUrlEncode(content string) string {
	value := make(url.Values)
	value.Add("", content)
	result := strings.Replace(value.Encode(), "=", "", 1)
	result =
		strings.Replace(
			strings.Replace(
				strings.Replace(
					result, "+", "%20", -1,
				), "*", "%2A", -1,
			), "%7E", "~", -1,
		)
	return result
} // specialUrlEncode()
