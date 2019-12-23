package incomingwebhook

type TextObject struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Emoji    bool   `json:"emoji,omitempty"`
	Verbatim bool   `json:"verbatim,omitempty"`
}

type ConfirmationDialogObject struct {
	Title   TextObject `json:"title"`
	Text    TextObject `json:"text"`
	Confirm TextObject `json:"confirm"`
	Deny    TextObject `json:"deny"`
}

type OptionObject struct {
	Text  TextObject `json:"text"`
	Value string     `json:"value"`
	URL   string     `json:"url,omitempty"`
}

type OptionGroupObject struct {
	Label   TextObject     `json:"label"`
	Options []OptionObject `json:"options"`
}

type Object interface {
	Object()
}

func (o TextObject) Object()               {}
func (o ConfirmationDialogObject) Object() {}
func (o OptionObject) Object()             {}
func (o OptionGroupObject) Object()        {}
