package service

import (
	"CRUD_REST_MUX_MySQL/model"
	"CRUD_REST_MUX_MySQL/repository"
	"CRUD_REST_MUX_MySQL/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var user model.User
var users []model.User

// GetUser by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	varID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := repository.GetUserByID(varID)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.ResponseWithJSON(w, http.StatusOK, user)
}

// GetUsers data
func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := repository.GetUsers()
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusOK, users)
}

// CreateUser post
func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal([]byte(body), &user)

	user, err := repository.CreateUser(user)
	if err != nil {
		// print stack trace
		// panic(err.Error())
		util.ResponseWithError(w, http.StatusConflict, err.Error())
		return
	}

	// json.NewEncoder(w).Encode(user)
	util.ResponseWithJSON(w, http.StatusOK, user)
}

// UpdateUser post
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.Unmarshal([]byte(body), &user)

	user, err := repository.UpdateUser(user)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusOK, user)
}

// DeleteUserByID ...
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	varID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		panic(err.Error())
	}

	err = repository.DeleteUserByID(varID)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusNoContent, user)
}
