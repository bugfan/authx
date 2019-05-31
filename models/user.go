package models

import (
	"authx/utils"
	"time"
)

func init() {
	Register(new(User))
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
