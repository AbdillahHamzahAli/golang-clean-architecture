package dto

import "time"

type UserResponse struct {
	ID        string    `json:"id,omitempty"`
	Username  string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Token     string    `json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required,max=100"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,max=100"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,max=100"`
	Password string `json:"password" binding:"required,max=100"`
}
