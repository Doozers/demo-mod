package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	demomod "moul.io/adapterkit-module-demo"

	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	err := basic(os.Args[1:])
	if err != nil {
		log.Fatalf("err: %+v", err)
	}

	return
}

func basic(args []string) error {
	rootFlagSet := flag.NewFlagSet("basic", flag.ExitOnError)

	root := ffcli.Command{
		ShortUsage: "basic [flags] <command>",
		FlagSet:    rootFlagSet,
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Subcommands: []*ffcli.Command{
			server(),
		},
		Exec: func(_ context.Context, _ []string) error {
			return flag.ErrHelp
		},
	}

	return root.ParseAndRun(context.Background(), args)
}

func server() *ffcli.Command {
	return &ffcli.Command{
		Name:       "server",
		ShortUsage: "basic server ${port}",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Exec: func(_ context.Context, _ []string) error {
			lis, err := net.Listen("tcp", "127.0.0.1:9314")
			if err != nil {
				return err
			}
			grpcServer := grpc.NewServer()

			demomod.RegisterDemomodSvcServer(grpcServer, demomod.New("test"))

			return grpcServer.Serve(lis)
		},
	}
}
