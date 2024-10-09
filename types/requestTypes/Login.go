package requestTypes

import "notes-back/types"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  types.User   `json:"user"`
}