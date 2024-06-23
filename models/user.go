package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"-"`
	Login    string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Role     string    `json:"role" binding:"required"`
	Salt     string    `json:"salt" binding:"required"`
}
