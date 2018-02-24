package main

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPayloadRelease(t *testing.T) {
	data, err := ioutil.ReadFile("./fixtures/payload.json")

	assert.Nil(t, err, "Cannot open file")
	assert.True(t, strings.Contains(string(data), "\"action\": \"published\","))

	payload, err := parsePayload(data)

	assert.Nil(t, err, "Cannot parse content")
	assert.Equal(t, "v0.0.1", payload.Release.Name)
	assert.Equal(t, "### Add\r\n\r\n* Initial commit", payload.Release.Body)
	assert.Equal(t, false, payload.Release.Draft)
	assert.Equal(t, false, payload.Release.Prerelease)

	date, err := time.Parse(time.RFC3339, "2018-02-23T21:45:57Z")

	assert.Nil(t, err)
	assert.Equal(t, date, payload.Release.Date)
}

func TestPayloadSender(t *testing.T) {
	data, _ := ioutil.ReadFile("./fixtures/payload.json")
	payload, _ := parsePayload(data)

	assert.Equal(t, "sbstjn", payload.Sender.Name)
}

func TestPayloadRepository(t *testing.T) {
	data, _ := ioutil.ReadFile("./fixtures/payload.json")
	payload, _ := parsePayload(data)

	assert.Equal(t, "github-releases-slack", payload.Repository.Name)
	assert.Equal(t, "sbstjn/github-releases-slack", payload.Repository.FullName)
	assert.Equal(t, "https://github.com/sbstjn/github-releases-slack", payload.Repository.URL)
}
