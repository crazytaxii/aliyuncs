# aliyuncs
阿里云通讯Golang SDK

## 准备
### 创建阿里云账号
为了访问短信服务，需要有一个阿里云账号。

### 获取阿里云访问密钥
阿里云访问密钥是阿里云为用户使用 API（非控制台）来访问其云资源设计的“安全口令”。
您可以用它来签名 API 请求内容以通过服务端的安全验证。

该访问秘钥成对（AccessKeyId 与 AccessKeySecret）生成和使用。

### 短信签名
根据用户属性来创建符合自身属性的签名信息。[签名申请手册](https://help.aliyun.com/document_detail/55327.html?spm=a2c4g.11186623.2.6.f6b36C)

**注意：短信签名需要审核通过后才可以使用。**

### 短信模版
短信模板，即具体发送的短信内容。[模版申请手册](https://help.aliyun.com/document_detail/55330.html?spm=a2c4g.11186623.2.7.f6b36C)

## 安装
```Bash
$ go get github.com/crazytaxii/aliyuncs
```

## 使用实例
### 初始化
```Go
import (
    cs "github.com/crazytaxii/aliyuncs"
)

func init() {
    cs.Init(
        "access_key_id",
        "access_key_secret",
    )
}
```

### 向指定手机号（用户）发送短信
```Go
func sendSMS() {
    param := map[string]string{
        "code": "123456",
    }

    _, err := cs.SendSMS(
        "13800000000", // 指定手机号
        "sign_name", // 短信签名
        param, // 短信模板变量
        "SMS_123456789", // 短信模板ID
    )
    if err != nil {
        fmt.Println("err:", err)
    }
}
```