package ds

type MessageBuilder struct {
	channelID Snowflake
	client    *Client
	MessageBody
}

func NewMessageBuilder(channelID Snowflake, client *Client) *MessageBuilder {
	return &MessageBuilder{channelID: channelID, client: client}
}

func (m *MessageBuilder) WithContent(c string) *MessageBuilder {
	m.Content = c
	return m
}

func (m *MessageBuilder) SetContent(c string) {
	m.Content = c
}

func (m *MessageBuilder) WithEmbeds(e []*Embed) *MessageBuilder {
	m.Embeds = e
	return m
}

func (m *MessageBuilder) SetEmbeds(e []*Embed) {
	m.Embeds = e
}

func (m *MessageBuilder) WithEmbed(e *Embed) *MessageBuilder {
	m.Embeds = append(m.Embeds, e)
	return m
}

func (m *MessageBuilder) AddEmbed(e *Embed) {
	m.Embeds = append(m.Embeds, e)
}

func (m *MessageBuilder) WithTTS(t bool) *MessageBuilder {
	m.TTS = &t
	return m
}

func (m *MessageBuilder) SetTTS(t bool) {
	m.TTS = &t
}

func (m *MessageBuilder) WithAttachments(a []*Attachment) *MessageBuilder {
	m.Attachments = a
	return m
}

func (m *MessageBuilder) SetAttachments(a []*Attachment) {
	m.Attachments = a
}

func (m *MessageBuilder) WithAttachment(a *Attachment) *MessageBuilder {
	m.Attachments = append(m.Attachments, a)
	return m
}

func (m *MessageBuilder) AddAttachment(a *Attachment) {
	m.Attachments = append(m.Attachments, a)
}

func (m *MessageBuilder) WithComponents(c []MessageComponent) *MessageBuilder {
	m.Components = c
	return m
}

func (m *MessageBuilder) WithComponent(c MessageComponent) *MessageBuilder {
	m.Components = append(m.Components, c)
	return m
}

func (m *MessageBuilder) SetComponents(c []MessageComponent) {
	m.Components = c
}

func (m *MessageBuilder) AddComponent(c MessageComponent) {
	m.Components = append(m.Components, c)
}

func (m *MessageBuilder) WithFiles(ms []*MessageFile) *MessageBuilder {
	m.Files = ms
	return m
}

func (m *MessageBuilder) SetFiles(ms []*MessageFile) {
	m.Files = ms
}

func (m *MessageBuilder) WithFile(ms *MessageFile) *MessageBuilder {
	m.Files = append(m.Files, ms)
	return m
}

func (m *MessageBuilder) AddFile(ms *MessageFile) {
	m.Files = append(m.Files, ms)
}

func (m *MessageBuilder) Send() (*Message, error) {
	return m.client.Api.SendMessage(m.channelID, &m.MessageBody)

}
