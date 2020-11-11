# REST-Golang-MUX-MySQL
Web service CRUD using Golang with gorilla-MUX for create REST api, MySQL as database, Viper as environment variable and sqlmock assert for Testing.


**Prerequisites**

1. [Go](https://golang.org/)
2. [Gorilla Mux](https://github.com/gorilla/mux)
3. [Mysql](https://www.mysql.com/downloads/)
4. [Viper](github.com/spf13/viper)
5. [SQLMock](github.com/DATA-DOG/go-sqlmock)
6. [Assert](github.com/stretchr/testify/assert)


**Getting Started**
1. Firstly, we need to get MUX, MySQL, Viper, sqlmock, assert library dependencies and install it
```
go get github.com/gorilla/mux  
go get github.com/go-sql-driver/mysql
go get github.com/spf13/viper
go get github.com/DATA-DOG/go-sqlmock
go get github.com/stretchr/testify/assert
```
2. Import dump.sql to your MySQL and configure your credential in folder resource
3. Open cmd in your project directory and type `go test -v` , you should get a response similar to the following:
![Alt text](asset/unitTesting.PNG?raw=true "Response Unit Testing")

4. To run application,open cmd in your project directory and type
```
go run main.go
```

**Sample Payload**
1. [Get User By Id](asset/getUserById.PNG)
2. [Get User Detail By Id](asset/getUserDetailById.PNG)
3. [Get All User](asset/getAllUser.PNG)
4. [Get All User Detail](asset/getAllUserDetail.PNG)
5. [Create User](asset/createUser.PNG)
6. [Create User Detail](asset/createUserDetail.PNG)
7. [Update User](asset/updateUser.PNG)
8. [Update User Detail](asset/updateUserDetail.PNG)
9. [Delete User By Id](asset/deleteUserById.PNG)
10. [Delete User Detail By Id](asset/deleteUserDetailById.PNG)
11. [Example error response,in case Update User Detail](asset/updateUserDetailError.PNG)