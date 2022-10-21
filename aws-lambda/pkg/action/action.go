package action

import (
	"fmt"

	demomod "moul.io/adapterkit-module-demo"

	"github.com/Doozers/demo-mod/aws-lambda/pkg/adapterKit"
	"github.com/Doozers/demo-mod/aws-lambda/pkg/types"
)

func SumAction(s *demomod.Sum_Request) (types.Response, error) {
	if s == nil {
		return types.Response{}, fmt.Errorf("`sum` field must be specified")
	}
	res, err := adapterKit.DemomodSvcSum(s.A, s.B)
	if err != nil {
		return types.Response{}, err
	}
	return types.Response{SumResult: res}, nil
}

func SayHelloAction() (types.Response, error) {
	res, err := adapterKit.DemomodSvcSayHello()
	if err != nil {
		return types.Response{}, err
	}
	return types.Response{HelloResult: res}, nil
}
