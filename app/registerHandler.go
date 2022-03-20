package app

import (
	"encoding/json"
	"github.com/jesserahman/goLangAuth/dto"
	"github.com/jesserahman/goLangAuth/service"
	"net/http"
)

type RegisterHandler struct {
	service service.RegisterService
}

// handleLogin will verify that the username and password matches the DB
func (handler *RegisterHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var request dto.NewRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		err := handler.service.RegisterUser(request)
		if err != nil {
			writeResponse(w, http.StatusUnauthorized, err)
		}
	}
}
