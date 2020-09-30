package main

import (
	"CRUD_REST_MUX_MySQL/configuration"
	"CRUD_REST_MUX_MySQL/router"
)

func main() {
	configuration.ReadConfig()
	router.WebServerConf("8999")
}
