package gateway

import (
	"context"
	"errors"
	"github.com/huiwanggo/easy-sms/messenger"
	"strings"
)

func init() {
	Register(&MitakeGateway{BaseGateway{Name: "mitake"}})
}

type MitakeGateway struct{ BaseGateway }

type MitakeResult map[string]string

func (gw *MitakeGateway) Send(ctx context.Context, phone messenger.Phone, message messenger.Message) (Result, error) {
	config := gw.GetConfig()
	username := config["username"]
	password := config["password"]
	callback := config["callback"]

	var result Result
	result.Status = Failure

	params := map[string]string{
		"username":     username,
		"password":     password,
		"CharsetURL":   "UTF8",
		"smsPointFlag": "1",
		"dstaddr":      phone.GetNumber(),
		"smbody":       message.GetContent(),
	}
	if callback != "" {
		params["response"] = callback
	}

	response, err := gw.GetClient().R().SetContext(ctx).SetQueryParams(params).Get("http://smsapi.mitake.com.tw/api/mtk/SmSend")
	if err != nil {
		return result, err
	}

	body := make(MitakeResult)
	bs := strings.Split(strings.ReplaceAll(string(response.Body()), "\r\n", "\n"), "\n")
	for _, s := range bs {
		field := strings.Split(s, "=")
		if len(field) == 2 {
			body[field[0]] = field[1]
		}
	}
	result.Result = body

	if body["statuscode"] == "0" || body["statuscode"] == "1" || body["statuscode"] == "2" || body["statuscode"] == "4" {
		result.Status = Success
		return result, nil
	}

	return result, errors.New(body["Error"])
}
