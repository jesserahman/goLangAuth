package app

import (
	"encoding/json"
	"goLangAuth/dto"
	"goLangAuth/service"
	"net/http"
)

type AuthHandler struct {
	service service.AuthService
}

// handleLogin will verify that the username and password matches the DB
func (handler *AuthHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAuthRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		signedTokenAsString, err := handler.service.VerifyCredentials(request)
		if err != nil {
			writeResponse(w, http.StatusUnauthorized, err)
		} else {
			writeResponse(w, http.StatusOK, signedTokenAsString)
		}
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
