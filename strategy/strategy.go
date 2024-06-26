package strategy

type Strategy interface {
	Apply([]string) []string
}
