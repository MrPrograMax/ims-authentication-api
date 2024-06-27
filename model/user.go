package model

type User struct {
	Id       int64  `json:"-"`
	FullName string `json:"full_name" binding:"required" db:"full_name"`

	Login    string `json:"login" binding:"required" db:"login"`
	Password string `json:"password" binding:"required" db:"password"`
	RoleId   int64  `json:"role_id" binding:"required" db:"role_id"`
}

type Role struct {
	Id   int64  `json:"-"`
	Name string `json:"name" binding:"required" db:"name"`
}
