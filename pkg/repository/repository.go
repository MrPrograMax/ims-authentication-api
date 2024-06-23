package repository

import "github.com/jmoiron/sqlx"

type Entity struct {
}

type Repository struct {
	Entity
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
