package service

import (
	"goLangAuth/domain"
	"goLangAuth/dto"
	"goLangAuth/errs"
)

type AuthService interface {
	VerifyCredentials(dto.NewAuthRequest) *errs.AppError
}

type AuthServiceImpl struct {
	repository domain.AuthRepository
}

func (service AuthServiceImpl) VerifyCredentials(request dto.NewAuthRequest) *errs.AppError {
	err := service.repository.CheckCredentials(request)
	if err != nil {
		return err
	}
	return nil
}

func NewAuthService(repo domain.AuthRepository) AuthServiceImpl {
	return AuthServiceImpl{repo}
}
