package dto

type NewAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
