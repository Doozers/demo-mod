package adapterKit

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	demomod "moul.io/adapterkit-module-demo"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	service := demomod.New("test")
	demomod.RegisterDemomodSvcServer(server, service)
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func DemomodSvcSum(A, B int64) (*demomod.Sum_Result, error) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := demomod.NewDemomodSvcClient(conn)
	req := &demomod.Sum_Request{A: A, B: B}
	ret, err := client.Sum(ctx, req)
	return ret, err
}

func DemomodSvcSayHello() (*demomod.HelloResult, error) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := demomod.NewDemomodSvcClient(conn)
	req := &demomod.Empty{}
	ret, err := client.SayHello(ctx, req)
	return ret, err
}
