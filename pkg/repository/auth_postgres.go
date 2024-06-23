package repository

import (
	"github.com/jmoiron/sqlx"
	"ims-authentication-api/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int64, error) {
	var id int64
	insertQuery := "INSERT INTO users (login, password, role_id) VALUES ($1, $2, $3) RETURNING id"

	row := r.db.QueryRow(insertQuery, user.Login, user.Password, user.RoleId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(login string) (models.User, error) {
	var user models.User

	query := "SELECT login, password, role_id FROM users WHERE login = $1"
	err := r.db.Get(&user, query, login)

	return user, err
}
