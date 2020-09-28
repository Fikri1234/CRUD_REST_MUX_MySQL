package model

// UserDetail struct
type UserDetail struct {
	ID      int64  `json:"id"`
	Address string `json:"address"`
	DOB     string `json:"dob"`
	POB     string `json:"pob"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	UserID  int64  `json:"userId"`
}

// UserDetails array type
type UserDetails []UserDetail
