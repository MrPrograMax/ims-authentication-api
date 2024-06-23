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
	insertQuery := "INSERT INTO users (login, password, role_id) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(insertQuery)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(user.Login, user.Password, user.Role)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *AuthPostgres) GetUser(login string) (models.User, error) {
	var user models.User

	query := "SELECT login, password, role_id FROM users WHERE login = $1"
	err := r.db.Get(user, query, login)

	return user, err
}
