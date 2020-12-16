package biz

import (
	"jmolboy/week04/internal/dao"
)

type User struct{}

func (U *User) Create(user dao.User) (insertId int, err error) {
	return user.Create()
}
