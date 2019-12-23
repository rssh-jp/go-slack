package incomingwebhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type SlackMessage struct {
	whURL string
}

func New(webhookURL string) *SlackMessage {
	return &SlackMessage{
		whURL: webhookURL,
	}
}

func (s SlackMessage) Send(mp MessagePayload) error {
	v, err := json.Marshal(mp)
	if err != nil {
		return err
	}

	contents := bytes.NewBuffer(v)

	err = s.request(contents)
	if err != nil {
		return err
	}

	return nil
}

func (s SlackMessage) SendJSON(j string) error {
	contents := bytes.NewBufferString(j)

	err := s.request(contents)
	if err != nil {
		return err
	}

	return nil
}

func (s SlackMessage) SendSimple(text string) error {
	m := MessagePayload{
		Text: text,
	}

	v, err := json.Marshal(m)
	if err != nil {
		return err
	}

	contents := bytes.NewBuffer(v)

	err = s.request(contents)
	if err != nil {
		return err
	}

	return nil
}

func (s SlackMessage) SendAttachment(obj Attachment) error {
	o := MessagePayload{
		Attachments: []Attachment{
			Attachment{
				Color:  obj.Color,
				Blocks: obj.Blocks,
			},
		},
	}

	jsonStr, err := json.Marshal(o)
	if err != nil {
		return err
	}

	contents := bytes.NewBuffer(jsonStr)

	err = s.request(contents)
	if err != nil {
		return err
	}

	return nil
}

func (s SlackMessage) request(contents io.Reader) error {
	req, err := http.NewRequest("POST", s.whURL, contents)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	b := make([]byte, 1024)
	n, err := res.Body.Read(b)
	if err != nil && err != io.EOF {
		return err
	}

	if !strings.Contains(string(b[:n]), "ok") {
		return errors.New(string(b[:n]))
	}

	return nil
}
