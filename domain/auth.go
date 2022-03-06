package domain

import (
	"github.com/golang-jwt/jwt"
	"goLangAuth/dto"
	"goLangAuth/errs"
	"time"
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

func (u User) GenerateAdminClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"username": u.Username,
		"role":     u.Role,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}
}

func (u User) GenerateUserClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"username":        u.Username,
		"customer_id":     u.CustomerId,
		"role":            u.Role,
		"account_numbers": u.AccountIDs,
		"exp":             time.Now().Add(time.Hour).Unix(),
	}
}
