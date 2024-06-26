package messenger

type Message struct {
	Content  string
	Template map[string]string
	Data     map[string]string
}

func (m Message) GetContent() string {
	return m.Content
}

func (m Message) GetTemplate(gateway string) string {
	template := m.Template[gateway]
	if template != "" {
		return template
	}
	return m.Template["default"]
}

func (m Message) GetData() map[string]string {
	return m.Data
}
