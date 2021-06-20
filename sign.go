package aliyuncs

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
)

/**
 * 签名
 */
func (c *Client) doSign(method string, param map[string]string) string {
	// 根据参数key顺序排序
	pList := make([]string, 0)
	for key, value := range param {
		pList = append(pList, fmt.Sprintf("%s=%s", key, specialUrlEncode(value)))
	}
	sort.Strings(pList)

	// 构造待签名的请求串
	sortedQueryString := strings.Join(pList, "&")
	str2Sign := fmt.Sprintf("%s&%s&%s", method, specialUrlEncode("/"),
		specialUrlEncode(sortedQueryString))

	// 签名采用HmacSHA1算法 + Base64，编码采用：UTF-8
	h := hmac.New(sha1.New, []byte(fmt.Sprintf("%s&", c.accessSecret)))
	h.Write([]byte(str2Sign))
	hashed := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashed)
} // doSign()
