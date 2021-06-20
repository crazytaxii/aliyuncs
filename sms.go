package aliyuncs

type smsContext struct {
	phone        string
	signName     string
	templateCode string
	CSTemplate
}

func (c smsContext) getPhone() string {
	return c.phone
}

func (c smsContext) genReqParams() (map[string]string, error) {
	str, err := c.CSTemplate.Marshal()
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"SignName":      c.signName,
		"PhoneNumbers":  c.phone,
		"TemplateCode":  c.templateCode,
		"TemplateParam": str,
	}, nil
}

type SendSMSResponse struct {
	commonResponse
	BizID string `json:"BizId"`
}

/**
 * 向指定手机号发送短信
 */
func (c *Client) SendSMS(phone, signName, templateCode string, template CSTemplate) (string, error) {
	ctx := smsContext{
		phone:        phone,
		signName:     signName,
		templateCode: templateCode,
		CSTemplate:   template,
	}

	p, err := c.genReqParams(ctx, ActionSMS)
	if err != nil {
		return "", err
	}

	resp := &SendSMSResponse{}
	if err := c.doGet(ctx, p, resp); err != nil {
		return "", err
	}
	return resp.BizID, nil
}
