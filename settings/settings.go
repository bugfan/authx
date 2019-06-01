package settings

import (
	"os"
	"strings"
)

var defaults map[string]string

func init() {
	defaults = map[string]string{
		"db_user":     "root",
		"db_password": "123456",
		"db_host":     "127.0.0.1:3306",
		"db_name":     "authx",
		"db_log":      "xorm.log",
		"mongo_host":  "127.0.0.1",
		"admin_host":  "localhost",
	}
}
func Get(key string) string {
	env := strings.TrimSpace(os.Getenv(key))
	if env != "" {
		return env
	}
	return defaults[key]
}
