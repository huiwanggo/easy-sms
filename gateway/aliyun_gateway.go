package gateway

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/huiwanggo/easy-sms/messenger"
	"net/url"
	"strings"
	"time"
)

func init() {
	Register(&AliyunGateway{BaseGateway{Name: "aliyun"}})
}

type AliyunGateway struct{ BaseGateway }

type AliyunResult map[string]string

func (gw *AliyunGateway) Send(ctx context.Context, phone messenger.Phone, message messenger.Message) (Result, error) {
	config := gw.GetConfig()
	accessKeyId := config["accessKeyId"]
	accessKeySecret := config["accessKeySecret"]
	signName := config["signName"]

	var result Result
	result.Status = Failure

	templateParam := message.GetData()
	marshal, err := json.Marshal(templateParam)
	if err != nil {
		return result, err
	}

	params := map[string]string{
		"Format":           "JSON",
		"Version":          "2017-05-25",
		"AccessKeyId":      accessKeyId,
		"Timestamp":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"SignatureNonce":   uuid.NewString(),
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureVersion": "1.0",
		"RegionId":         "cn-hangzhou",
		"Action":           "SendSms",
		"PhoneNumbers":     phone.GetNumber(),
		"SignName":         signName,
		"TemplateCode":     message.GetTemplate(gw.Name),
		"TemplateParam":    string(marshal),
	}
	params["Signature"] = AliyunGenerateSign("GET", accessKeySecret, params)

	response, err := gw.GetClient().R().SetContext(ctx).SetQueryParams(params).Get("https://dysmsapi.aliyuncs.com")
	if err != nil {
		return result, err
	}

	body := make(AliyunResult)
	err = json.Unmarshal(response.Body(), &body)
	if err != nil {
		return result, err
	}
	result.Result = body

	if body["Code"] == "OK" {
		result.Status = Success
	}

	return result, errors.New(body["Message"])
}

func AliyunGenerateSign(method string, accessKeySecret string, params map[string]string) string {
	urlEncoder := url.Values{}
	for key, value := range params {
		urlEncoder.Add(key, value)
	}
	stringToSign := urlEncoder.Encode()

	stringToSign = strings.ReplaceAll(stringToSign, "+", "%20")
	stringToSign = strings.ReplaceAll(stringToSign, "*", "%2A")
	stringToSign = strings.ReplaceAll(stringToSign, "%7E", "~")

	stringToSign = url.QueryEscape(stringToSign)

	stringToSign = method + "&%2F&" + stringToSign

	key := []byte(accessKeySecret + "&")
	h := hmac.New(sha1.New, key)
	h.Write([]byte(stringToSign))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
