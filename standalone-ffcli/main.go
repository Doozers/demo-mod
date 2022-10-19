package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
)

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
			sayHello(),
			echoStream(),
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
			result := SumAction(A, B)

			fmt.Printf("%d + %d = %d", A, B, result)
			return nil
		},
	}
}

func sayHello() *ffcli.Command {
	return &ffcli.Command{
		Name:       "sayHello",
		ShortUsage: "demo-mod sayHello",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Exec: func(_ context.Context, _ []string) error {
			result := SayHelloAction()

			fmt.Println(result)
			return nil
		},
	}
}

func echoStream() *ffcli.Command {
	return &ffcli.Command{
		Name:       "echoStream",
		ShortUsage: "demo-mod echoStream",
		Options:    []ff.Option{ff.WithEnvVarNoPrefix()},
		Exec: func(_ context.Context, _ []string) error {
			receive := make(chan string)

			go func() {
				Reader := bufio.NewReader(os.Stdin)
				var input string
				for {
					input, _ = Reader.ReadString('\n')
					fmt.Print(" >> ")

					fmt.Println(len(input) - 1)
					if len(input[:len(input)-1]) > 1 {
						receive <- EchoStreamAction(input[:len(input)-1])
					} else {
						receive <- ""
						return
					}
				}
			}()

			go func() {
				for {
					text, ok := <-receive
					if !ok || text == "" {
						return
					}

					fmt.Println(text)
				}
			}()

			for runtime.NumGoroutine() > 1 {
				time.Sleep(1 * time.Second)
			}
			return nil
		},
	}

}
