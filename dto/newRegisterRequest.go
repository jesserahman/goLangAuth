package dto

type NewRegisterRequest struct {
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	Role       string  `json:"role"`
	CustomerId *string `json:"customer_id"`
}
