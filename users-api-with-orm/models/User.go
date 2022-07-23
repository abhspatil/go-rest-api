package models

import "time"

type User struct {
	Id       int       `json:"id" gorm:"primaryKey"`
	FullName string    `json:"fullName"`
	DOB      time.Time `json:"dob"`
	Salary   float32   `json:"salary"`
}
