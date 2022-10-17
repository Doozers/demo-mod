package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	demomod "moul.io/adapterkit-module-demo"
)

func getClient() (demomod.DemomodSvcClient, error) {
	conn, err := grpc.Dial("127.0.0.1:9314", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return demomod.NewDemomodSvcClient(conn), nil
}
