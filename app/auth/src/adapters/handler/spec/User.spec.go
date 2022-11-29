package spec

import (
	"github.com/reizt/portfolio/src/core"
)

type GetUsersRequest struct{}

type GetUsers200Response struct {
	Data []core.User `json:"data"`
}
