package main

import (
	"log"

	"github.com/devstackq/go-clean/config"
	"github.com/devstackq/go-clean/server"
	"github.com/spf13/viper"
)

func main() {
	if err := config.GetConfig(); err != nil {
		log.Println(err, "viper")
		return
	}
	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Println(err)
		return
	}
}
