package dao

import (
	"database/sql"

	"github.com/pkg/errors"
)

var (
	UserNotFound = errors.New("user not found")
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
}

func (u *User) MockDataNotFound() (err error) {
	return UserNotFound
}

func (u *User) Detail() (err error) {
	query := "SELECT id,name FROM users where id = ? limit 1"
	err = DB.QueryRow(query, u.ID).Scan(&u.ID, &u.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.Wrap(err, u.MockDataNotFound().Error())
		}
		return
	}
	return
}
