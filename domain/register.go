package domain

import (
	"github.com/jesserahman/goLangAuth/dto"
	"github.com/jesserahman/goLangAuth/errs"
)

type RegisterRepository interface {
	RegisterNewUser(request dto.NewRegisterRequest) *errs.AppError
}
