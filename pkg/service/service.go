package service

import (
	"ims-authentication-api/model"
	"ims-authentication-api/pkg/repository"
)

type Entity struct {
}

type Authorization interface {
	CreateUser(user model.User) (int64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int64, error)
}

type Services struct {
	Entity
	Authorization
}

func NewService(repos *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(repos.Authorization),
	}
}
