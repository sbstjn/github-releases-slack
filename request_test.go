package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	req := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"token":   "example-token",
			"channel": "example-channel",
		},
		Body: "{}",
	}

	request, err := parseRequest(req)

	assert.Nil(t, err)
	assert.NotNil(t, request)

	assert.Equal(t, "example-token", request.Token)
	assert.Equal(t, "example-channel", request.Channel)
}
