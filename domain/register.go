package domain

import (
	"goLangAuth/dto"
	"goLangAuth/errs"
)

type RegisterRepository interface {
	RegisterNewUser(request dto.NewRegisterRequest) *errs.AppError
}
