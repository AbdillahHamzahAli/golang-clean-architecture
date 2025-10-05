package repository

import (
	"context"

	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/domain/entity"
)

type PostRepository interface {
	Create(context.Context, entity.Post) (entity.Post, error)
	Update(context.Context, entity.Post) (entity.Post, error)
	GetById(context.Context, string) (entity.Post, bool, error)
	Delete(context.Context, string) (bool, error)
}