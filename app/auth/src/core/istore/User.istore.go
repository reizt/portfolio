package istore

import (
	"github.com/reizt/portfolio/src/core"
)

type IUserStore interface {
	FindMany(input IUserStoreFindManyInput) (*([]core.User), error)
}

type IUserStoreFindManyInput struct {
	//
}
