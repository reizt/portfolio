package handler

import "encoding/json"

type Ctx interface {
	Body() (*([]byte), error)
	Status(status int) *Ctx
	SendJson(body any) error
}

func body[T any](c Ctx) (*T, error) {
	bytes, err := (c).Body()
	if err != nil {
		return nil, err
	}
	b := new(T)
	if err := json.Unmarshal(*bytes, b); err != nil {
		return nil, err
	}
	return b, nil
}
