package main

import (
	"CRUD_REST_MUX_MySQL/configuration"
	"CRUD_REST_MUX_MySQL/model"
	"CRUD_REST_MUX_MySQL/router"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/spf13/viper"
)

var db sql.DB
var client = &http.Client{}

func TestMain(m *testing.T) {
	viper.SetConfigFile("./resource/properties-test.yaml")

	viper.ReadInConfig()

}

var user = &model.User{
	ID:       1999999,
	Username: "haha",
	Password: "pass123",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock, error) {

	db, mock, err := sqlmock.New()

	return db, mock, err
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {

	rr := httptest.NewRecorder()

	r := router.NewRouter()
	r.Use(configuration.CORS)
	// http.Handle("/", r)

	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}
	client.Transport = transport

	r.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// test API
func TestAPIGetByIDUser(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/user/1", nil)
	if err != nil {
		t.Errorf("Expected sdf code %v. Go", err)
	}
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestAPIGetByIDUserNotFound(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/user/1999999999", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestAPIGetAllUser(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/user/", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

// func TestAPICreateUser(t *testing.T) {
// 	payload := []byte(`{"Username":"wiro", "password":"pass345"}`)

// 	req, _ := http.NewRequest("POST", "/api/user/", bytes.NewBuffer(payload))
// 	resp := executeRequest(req)

// 	checkResponseCode(t, http.StatusCreated, resp.Code)

// 	var m map[string]interface{}
// 	json.Unmarshal(resp.Body.Bytes(), &m)

// 	if m["Username"] != "wiro" {
// 		t.Errorf("Expected user name to be 'wiro'. Got '%v'", m["Username"])
// 	}

// 	if m["ID"] != 1 {
// 		t.Errorf("Expected product ID to be '1'. Got '%v'", m["ID"])
// 	}
// }

// Test repo
func TestFindUserById(t *testing.T) {
	_, mock, err := NewMock()
	if err != nil {
		fmt.Printf("error mock: " + err.Error())
	}

	// simulate any sql driver behavior in tests, without needing a real database connection
	query := "select id, user_name, password from m_user where id = \\?"

	rows := sqlmock.NewRows([]string{"id", "user_name", "password"}).
		AddRow(user.ID, user.Username, user.Password)

	mock.ExpectQuery(query).WithArgs(user.ID).WillReturnRows(rows)
	// ------------ end of mock ---------------

	assert.NotNil(t, user)
}

func TestFindUserByIdError(t *testing.T) {
	db, mock, err := NewMock()
	if err != nil {
		fmt.Printf("error mock: " + err.Error())
	}

	db = configuration.Connect()

	defer db.Close()

	// simulate any sql driver behavior in tests, without needing a real database connection
	query := "select id, user_name, password from m_user where id = \\?"

	rows := sqlmock.NewRows([]string{"id", "user_name", "password"}).
		AddRow(user.ID, user.Username, user.Password)

	mock.ExpectQuery(query).WithArgs(user.ID).WillReturnRows(rows)
	// ------------ end of mock ---------------

	// Context like a timeout or deadline or a channel to indicate stop working and return
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := new(model.User)
	err = db.QueryRowContext(ctx, "select id, user_name, password from m_user where id = ?", user.ID).Scan(&res.ID, &res.Username, &res.Password)

	assert.Empty(t, res)
	assert.Error(t, err)
}

func TestFindAllUser(t *testing.T) {
	users := make([]*model.User, 0)

	db, mock, err := NewMock()
	if err != nil {
		fmt.Printf("error mock: " + err.Error())
	}

	db = configuration.Connect()

	defer db.Close()

	// simulate any sql driver behavior in tests, without needing a real database connection
	query := "select id, user_name, password from m_user where id = ?"
	rows := sqlmock.NewRows([]string{"id", "user_name", "password"}).
		AddRow(user.ID, user.Username, user.Password)

	mock.ExpectQuery(query).WithArgs(user.ID).WillReturnRows(rows)
	// ------------ end of mock ---------------

	// Context like a timeout or deadline or a channel to indicate stop working and return
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := db.QueryContext(ctx, "select id, user_name, password from m_user")
	defer res.Close()

	for res.Next() {
		user := new(model.User)
		err = res.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
		)

		users = append(users, user)
	}

	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	// assert.Len(t, users, 1)
}

func TestCreateUser(t *testing.T) {
	_, mock, err := NewMock()
	if err != nil {
		fmt.Printf("error mock: " + err.Error())
	}

	query := "insert into m_user \\(user_name, password\\) values \\(\\?, \\?\\)"

	rows := mock.ExpectPrepare(query)
	rows.ExpectExec().WithArgs(user.Username, user.Password).WillReturnResult(sqlmock.NewResult(0, 0))
}

func TestUpdateUser(t *testing.T) {
	_, mock, err := NewMock()
	if err != nil {
		fmt.Printf("error mock: " + err.Error())
	}

	query := "update m_user set user_name =\\?, password =\\? where id =\\?"

	rows := mock.ExpectPrepare(query)
	rows.ExpectExec().WithArgs(user.Username, user.Password, user.ID).WillReturnResult(sqlmock.NewResult(0, 1))
}

func TestDeleteUser(t *testing.T) {
	_, mock, err := NewMock()
	if err != nil {
		fmt.Printf("error mock: " + err.Error())
	}

	query := "delete from m_user where id =\\?"

	rows := mock.ExpectPrepare(query)
	rows.ExpectExec().WithArgs(user.ID).WillReturnResult(sqlmock.NewResult(0, 1))
}
