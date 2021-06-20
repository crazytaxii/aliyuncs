package aliyuncs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (c *Client) doGet(ctx CSContext, reqParams map[string]string, resMsg ResponseMessage) error {
	req, err := http.NewRequest(http.MethodGet, BaseURL, nil)
	if err != nil {
		return err
	}

	// join query string
	q := req.URL.Query()
	for k, v := range reqParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.Get(req.URL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, resMsg); err != nil {
		return err
	}
	if err := resMsg.error(ctx); err != nil {
		return err
	}
	return nil
}
