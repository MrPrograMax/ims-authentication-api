package repository

import (
	"github.com/jmoiron/sqlx"
	"ims-authentication-api/model"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int64, error) {
	var id int64
	insertQuery := "INSERT INTO user (full_name, login, password, role_id) VALUES ($1, $2, $3) RETURNING id"

	row := r.db.QueryRow(insertQuery, user.Login, user.Password, user.RoleId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(login string) (model.User, error) {
	var user model.User

	query := "SELECT full_name, login, password, role_id FROM users WHERE login = $1"
	err := r.db.Get(&user, query, login)

	return user, err
}
