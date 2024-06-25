package repository

import (
	"github.com/jmoiron/sqlx"
	"ims-authentication-api/model"
)

type Authorization interface {
	CreateUser(user model.User) (int64, error)
	GetUser(login string) (model.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
