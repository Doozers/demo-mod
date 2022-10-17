package main

import (
	"context"

	demomod "moul.io/adapterkit-module-demo"
)

func sumAction(c demomod.DemomodSvcClient, A, B int64) (int64, error) {
	result, err := c.Sum(context.Background(), &demomod.Sum_Request{
		A: A,
		B: B,
	})
	if err != nil {
		return 0, err
	}

	return result.C, nil
}
