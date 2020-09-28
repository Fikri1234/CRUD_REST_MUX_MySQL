package model

// User struct
type User struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Employee *Employee `json:"employee"`
}

// Users array type
type Users []User
