package service

import (
	"ims-authentication-api/models"
	"ims-authentication-api/pkg/repository"
)

type Entity struct {
}

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Services struct {
	Entity
	Authorization
}

func NewService(repos *repository.Repository) *Services {
	return &Services{}
}
