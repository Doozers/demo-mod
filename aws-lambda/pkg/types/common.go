package types

import (
	demomod "moul.io/adapterkit-module-demo"
)

type Request struct {
	RequestType string               `json:"requestType"`
	Sum         *demomod.Sum_Request `json:"sumRequest"`
}

type Response struct {
	HelloResult *demomod.HelloResult `json:"helloResult,omitempty"`
	SumResult   *demomod.Sum_Result  `json:"sumResult,omitempty"`
}
