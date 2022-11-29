package handler

import (
	"encoding/json"

	"github.com/reizt/portfolio/src/adapters/handler/spec"
	"github.com/reizt/portfolio/src/core"
	"github.com/reizt/portfolio/src/core/app"
)

var (
	sayHello400Sentence = "invalid input"
	sayHello500Sentence = "an error occured"
)

func (h *handler) SayHello(c Ctx) error {
	// Make input
	body, err := body[spec.SayHelloRequest](c)
	if err != nil {
		if _, ok := err.(*json.SyntaxError); ok {
			resBody := spec.SayHello400Response{Error: spec.SayHello400ResponseError{Message: sayHello400Sentence}}
			return (*c.Status(400)).SendJson(resBody)
		}
		resBody := spec.SayHello500Response{Error: spec.SayHello500ResponseError{Message: sayHello500Sentence}}
		return (*c.Status(500)).SendJson(resBody)
	}
	input := app.SayHelloInput{
		Name: body.Name,
	}

	// Do action
	sentence, err := h.app.SayHello(input)

	// Make output
	if err != nil {
		if err == core.ErrInvalidInput {
			resBody := spec.SayHello400Response{Error: spec.SayHello400ResponseError{Message: sayHello400Sentence}}
			return (*c.Status(400)).SendJson(resBody)
		}
		resBody := spec.SayHello500Response{Error: spec.SayHello500ResponseError{Message: sayHello500Sentence}}
		return (*c.Status(500)).SendJson(resBody)
	}

	resBody := spec.SayHello200Response{Data: spec.SayHello200ResponseData{Sentence: *sentence}}
	return (*c.Status(200)).SendJson(resBody)
}
