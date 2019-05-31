package main

import (
	"authx/apis"
	"os"

	"authx/models"

	"authx/settings"
)

func main() {
	// init database
	_, err := models.SetEngine(&models.Config{
		User:     settings.Get("db_user"),
		Password: settings.Get("db_password"),
		Host:     settings.Get("db_host"),
		Name:     settings.Get("db_name"),
		Log:      settings.Get("db_log"),
	})
	if err != nil {
		os.Exit(-1)
	}
	// run api server
	apis.NewAPIServer().G.Run(":9991")
}
