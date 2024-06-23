package models

type User struct {
	Id       int64  `json:"-"`
	Login    string `json:"login" binding:"required" db:"login"`
	Password string `json:"password" binding:"required" db:"password"`
	RoleId   int64  `json:"role_id" binding:"required" db:"role_id"`
}

type Role struct {
	Id   int64  `json:"-"`
	Name string `json:"name" binding:"required" db:"name"`
}
