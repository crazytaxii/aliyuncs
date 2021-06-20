# aliyuncs

阿里云通讯（原阿里大于）Golang SDK

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

## 使用

```go
import (
    cs "github.com/crazytaxii/aliyuncs"
)


func main() {
    tmpl := cs.CSTemplate{
        "code": "123456",
    }

    client := cs.NewClient("access_id", "access_secret", 10 * time.Second)
    if _, err := client.SendSMS(
        "13712341234", // 指定手机号
        "sign_name", // 短信签名
        "SMS_123456789", // 短信模板ID
        tmpl, // 短信模板变量
    ); err != nil {
        log.Fatal("err:", err)
    }
}
```
