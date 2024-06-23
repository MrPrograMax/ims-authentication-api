package repository

import (
	"database/sql"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

//func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
//	var id int
//	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
//
//	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
//	if err := row.Scan(&id); err != nil {
//		return 0, err
//	}
//
//	return id, nil
//}
