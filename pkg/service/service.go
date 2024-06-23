package service

import "ims-authentication-api/pkg/repository"

type Entity struct {
}

type Services struct {
	Entity
}

func NewService(repos *repository.Repository) *Services {
	return &Services{}
}
