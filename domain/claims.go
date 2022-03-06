package domain

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username       string   `json:"username"`
	CustomerId     string   `json:"customer_id"`
	Role           string   `json:"role"`
	AccountNumbers []string `json:"account_numbers"`
	Exp            int64    `json:"exp"`
}

func BuildClaimsFromMap(mapClaims jwt.MapClaims) (*Claims, error) {
	// Convert map to json string
	jsonStr, err := json.Marshal(mapClaims)
	if err != nil {
		return nil, err
	}

	// Convert json string to struct
	var userClaims Claims
	if unmarshallErr := json.Unmarshal(jsonStr, &userClaims); err != nil {
		return nil, unmarshallErr
	}

	return &userClaims, nil
}
