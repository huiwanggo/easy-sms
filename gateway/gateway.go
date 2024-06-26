package gateway

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/huiwanggo/easy-sms/messenger"
)

const Success = "success"
const Failure = "failure"

type Config map[string]string

type Result struct {
	Status string
	Result interface{}
}

type Gateway interface {
	GetName() string
	SetConfig(config Config)
	GetConfig() Config
	GetClient() *resty.Client
	SetClient(client *resty.Client)
	Send(ctx context.Context, phone messenger.Phone, message messenger.Message) (Result, error)
}

var registeredGateways = make(map[string]Gateway)

func Register(gateway Gateway) {
	if gateway == nil {
		panic("gateway: Register gateway is nil")
	}
	if gateway.GetName() == "" {
		panic("gateway: Register gateway name is empty")
	}
	gateway.SetClient(resty.New())
	registeredGateways[gateway.GetName()] = gateway
}

func GetGateway(name string) Gateway {
	return registeredGateways[name]
}

type BaseGateway struct {
	Name   string
	Client *resty.Client
	Config Config
}

func (g *BaseGateway) GetName() string {
	return g.Name
}

func (g *BaseGateway) GetConfig() Config {
	return g.Config
}

func (g *BaseGateway) SetConfig(config Config) {
	g.Config = config
}

func (g *BaseGateway) GetClient() *resty.Client {
	return g.Client
}

func (g *BaseGateway) SetClient(client *resty.Client) {
	g.Client = client
}
