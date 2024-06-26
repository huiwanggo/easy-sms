package messenger

type Phone struct {
	Number  string
	IDDCode string
}

func (p *Phone) GetNumber() string {
	return p.Number
}

func (p *Phone) GetUniversalNumber() string {
	return p.GetPrefixedIDDCode("00") + p.Number
}

func (p *Phone) GetZeroPrefixedNumber() string {
	return p.GetPrefixedIDDCode("+") + p.Number
}

func (p *Phone) GetPrefixedIDDCode(prefix string) string {
	if p.IDDCode != "" {
		return prefix + p.IDDCode
	}
	return p.IDDCode
}
