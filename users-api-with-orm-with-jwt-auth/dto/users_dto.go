package dto

import "time"

type CreateUserReq struct {
	FullName string    `json:"fullName"`
	DOB      time.Time `json:"dob"`
	Salary   float32   `json:"salary"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Email    string    `json:"email"`
}

type UpdateUserReq struct {
	Id       int       `json:"id"`
	FullName string    `json:"fullName"`
	DOB      time.Time `json:"dob"`
	Salary   float32   `json:"salary"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Email    string    `json:"email"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	TokenString string `json:"token"`
}
