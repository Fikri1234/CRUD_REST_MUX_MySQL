# REST-Golang-MUX-MySQL
Web service CRUD using Golang with gorilla-MUX for create REST api and MySQL as database


**Prerequisites**

1. [Go](https://golang.org/)
2. [Gorilla Mux](https://github.com/gorilla/mux)
3. [Mysql](https://www.mysql.com/downloads/)


**Getting Started**
1. Firstly, we need to get MUX and MySQL library dependencies and install it
```
go get github.com/gorilla/mux  
go get github.com/go-sql-driver/mysql
```
2. Import dump.sql to your MySQL
3. Open cmd in your project directory and run command 
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