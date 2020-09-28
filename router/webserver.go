package router

import (
	"log"
	"net/http"
)

var client = &http.Client{}

// WebServerConf ...
func WebServerConf(port string) {
	// ==== old ways config from router.go
	r := NewRouter()
	http.Handle("/", r)

	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}
	client.Transport = transport

	// // here is service router
	// userRouter := service.UserRouter()
	// userDtlRouter := service.UserDetailRouter()

	// http.Handle("/", userRouter)
	// http.Handle("/", r)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
