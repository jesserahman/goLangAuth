package service

import "goLangAuth/domain"

type AuthService interface {
}

type AuthServiceImpl struct {
	repository domain.AuthRepository
}

func NewAuthService(repo domain.AuthRepository) AuthServiceImpl {
	return AuthServiceImpl{repo}
}
