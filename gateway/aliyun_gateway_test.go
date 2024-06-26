package gateway

import (
	"testing"
)

func TestAliyunGenerateSign(t *testing.T) {
	params := map[string]string{
		"Format":           "JSON",
		"Version":          "2017-05-25",
		"AccessKeyId":      "id",
		"Timestamp":        "2024-06-22T05:40:29Z",
		"SignatureNonce":   "24da9959-2b0a-428f-95db-484f55a666b1",
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureVersion": "1.0",
		"RegionId":         "cn-hangzhou",
		"Action":           "SendSms",
		"PhoneNumbers":     "18888888888",
		"SignName":         "阿里云短信测试",
		"TemplateCode":     "SMS_154950909",
		"TemplateParam":    "{\"code\":\"1234\"}",
	}
	sign := AliyunGenerateSign("GET", "secret", params)
	expect := "lPBRSvsWtyHndM8/8OiXd2QkEWI="
	if sign != expect {
		t.Error("generate generateSign error")
	}
}
