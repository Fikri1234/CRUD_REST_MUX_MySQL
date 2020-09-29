package repository

import (
	"CRUD_REST_MUX_MySQL/configmysql"
	"CRUD_REST_MUX_MySQL/model"
	"log"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	_ "github.com/go-sql-driver/mysql"
)

// var db *sql.DB
// var err error

// var client = &http.Client{}

// func initialize() {
// 	db = configmysql.Connect()

// 	// defer db.Close()

// 	var transport http.RoundTripper = &http.Transport{
// 		DisableKeepAlives: true,
// 	}
// 	client.Transport = transport
// }

// GetUserByID ...
func GetUserByID(id int64) (model.User, error) {
	db := configmysql.Connect()

	var user model.User

	result, err := db.Query("select id, user_name, password from m_user where id = ?", id)
	if err != nil {
		// print stack trace
		log.Println("Error query user: " + err.Error())
		return user, err
	}

	defer db.Close()

	for result.Next() {
		err := result.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

// GetUsers ...
func GetUsers() ([]model.User, error) {
	db := configmysql.Connect()

	var user model.User
	var users []model.User

	rows, err := db.Query("select id, user_name, password from m_user")
	if err != nil {
		log.Print(err)
	}

	defer db.Close()

	var isError bool
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			isError = true
			break
		} else {
			users = append(users, user)
		}
	}

	if isError {
		return users, err
	}

	return users, nil
}

// CreateUser ...
func CreateUser(usr model.User) (model.User, error) {
	db := configmysql.Connect()
	defer db.Close()

	var user model.User

	crt, err := db.Prepare("insert into m_user (user_name, password) values (?, ?)")
	if err != nil {
		return user, err
	}
	res, err := crt.Exec(usr.Username, usr.Password)
	if err != nil {
		return user, err
	}

	rowID, err := res.LastInsertId()
	if err != nil {
		return user, err
	}

	user.ID = int64(rowID)

	// find user by id
	resval, err := GetUserByID(user.ID)
	if err != nil {
		return user, err
	}

	return resval, nil
}

// UpdateUser ...
func UpdateUser(usr model.User) (model.User, error) {
	db := configmysql.Connect()
	defer db.Close()

	var user model.User

	crt, err := db.Prepare("update m_user set user_name =?, password =? where id=?")
	if err != nil {
		return user, err
	}
	_, queryError := crt.Exec(usr.Username, usr.Password, usr.ID)
	if queryError != nil {
		return user, err
	}

	// find user by id
	res, err := GetUserByID(usr.ID)
	if err != nil {
		return user, err
	}

	return res, nil
}

// DeleteUserByID ...
func DeleteUserByID(id int64) error {
	db := configmysql.Connect()
	defer db.Close()

	crt, err := db.Prepare("delete from m_user where id=?")
	if err != nil {
		return err
	}
	_, queryError := crt.Exec(id)
	if queryError != nil {
		return err
	}

	return nil
}
