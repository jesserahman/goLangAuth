package domain

import (
	"goLangAuth/dto"
	"goLangAuth/errs"
)

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
type AuthRepository interface {
	CheckCredentials(dto.NewAuthRequest) *errs.AppError
}
