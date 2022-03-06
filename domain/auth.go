package domain

import (
	"goLangAuth/dto"
	"goLangAuth/errs"
)

type User struct {
	Username   string  `db:"username"`
	CustomerId *string `db:"customer_id"`
	Role       string  `db:"role"`
	AccountIDs *string `db:"account_ids"`
}
type AuthRepository interface {
	CheckCredentials(dto.NewAuthRequest) (*User, *errs.AppError)
}
