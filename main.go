package main

import (
	"CRUD_REST_MUX_MySQL/configuration"
	"CRUD_REST_MUX_MySQL/router"

	"github.com/spf13/viper"
)

func main() {
	configuration.ReadConfig()
	port := viper.GetString("PORT")
	router.WebServerConf(port)
}
