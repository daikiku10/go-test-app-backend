package model

import "time"

type User struct {
	ID             int64     `json:"id" db:"id"`
	FirstName      string    `json:"first_name" db:"first_name"`
	FirstNameKana  string    `json:"first_name_kana" db:"first_name_kana"`
	FamilyName     string    `json:"family_name" db:"family_name"`
	FamilyNameKana string    `json:"family_name_kana" db:"family_name_kana"`
	Password       string    `json:"password" db:"password"`
	Email          string    `json:"email" db:"email"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdateAt       time.Time `json:"updateAt" db:"update_at"`
}

type Users []*User
