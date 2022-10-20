package main

import (
	"fmt"

	"github.com/Doozers/demo-mod/aws-lambda/pkg/action"
	"github.com/Doozers/demo-mod/aws-lambda/pkg/types"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(req types.Request) (types.Response, error) {
	switch req.RequestType {
	case "sum":
		return action.SumAction(req.Sum)

	case "sayHello":
		return action.SayHelloAction()

	default:
		return types.Response{}, fmt.Errorf("invalid request type, %s", req.RequestType)
	}
}

func main() {
	lambda.Start(HandleRequest)
}
