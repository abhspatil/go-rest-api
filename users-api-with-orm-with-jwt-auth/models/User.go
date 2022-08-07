package models

import "time"

type User struct {
	Id       int       `json:"id" gorm:"primaryKey"`
	Email    string    `gorm:"unique" json:"email"`
	FullName string    `json:"fullName"`
	DOB      time.Time `json:"dob"`
	Salary   float32   `json:"salary"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
}
