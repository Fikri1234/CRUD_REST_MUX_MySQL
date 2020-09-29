package service

import (
	"CRUD_REST_MUX_MySQL/model"
	"CRUD_REST_MUX_MySQL/repository"
	"CRUD_REST_MUX_MySQL/util"
	"encoding/json"
	"io/ioutil"
	"log"
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

var userDtl model.UserDetail
var userDtls []model.UserDetail

// GetUserDetailByID by id
func GetUserDetailByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	varID, err := strconv.ParseInt(params["id"], 10, 64)
	log.Print("userDtls")
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	userDtl, err := repository.GetUserDetailByID(varID)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.ResponseWithJSON(w, http.StatusOK, userDtl)
}

// GetUserDetails data
func GetUserDetails(w http.ResponseWriter, r *http.Request) {

	userDtls, err := repository.GetUserDetails()
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.ResponseWithJSON(w, http.StatusOK, userDtls)
}

// CreateUserDetail post
func CreateUserDetail(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal([]byte(body), &userDtl)

	userDtl, err := repository.CreateUserDetail(userDtl)
	if err != nil {
		util.ResponseWithError(w, http.StatusConflict, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusCreated, userDtl)
}

// UpdateUserDetail put
func UpdateUserDetail(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.Unmarshal([]byte(body), &userDtl)

	usrDtl, err := repository.UpdateUserDetail(userDtl)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusOK, usrDtl)
}

// DeleteUserDetailByID ...
func DeleteUserDetailByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	varID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		panic(err.Error())
	}

	err = repository.DeleteUserDetailByID(varID)
	if err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.ResponseWithJSON(w, http.StatusNoContent, userDtl)
}
