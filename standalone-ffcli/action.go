package main

import (
	"os"
)

func SumAction(A, B int64) int64 {
	return A + B
}

func SayHelloAction() string {
	return "hello! " + os.Getenv("USER")
}

func EchoStreamAction(sent string) string {
	return sent
}
