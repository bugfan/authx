package models

import (
	"authx/utils"
	"time"
)

func init() {
	Register(new(User))
}

func initDefaultUser() {
	user := User{
		Name:     "admin",
		Password: "123456",
	}
	has, _ := x.Where("name=?", user.Name).Exist(&user)
	if !has {
		x.Insert(&user)
	}
}

type User struct {
	ID       int64
	Name     string
	Password string
	JWT      string
	Created  time.Time
	Updated  time.Time
}

func (s *User) Exist() bool {
	s.Password = utils.EncryptedPassword(s.Password)
	has, _ := x.Exist(s)
	return has
}
