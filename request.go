package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
)

// Request contains data from API Gateway invocation parameters
type Request struct {
	Token   string
	Channel string
	Payload Payload
}

func parseRequest(req events.APIGatewayProxyRequest) (*Request, error) {
	payload, err := parsePayload([]byte(req.Body))

	if err != nil {
		return nil, errors.New("Unable to parse request payload")
	}

	return &Request{
		req.PathParameters["token"],
		req.PathParameters["channel"],
		*payload,
	}, nil
}
