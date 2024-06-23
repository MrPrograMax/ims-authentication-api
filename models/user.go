package models

type User struct {
	Id       int64  `json:"-"`
	Login    string `json:"username" binding:"required" db:"login"`
	Password []byte `json:"password" binding:"required" db:"password"`
	Role     int64  `json:"role" binding:"required" db:"role_id"`
}

type Role struct {
	Id   int64  `json:"-"`
	Name string `json:"name" binding:"required" db:"name"`
}
