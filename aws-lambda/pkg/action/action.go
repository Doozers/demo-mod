package action

import (
	"fmt"

	"github.com/Doozers/demo-mod/aws-lambda/pkg/types"
)

func SumAction(s *types.Sum) (types.Response, error) {
	if s == nil {
		return types.Response{}, fmt.Errorf("`sum` field must be specified")
	}
	return types.Response{Result: s.A + s.B}, nil
}

func SayHelloAction() (types.Response, error) {
	return types.Response{Message: "Hello World !"}, nil
}
