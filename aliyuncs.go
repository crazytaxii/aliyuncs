package aliyuncs

import ()

const DOMAIN = "dysmsapi.aliyuncs.com"

type AliyunCsConfig struct {
	AccessKeyId  string
	AccessSecret string
	SignName     string
}

var config = new(AliyunCsConfig)

/**
 * 阿里云通讯初始化
 */
func Init(accessKeyId string, accessSecret string, signName string) {
	config.AccessKeyId = accessKeyId
	config.AccessSecret = accessSecret
	config.SignName = signName
} // Init()
