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

// handleLogin will verify that the token is valid
func (handler *AuthHandler) handleVerify(w http.ResponseWriter, r *http.Request) {
	urlParams := make(map[string]string)

	for k, _ := range r.URL.Query() {
		urlParams[k] = r.URL.Query().Get(k)
	}

	if urlParams["token"] != "" {
		isAuthorized, appError := handler.service.VerifyToken(urlParams)
		if appError != nil {
			writeResponse(w, http.StatusForbidden, notAuthorizedResponse("verify token error"))
		} else {
			if isAuthorized {
				writeResponse(w, http.StatusOK, authorizedResponse())
			} else {
				writeResponse(w, http.StatusForbidden, notAuthorizedResponse("unauthorized"))
			}
		}
	} else {
		writeResponse(w, http.StatusForbidden, notAuthorizedResponse("missing token"))
	}
}

func notAuthorizedResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"isAuthorized": false,
		"message":      msg,
	}
}

func authorizedResponse() map[string]bool {
	return map[string]bool{"isAuthorized": true}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
