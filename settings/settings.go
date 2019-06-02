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
		"authx_host":  "localhost",
		"authx_port":  "9993",
	}
}
func Get(key string) string {
	env := strings.TrimSpace(os.Getenv(strings.ToUpper(key)))
	if env != "" {
		return env
	}
	return defaults[key]
}
