package domain

import (
	"goLangAuth/dto"
	"goLangAuth/errs"
)

type User struct {
	Username   string `db:"username"`
	Password   string `db:"password"`
	Role       string `db:"role"`
	CustomerId string `db:"customer_id"`
	CreatedOn  string `db:"created_on"`
}
type AuthRepository interface {
	CheckCredentials(dto.NewAuthRequest) (*User, *errs.AppError)
}
