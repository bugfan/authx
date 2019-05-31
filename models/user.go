package models

import "time"

func init() {
	Register(new(User))
}

type User struct {
	ID       int64
	Name     string
	Password string
	Created  time.Time
	Updated  time.Time
}
