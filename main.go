package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Person struct {
	Name string `json:"name"`
}

var people = []Person{
	{"Alex"},
	{"Bob"},
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	msg, err := json.Marshal(people)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       "Oops!!something went wrong",
		}, nil
	}
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(msg),
	}, nil
}
