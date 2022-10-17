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

func sayHelloAction(c demomod.DemomodSvcClient) (string, error) {
	result, err := c.SayHello(context.Background(), &demomod.Empty{})
	if err != nil {
		return "", err
	}

	return result.Msg, nil
}

func echoStreamAction(c demomod.DemomodSvcClient, send, receive *chan string) error {
	stream, err := c.EchoStream(context.Background())
	if err != nil {
		return err
	}

	go func() {
		for {
			text, ok := <-*send
			if !ok {
				return
			}
			err := stream.Send(&demomod.EchoStream_Input{Text: text})
			if err != nil {
				return
			}
		}
	}()

	go func() {
		for {
			result, err := stream.Recv()
			if err != nil {
				return
			}
			*receive <- result.GetReply()
		}
	}()

	return nil
}
