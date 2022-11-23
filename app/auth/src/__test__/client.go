package __test__

import (
	"net/http"
)

func newClient() client {
	baseUrl := "http://localhost:6060"
	c := client{baseUrl}
	return c
}

type client struct {
	baseUrl string
}

func (c client) getUsers() (resp *http.Response, err error) {
	return http.Get(c.baseUrl + "/users")
}
