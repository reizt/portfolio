package app

import (
	"fmt"
)

type SayHelloInput struct {
	Name string
}

func (a *App) SayHello(input SayHelloInput) (*string, error) {
	sentence := fmt.Sprintf("Hello, %s!", input.Name)
	return &sentence, nil
}
