package main

import (
	"context"
	"fmt"

	"github.com/Doozers/aws-lambda-test/pkg/action"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Request struct {
	RequestType string      `json:"requestType"`
	Sum         *action.Sum `json:"sum"`
}

func HandleRequest(req Request) (string, error) {
	switch req.RequestType {
	case "sum":
		return action.SumAction(req.Sum)

	case "sayHello":
		return action.SayHelloAction()

	default:
		return "", fmt.Errorf("invalid request type, %s", req.RequestType)
	}
}

func main() {
	lambda.Start(HandleRequest)
}
