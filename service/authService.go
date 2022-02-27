package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"goLangAuth/domain"
	"goLangAuth/dto"
	"goLangAuth/errs"
	"log"
	"os"
	"time"
)

type AuthService interface {
	VerifyCredentials(dto.NewAuthRequest) (*string, *errs.AppError)
	VerifyToken(urlParams map[string]string) (bool, error)
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
		"exp":         time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedTokenAsString, signStringErr := token.SignedString([]byte(os.Getenv("HMAC_SAMPLE_SECRET")))
	if signStringErr != nil {
		fmt.Println("error signing token")
	}

	return &signedTokenAsString, nil
}

func (service AuthServiceImpl) VerifyToken(urlParams map[string]string) (bool, error) {
	// convert string token to JWT struct
	if jwtToken, err := jwtTokenFromString(urlParams["token"]); err != nil {
		return false, err
	} else {
		//	check the validity of the token: expiration time & signature
		if jwtToken.Valid {
			return true, nil
		}
	}

	return false, nil
}

func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("HMAC_SAMPLE_SECRET")), nil
	})

	if err != nil {
		log.Println("Error while parsing token: ", err.Error())
		return nil, err
	}
	return token, nil
}

func NewAuthService(repo domain.AuthRepository) AuthServiceImpl {
	return AuthServiceImpl{repo}
}
