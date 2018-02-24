package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMessageFromTemplate(t *testing.T) {
	template := MessageTemplateData{
		URL:      "https://example.com/example/project",
		FullName: "example/project",
		Version:  "v0.1.0",
	}

	message, err := messageFromTemplate(template)

	assert.Nil(t, err)
	assert.NotNil(t, message)
}

func TestMessageTemplateFromPayload(t *testing.T) {
	date, _ := time.Parse(time.RFC3339, "2018-02-23T21:45:57Z")
	payload := Payload{
		Release: PayloadRelease{
			Name:       "v0.1.0",
			Body:       "### Add\r\n\r\n* Initial commit",
			Date:       date,
			Draft:      false,
			Prerelease: false,
		},
		Repository: PayloadRepository{
			Name:     "project",
			FullName: "example/project",
			URL:      "https://github.com/example/project",
		},
		Sender: PayloadSender{
			Name: "sbstjn",
		},
	}

	m := messageTemplateFromPayloadForChannel(payload, "#example")

	assert.Equal(t, "example/project", m.FullName)
	assert.Equal(t, "https://github.com/example/project", m.URL)
	assert.Equal(t, "#example", m.Channel)
	assert.Equal(t, "v0.1.0", m.Version)
}
