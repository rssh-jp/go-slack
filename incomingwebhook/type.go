package incomingwebhook

type MessagePayload struct {
	Text        string       `json:"text"`
	Blocks      []Block      `json:"blocks,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	ThreadTS    string       `json:"thread_ts,omitempty"`
	Mrkdwn      bool         `json:"mrkdwn,omitempty"`
}

type Attachment struct {
	Blocks []Block `json:"blocks,omitempty"`
	Color  string  `json:"color,omitempty"`

	Fallback      string  `json:"fallback,omitempty"`
	Pretext       string  `json:"pretext,omitempty"`
	AuthorName    string  `json:"author_name,omitempty"`
	AuthorLink    string  `json:"author_link,omitempty"`
	Title         string  `json:"title,omitempty"`
	TitleLink     string  `json:"title_link,omitempty"`
	Text          string  `json:"text,omitempty"`
	ImageURL      string  `json:"image_url,omitempty"`
	ThumbURL      string  `json:"thumb_url,omitempty"`
	Footer        string  `json:"footer,omitempty"`
	FooterIcon    string  `json:"footer_icon,omitempty"`
	UnixTimeStamp int     `json:"ts,omitempty"`
	Fields        []Field `json:"fields,omitempty"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}
