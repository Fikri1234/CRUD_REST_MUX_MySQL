package model

// Employee struct
type Employee struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Position  string `json:"position"`
}

// Employees array struct
type Employees []Employee
