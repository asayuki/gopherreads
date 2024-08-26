package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type UserAuth struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}
