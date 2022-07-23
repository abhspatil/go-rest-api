package dto

import "time"

type CreateUserReq struct {
	FullName string    `json:"fullName"`
	DOB      time.Time `json:"dob"`
	Salary   float32   `json:"salary"`
}

type UpdateUserReq struct {
	Id       int       `json:"id"`
	FullName string    `json:"fullName"`
	DOB      time.Time `json:"dob"`
	Salary   float32   `json:"salary"`
}
