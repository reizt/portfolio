package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/reizt/portfolio/src/core"
	"github.com/reizt/portfolio/src/core/istore"
)

// Implements istore.IUserStore
type userStore struct {
	db *sql.DB
}

func (s userStore) FindMany(input istore.IUserStoreFindManyInput) (*([]core.User), error) {
	users := make([]core.User, 0)

	query := fmt.Sprintf(`SELECT * FROM %s`, userTableName)
	rows, err := s.db.Query(query)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	for rows.Next() {
		user := &core.User{}
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			log.Fatal(err.Error())
			return nil, err
		}
		users = append(users, *user)
	}

	return &users, nil
}
