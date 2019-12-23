package incomingwebhook

type ButtonElement struct {
	Type     string                    `json:"type"`
	Text     TextObject                `json:"text"`
	ActionID string                    `json:"action_id"`
	URL      string                    `json:"url,omitempty"`
	Value    string                    `json:"value,omitempty"`
	Style    string                    `json:"style,omitempty"`
	Confirm  *ConfirmationDialogObject `json:"confirm,omitempty"`
}

type DatePickerElement struct {
	Type        string                    `json:"type"`
	ActionID    string                    `json:"action_id"`
	Placeholder TextObject                `json:"placeholder,omitempty"`
	InitialDate string                    `json:"initial_date,omitempty"`
	Confirm     *ConfirmationDialogObject `json:"confirm,omitempty"`
}

type ImageElement struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

type OverflowMenuElement struct {
	Type     string                    `json:"type"`
	ActionID string                    `json:"action_id"`
	Options  []OptionObject            `json:"options"`
	Confirm  *ConfirmationDialogObject `json:"confirm,omitempty"`
}

type PlainTextInputElement struct {
	Type         string     `json:"type"`
	ActionID     string     `json:"action_id"`
	Placeholder  TextObject `json:"placeholder,omitempty"`
	Initialvalue string     `json:"initial_value,omitempty"`
	Multiline    bool       `json:"multiline,omitempty"`
	MinLength    int        `json:"min_length,omitempty"`
	MaxLength    int        `json:"max_length,omitempty"`
}

type RadioButtonGroupElement struct {
	Type          string                    `json:"type"`
	ActionID      string                    `json:"action_id"`
	Options       []OptionObject            `json:"options"`
	InitialOption *OptionObject             `json:"initial_option,omitempty"`
	Confirm       *ConfirmationDialogObject `json:"confirm,omitempty"`
}

type Element interface {
	Element()
}

func (e ButtonElement) Element()           {}
func (e DatePickerElement) Element()       {}
func (e ImageElement) Element()            {}
func (e OverflowMenuElement) Element()     {}
func (e PlainTextInputElement) Element()   {}
func (e RadioButtonGroupElement) Element() {}
