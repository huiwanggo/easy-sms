# easy-sms
Go 简单便捷的短信发送组件

## 安装

```
go get github.com/huiwanggo/easy-sms
```

## 使用

```go
package main

import (
	"context"
	"fmt"
	"github.com/huiwanggo/easy-sms"
	"github.com/huiwanggo/easy-sms/gateway"
	"github.com/huiwanggo/easy-sms/messenger"
	"github.com/huiwanggo/easy-sms/strategy"
	"time"
)

func main() {
	sms := easysms.New(
		map[string]gateway.Config{
			"aliyun": {
				"accessKeyId":     "accessKeyIdTest",
				"accessKeySecret": "accessKeySecretTest",
				"signName":        "阿里云短信测试",
			},
			"mitake": {
				"username": "username",
				"password": "password",
				"callback": "url",
			},
		},
		[]string{"aliyun", "mitake"},
		easysms.WithTimeout(5*time.Second),
		easysms.WithStrategy(strategy.RandomStrategy{}),
	)
	phone := messenger.Phone{Number: "18888888888"}
	message := messenger.Message{
		Content:  "您的验证码为: 1234",
		Template: map[string]string{"aliyun": "SMS_154950909"},
		Data:     map[string]string{"code": "1234"},
	}
	results, err := sms.Send(context.Background(), phone, message)
	fmt.Println(results, err)
}
```

### easysms.New

- 第一个参数: 配置 map[string]gateway.Config 配置各个短信平台配置
- 第二个参数: 允许使用的短信平台，发送策略为 OrderStrategy 时候 按照配置顺序发送
- 第三个参数: 可选参数 easysms.WithTimeout(timeout) 发送超时时间
- 第四个参数: 可选参数 easysms.WithStrategy(strategy) 发送策略：OrderStrategy 顺序发送，RandomStrategy 随机发送

### messenger.Phone

抽象了手机号

- Number 为 string 格式，手机号
- IDDCode 为 string 格式，国际码，国际短信需要指定

### messenger.Message

抽象了短信内容

- Template 为 map[string]string 格式，key 为短信平台名称，value 为短信模板编号，支持多个平台，key 可为 default，未指定平台时使用
- Data 为 map[string]string 格式，key 为短信模板占位符，value 为占位符值，支持多个平台，key 可为 default，未指定平台时使用
- Content 为 string 格式，短信内容

## 支持平台

### 阿里云短信

- [阿里云短信服务](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-sendsms?spm=a2c4g.11186623.0.0.4f6a5f9fiBc59t)
- 平台配置

```
map[string]gateway.Config{
    "aliyun": {
        "accessKeyId":     "accessKeyIdTest",
        "accessKeySecret": "accessKeySecretTest",
        "signName":        "阿里云短信测试",
    },
}
```

- 短信内容（模板编号+模板变量）

```
messenger.Message{
    Template: map[string]string{"aliyun": "SMS_154950909"},
    Data:     map[string]string{"code": "1234"},
}
```

### 台湾三竹简讯

- [三竹简讯](https://sms.mitake.com.tw/common/index.jsp?t=1722999600763)
- 平台配置

```
map[string]gateway.Config{
    "mitake": {
        "username": "username",
        "password": "password",
        "callback": "发送状态回调地址（可选）",
    },
}
```

- 短信内容（文本）

```
messenger.Message{
    Content: "您的验证码为: 1234",
}
```
