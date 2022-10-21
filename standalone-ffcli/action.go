package main

import (
	"github.com/Doozers/demo-mod/standalone-ffcli/pkg/adapterKit"
)

func SumAction(A, B int64) (int64, error) {
	res, err := adapterKit.DemomodSvcSum(A, B)
	if err != nil {
		return 0, err
	}
	return res.C, nil
}

func SayHelloAction() (string, error) {
	res, err := adapterKit.DemomodSvcSayHello()
	if err != nil {
		return "", err
	}
	return res.Msg, nil
}

func EchoStreamAction(sent string) {
	panic("not implemented")
}
