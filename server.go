package authx

import (
	"authx/apis"
	"authx/models"
	"authx/settings"
	"log"
	"os"
)

// addr --> ip:port --> "127.0.0.1:8080"
func Run(addr string) {
	// init database
	_, err := models.SetEngine(&models.Config{
		Obj:      settings.Get("db_obj"),
		User:     settings.Get("db_user"),
		Password: settings.Get("db_password"),
		Host:     settings.Get("db_host"),
		Name:     settings.Get("db_name"),
		Log:      settings.Get("db_log"),
	})
	if err != nil {
		log.Fatal("authx链接数据库失败:", err)
		os.Exit(-1)
	}
	// run api server
	apis.NewAPIServer().G.Run(addr)
}
