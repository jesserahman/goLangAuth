package service

import (
	"fmt"
	"github.com/jesserahman/goLangAuth/domain"
	"github.com/jesserahman/goLangAuth/dto"
	"github.com/jesserahman/goLangAuth/errs"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
)

type AuthService interface {
	VerifyCredentials(dto.NewAuthRequest) (*string, *errs.AppError)
	VerifyToken(urlParams map[string]string) (bool, error)
}

type AuthServiceImpl struct {
	repository      domain.AuthRepository
	rolePermissions domain.RolePermissions
}

func (service AuthServiceImpl) VerifyCredentials(request dto.NewAuthRequest) (*string, *errs.AppError) {
	user, err := service.repository.CheckCredentials(request)
	if err != nil {
		return nil, err
	}

	var claims jwt.MapClaims
	if user.Role == "admin" {
		claims = user.GenerateAdminClaims()
	} else {
		claims = user.GenerateUserClaims()
	}

	// newWithClaims takes the signing method in this case it's HS256 and then generates the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token with the secret
	signedTokenAsString, signStringErr := token.SignedString([]byte(os.Getenv("HMAC_SAMPLE_SECRET")))
	if signStringErr != nil {
		fmt.Println("error signing token")
	}

	return &signedTokenAsString, nil
}

func (service AuthServiceImpl) VerifyToken(urlParams map[string]string) (bool, error) {
	// get Token from the URL and convert the string token to JWT struct
	if jwtToken, err := jwtTokenFromString(urlParams["token"]); err != nil {
		return false, err
	} else {

		//	check the validity of the token: expiration time & signature
		if jwtToken.Valid {

			// get mapClaims from JWT token
			mapClaims := jwtToken.Claims.(jwt.MapClaims)

			// convert map to Claim struct - Claims contains all the user info that was sent in the token
			userClaims, claimsErr := domain.BuildClaimsFromMap(mapClaims)
			if claimsErr != nil {
				fmt.Println("error: ", err.Error())
				return false, err
			}

			// verify user permissions
			isAuthorized := service.rolePermissions.IsAuthorizedFor(userClaims.Role, urlParams["routeName"])
			if !isAuthorized {
				userUnauthorizedError := fmt.Errorf("user does not have permission to access this route")
				fmt.Println(userUnauthorizedError)
				return false, userUnauthorizedError
			} else {

				if userClaims.Role == "user" {
					if urlParams["customer_id"] == userClaims.CustomerId {
						return true, nil
					} else {
						userUnauthorizedError := fmt.Errorf("customer ID in URL does not match customer ID in JWT token")
						fmt.Println(userUnauthorizedError)
						return false, userUnauthorizedError
					}
				}
				return true, nil
			}
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

func NewAuthService(repo domain.AuthRepository, permissions domain.RolePermissions) AuthServiceImpl {
	return AuthServiceImpl{repo, permissions}
}
