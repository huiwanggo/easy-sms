package strategy

type OrderStrategy struct{}

func (s OrderStrategy) Apply(gateways []string) []string {
	return gateways
}
