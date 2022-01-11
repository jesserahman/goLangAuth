package app

import (
	"fmt"
	"goLangAuth/service"
	"net/http"
)

type AuthHandler struct {
	service service.AuthService
}

func (handler *AuthHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
