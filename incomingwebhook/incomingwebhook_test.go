package incomingwebhook

import (
	"testing"
)

const (
	webhookURL = "https://hooks.slack.com/services/xxxxxxxxx/xxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxx"
)

var (
	s *SlackMessage
)

func preprocess() error {
	s = New(webhookURL)
	return nil
}

func postprocess() error {
	return nil
}

func TestMain(m *testing.M) {
	preprocess()
	defer postprocess()

	m.Run()
}

func TestSuccessSimple(t *testing.T) {
	err := s.SendSimple("simple text!")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSuccess(t *testing.T) {
	m := MessagePayload{
		Attachments: []Attachment{
			Attachment{
				Color: "#33dd33",
				Blocks: []Block{
					SectionBlock{
						Type: "section",
						Text: TextObject{
							Type: "mrkdwn",
							Text: "A message *with some bold text* and _some italicized text_.",
						},
					},
					SectionBlock{
						Type: "section",
						Text: TextObject{
							Type: "plain_text",
							Text: "block_text",
						},
						Fields: []TextObject{
							TextObject{
								Type: "plain_text",
								Text: "block_field_text",
							},
						},
						Accessory: ButtonElement{
							Type: "button",
							Text: TextObject{
								Type: "plain_text",
								Text: "block-section-accessory-text",
							},
							ActionID: "100",
							Style:    "primary",
						},
					},
				},
			},
			Attachment{
				Color: "good",
				Text:  "nice",
			},
		},
	}

	err := s.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSuccessAttachment(t *testing.T) {
	obj := Attachment{
		Color: "#33dd33",
		Blocks: []Block{
			SectionBlock{
				Type: "section",
				Text: TextObject{
					Type: "mrkdwn",
					Text: "A message *with some bold text* and _some italicized text_.",
				},
			},
			SectionBlock{
				Type: "section",
				Text: TextObject{
					Type: "plain_text",
					Text: "block_text",
				},
				Fields: []TextObject{
					TextObject{
						Type: "plain_text",
						Text: "block_field_text",
					},
				},
				Accessory: ButtonElement{
					Type: "button",
					Text: TextObject{
						Type: "plain_text",
						Text: "block-section-accessory-text",
					},
					ActionID: "100",
					Style:    "primary",
				},
			},
		},
	}

	err := s.SendAttachment(obj)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChannelLink(t *testing.T) {
	m := MessagePayload{
		Blocks: []Block{
			SectionBlock{
				Type: "section",
				Text: TextObject{
					Type: "plain_text",
					Text: "Invitation!!",
				},
			},
		},
		Attachments: []Attachment{
			Attachment{
				Color: "#dddd44",
				Blocks: []Block{
					SectionBlock{
						Type: "section",
						Text: TextObject{
							Type: "mrkdwn",
							Text: "Join us. <#XXXXXXXX>?",
						},
					},
				},
			},
		},
	}

	err := s.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestButton(t *testing.T) {
	m := MessagePayload{
		Blocks: []Block{
			ActionBlock{
				Type:    "actions",
				BlockID: "blockid",
				Elements: []Element{
					ButtonElement{
						Type: "button",
						Text: TextObject{
							Type: "plain_text",
							Text: "Primary",
						},
						ActionID: "abc",
						Style:    "primary",
						URL:      "https://api.slack.com/reference/block-kit/block-elements#button",
					},
					ButtonElement{
						Type: "button",
						Text: TextObject{
							Type: "plain_text",
							Text: "Cancel",
						},
						ActionID: "def",
						URL:      "https://api.slack.com/reference/block-kit/block-elements",
					},
				},
			},
		},
	}

	err := s.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJSON(t *testing.T) {
	json := `
{
    "blocks": [
        {
            "type": "section",
            "text": {
                "type": "mrkdwn",
                "text": "*Where should we order lunch from?* Poll by <fakeLink.toUser.com|Mark>"
            }
        },
        {
            "type": "divider"
        },
        {
            "type": "section",
            "text": {
                "type": "mrkdwn",
                "text": ":sushi: *Ace Wasabi Rock-n-Roll Sushi Bar*\nThe best landlocked sushi restaurant."
            },
            "accessory": {
                "type": "button",
                "text": {
                    "type": "plain_text",
                    "emoji": true,
                    "text": "Vote"
                },
                "value": "click_me_123"
            }
        },
        {
            "type": "context",
            "elements": [
                {
                    "type": "image",
                    "image_url": "https://api.slack.com/img/blocks/bkb_template_images/profile_1.png",
                    "alt_text": "Michael Scott"
                },
                {
                    "type": "image",
                    "image_url": "https://api.slack.com/img/blocks/bkb_template_images/profile_2.png",
                    "alt_text": "Dwight Schrute"
                },
                {
                    "type": "image",
                    "image_url": "https://api.slack.com/img/blocks/bkb_template_images/profile_3.png",
                    "alt_text": "Pam Beasely"
                },
                {
                    "type": "plain_text",
                    "emoji": true,
                    "text": "3 votes"
                }
            ]
        },
        {
            "type": "section",
            "text": {
                "type": "mrkdwn",
                "text": ":hamburger: *Super Hungryman Hamburgers*\nOnly for the hungriest of the hungry."
            },
            "accessory": {
                "type": "button",
                "text": {
                    "type": "plain_text",
                    "emoji": true,
                    "text": "Vote"
                },
                "value": "click_me_123"
            }
        },
        {
            "type": "context",
            "elements": [
                {
                    "type": "image",
                    "image_url": "https://api.slack.com/img/blocks/bkb_template_images/profile_4.png",
                    "alt_text": "Angela"
                },
                {
                    "type": "image",
                    "image_url": "https://api.slack.com/img/blocks/bkb_template_images/profile_2.png",
                    "alt_text": "Dwight Schrute"
                },
                {
                    "type": "plain_text",
                    "emoji": true,
                    "text": "2 votes"
                }
            ]
        },
        {
            "type": "section",
            "text": {
                "type": "mrkdwn",
                "text": ":ramen: *Kagawa-Ya Udon Noodle Shop*\nDo you like to shop for noodles? We have noodles."
            },
            "accessory": {
                "type": "button",
                "text": {
                    "type": "plain_text",
                    "emoji": true,
                    "text": "Vote"
                },
                "value": "click_me_123"
            }
        },
        {
            "type": "context",
            "elements": [
                {
                    "type": "mrkdwn",
                    "text": "No votes"
                }
            ]
        },
        {
            "type": "divider"
        },
        {
            "type": "actions",
            "elements": [
                {
                    "type": "button",
                    "text": {
                        "type": "plain_text",
                        "emoji": true,
                        "text": "Add a suggestion"
                    },
                    "value": "click_me_123"
                }
            ]
        }
    ]
}
    `

	err := s.SendJSON(json)
	if err != nil {
		t.Fatal(err)
	}
}
