package action

import (
	"fmt"
)

type Sum struct {
	A int `json:"a"`
	B int `json:"b"`
}

func SumAction(s *Sum) (string, error) {
	if s == nil {
		return "", fmt.Errorf("`sum` field must be specified")
	}
	return fmt.Sprintf("%d", s.A+s.B), nil
}

func SayHelloAction() (string, error) {
	return "Hello World!", nil
}
