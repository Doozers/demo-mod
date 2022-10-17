package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
)

var opts struct {
	cmd string
	arg string
}

func main() {
	if err := demoMod(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
	return
}

func demoMod(args []string) error {
	rootFlagSet := flag.NewFlagSet("demo-mod", flag.ExitOnError)

	root := &ffcli.Command{
		FlagSet:    rootFlagSet,
		ShortUsage: "demo-mod [flags] <command> [args...]",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Subcommands: []*ffcli.Command{
			sum(),
		},
		Exec: func(_ context.Context, _ []string) error {
			return flag.ErrHelp
		},
	}

	return root.ParseAndRun(context.Background(), args)
}

func sum() *ffcli.Command {
	var A int64
	var B int64

	sumFs := flag.NewFlagSet("sum", flag.ExitOnError)
	sumFs.Int64Var(&A, "A", 0, "first number")
	sumFs.Int64Var(&B, "B", 0, "second number")

	return &ffcli.Command{
		Name:       "sum",
		ShortUsage: "demo-mod sum ${A} ${B}",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		FlagSet:    sumFs,
		Exec: func(_ context.Context, _ []string) error {
			c, err := getClient()
			if err != nil {
				return err
			}

			result, err := sumAction(c, A, B)
			if err != nil {
				return err
			}

			fmt.Printf("%d + %d = %d", A, B, result)
			return nil
		},
	}
}
