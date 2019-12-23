package incomingwebhook

type ActionBlock struct {
	Type     string    `json:"type"`
	Elements []Element `json:"elements"`
	BlockID  string    `json:"block_id,omitempty"`
}

type ContextBlock struct {
	Type     string    `json:"type"`
	Elements []Element `json:"elements"`
	BlockID  string    `json:"block_id,omitempty"`
}

type DividerBlock struct {
	Type    string `json:"type"`
	BlockID string `json:"block_id,omitempty"`
}

type ImageBlock struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
	Title    string `json:"title,omitempty"`
	BlockID  string `json:"block_id,omitempty"`
}

type FileBlock struct {
	Type       string `json:"type"`
	ExternalID string `json:"external_id"`
	Source     string `json:"source"`
	BlockID    string `json:"block_id,omitempty"`
}

type InputBlock struct {
	Type     string     `json:"type"`
	Label    TextObject `json:"label"`
	Element  Element    `json:"element"`
	BlockID  string     `json:"block_id,omitempty"`
	Hint     TextObject `json:"hint"`
	Optional bool       `json:"optional,omitempty"`
}

type SectionBlock struct {
	Type      string       `json:"type"`
	Text      TextObject   `json:"text"`
	BlockID   string       `json:"block_id,omitempty"`
	Fields    []TextObject `json:"fields,omitempty"`
	Accessory Element      `json:"accessory,omitempty"`
}

type Block interface {
	Block()
}

func (b ActionBlock) Block()  {}
func (b ContextBlock) Block() {}
func (b DividerBlock) Block() {}
func (b ImageBlock) Block()   {}
func (b FileBlock) Block()    {}
func (b InputBlock) Block()   {}
func (b SectionBlock) Block() {}
