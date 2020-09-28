package router

import (
	"CRUD_REST_MUX_MySQL/service"

	"github.com/gorilla/mux"
)

// NewRouter group of router from routes
// func NewRouter() *mux.Router {

// 	router := mux.NewRouter().StrictSlash(true)
// 	for _, route := range routesUser {
// 		router.Methods(route.Method).
// 			Path(route.Pattern).
// 			Name(route.Name).
// 			Handler(route.HandlerFunc)
// 	}

// 	return router
// }

// NewRouter list of router endpoint
func NewRouter() *mux.Router {

	r := mux.NewRouter()

	// handle user endpoint
	r.HandleFunc("/api/user/{id}", service.GetUser).Methods("GET")
	r.HandleFunc("/api/user/", service.GetUsers).Methods("GET")
	r.HandleFunc("/api/user/", service.CreateUser).Methods("POST")
	r.HandleFunc("/api/user/", service.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/user/{id}", service.DeleteUserByID).Methods("DELETE")

	// handle user detail endpoint
	r.HandleFunc("/api/user/dtl/{id}", service.GetUserDetailByID).Methods("GET")
	r.HandleFunc("/api/user/dtl/", service.GetUserDetails).Methods("GET")
	r.HandleFunc("/api/user/dtl/", service.CreateUserDetail).Methods("POST")
	r.HandleFunc("/api/user/dtl/", service.UpdateUserDetail).Methods("PUT")
	r.HandleFunc("/api/user/dtl/{id}", service.DeleteUserDetailByID).Methods("DELETE")

	return r
}
