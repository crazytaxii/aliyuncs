package aliyuncs

type callContext struct {
	phone   string
	showAs  string
	ttsCode string
	CSTemplate
}

func (c callContext) getPhone() string {
	return c.phone
}

func (c callContext) genReqParams() (map[string]string, error) {
	str, err := c.CSTemplate.Marshal()
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"CalledNumber":     c.phone,
		"CalledShowNumber": c.showAs,
		"TtsCode":          c.ttsCode,
		"TtsParam":         str,
	}, nil
}

type SingleCallByTTSResponse struct {
	commonResponse
	CallID string `json:"CallId"`
}

func (c *Client) CallByTTS(phone, showAs, ttsCode string, template CSTemplate) (string, error) {
	ctx := callContext{
		phone:      phone,
		showAs:     showAs,
		ttsCode:    ttsCode,
		CSTemplate: template,
	}

	p, err := c.genReqParams(ctx, ActionCall)
	if err != nil {
		return "", err
	}

	resp := &SingleCallByTTSResponse{}
	if err := c.doGet(ctx, p, resp); err != nil {
		return "", err
	}
	return resp.CallID, nil
}
