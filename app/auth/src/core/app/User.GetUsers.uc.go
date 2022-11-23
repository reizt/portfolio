package app

import (
	"github.com/reizt/portfolio/src/core"
	"github.com/reizt/portfolio/src/core/istore"
)

type GetUsersInput struct {
	//
}

func (a *App) GetUsers(input GetUsersInput) (*([]core.User), error) {
	users, err := a.Store.User.FindMany(istore.IUserStoreFindManyInput{})
	if err != nil {
		return nil, err
	}
	return users, nil
}
