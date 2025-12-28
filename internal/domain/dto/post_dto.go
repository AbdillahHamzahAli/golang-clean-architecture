package dto

import (
	"time"

	"github.com/google/uuid"
)

type PostResponse struct {
	ID        string    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type PostCreateRequest struct {
	Title   string    `json:"title" binding:"required,max=100"`
	Content string    `json:"content" binding:"required,max=100"`
	UserID  uuid.UUID `json:"user_id" binding:"required"`
}

type PostUpdateRequest struct {
	ID      string    `json:"id" binding:"required"`
	Title   string    `json:"title" binding:"required,max=100"`
	Content string    `json:"content" binding:"required,max=100"`
	UserID  uuid.UUID `json:"user_id" binding:"required"`
}
