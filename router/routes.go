package router

import (
	"CRUD_REST_MUX_MySQL/service"
	"net/http"
)

// Route function name, HTTP method, url that will execute, method service
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes define of array Route
type Routes []Route

var routesUser = Routes{
	Route{
		"GetUser",
		"GET",
		"/api/user/{id}",
		service.GetUser,
	},
	Route{
		"GetUsers",
		"GET",
		"/api/user/",
		service.GetUsers,
	},
	Route{
		"CreateUser",
		"POST",
		"/api/user/",
		service.CreateUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/api/user/",
		service.UpdateUser,
	},
	Route{
		"DeleteUserByID",
		"DELETE",
		"/api/user/{id}",
		service.DeleteUserByID,
	},
}
