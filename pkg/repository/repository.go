package repository

import (
	"github.com/jmoiron/sqlx"
	"ims-authentication-api/models"
)

type Authorization interface {
	CreateUser(user models.User) (int64, error)
	GetUser(login string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
