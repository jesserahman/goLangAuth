package service

import (
	"github.com/jesserahman/goLangAuth/domain"
	"github.com/jesserahman/goLangAuth/dto"
	"github.com/jesserahman/goLangAuth/errs"
)

type RegisterService interface {
	RegisterUser(request dto.NewRegisterRequest) *errs.AppError
}

type RegisterServiceImpl struct {
	repository domain.RegisterRepository
}

func (service RegisterServiceImpl) RegisterUser(request dto.NewRegisterRequest) *errs.AppError {
	err := service.repository.RegisterNewUser(request)
	if err != nil {
		return err
	}
	return nil
}

func NewRegisterService(repo domain.RegisterRepository) RegisterServiceImpl {
	return RegisterServiceImpl{repo}
}
