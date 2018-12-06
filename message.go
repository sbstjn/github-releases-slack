package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
)

// MessageTemplateData contains data sent to Slack
type MessageTemplateData struct {
	URL      string
	FullName string
	Version  string
	Notes    string
	Channel  string
}

// MessageResponse contains data received from Slack
type MessageResponse struct {
	OK bool `json:"ok"`
}

func messageFromRequest(request Request) ([]byte, error) {
	template := messageTemplateFromPayloadForChannel(request.Payload, request.Channel)

	return messageFromTemplate(template)
}

func messageTemplateFromPayloadForChannel(payload Payload, channel string) MessageTemplateData {
	return MessageTemplateData{
		payload.Repository.URL,
		payload.Repository.FullName,
		payload.Release.Name,
		payload.Release.Body,
		channel,
	}
}

func messageFromTemplate(data MessageTemplateData) ([]byte, error) {
	t, _ := template.ParseFiles("templates/message.json")

	var message bytes.Buffer
	if err := t.Execute(&message, data); err != nil {
		return nil, err
	}

	return message.Bytes(), nil
}

func postMessageToSlack(message []byte, token string) (*MessageResponse, error) {
	url := "https://slack.com/api/chat.postMessage"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var response MessageResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
