package easysms

import (
	"context"
	"fmt"
	"github.com/huiwanggo/easy-sms/gateway"
	"github.com/huiwanggo/easy-sms/messenger"
	"github.com/huiwanggo/easy-sms/strategy"
	"time"
)

type Sms struct {
	gateways map[string]gateway.Config // 可用的网关配置
	allows   []string                  // 允许可用的网关
	timeout  time.Duration             // 请求超时时间
	strategy strategy.Strategy         // 网关调用策略
}

func New(gateways map[string]gateway.Config, allows []string, opts ...Option) *Sms {
	m := &Sms{
		allows:   allows,
		gateways: gateways,
		timeout:  time.Second * 5,
		strategy: strategy.OrderStrategy{},
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

type Option func(m *Sms)

func WithTimeout(timeout time.Duration) Option {
	return func(m *Sms) {
		m.timeout = timeout
	}
}

func WithStrategy(strategy strategy.Strategy) Option {
	return func(m *Sms) {
		m.strategy = strategy
	}
}

func (s *Sms) Send(ctx context.Context, phone messenger.Phone, message messenger.Message) (map[string]gateway.Result, error) {
	var err error
	results := make(map[string]gateway.Result)

	allows := s.strategy.Apply(s.allows)
	for _, allow := range allows {
		gw := gateway.GetGateway(allow)
		if gw != nil {
			gw.SetConfig(s.gateways[allow])
			gw.GetClient().SetTimeout(s.timeout)
			result, err := gw.Send(ctx, phone, message)
			results[gw.GetName()] = result
			if err == nil {
				break
			}
		} else {
			err = fmt.Errorf("gateway %s is not found", allow)
		}
	}

	return results, err
}
