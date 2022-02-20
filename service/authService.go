package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"goLangAuth/domain"
	"goLangAuth/dto"
	"goLangAuth/errs"
	"os"
)

type AuthService interface {
	VerifyCredentials(dto.NewAuthRequest) (*string, *errs.AppError)
}

type AuthServiceImpl struct {
	repository domain.AuthRepository
}

func (service AuthServiceImpl) VerifyCredentials(request dto.NewAuthRequest) (*string, *errs.AppError) {
	user, err := service.repository.CheckCredentials(request)
	if err != nil {
		return nil, err
	}

	var claims = jwt.MapClaims{
		"customer_id": user.CustomerId,
		"username":    user.Username,
		"role":        user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedTokenAsString, signStringErr := token.SignedString([]byte(os.Getenv("HMAC_SAMPLE_SECRET")))
	if signStringErr != nil {
		fmt.Println("error signing token")
	}

	return &signedTokenAsString, nil
}

func NewAuthService(repo domain.AuthRepository) AuthServiceImpl {
	return AuthServiceImpl{repo}
}
