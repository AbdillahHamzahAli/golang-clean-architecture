package repository

import (
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/domain/entity"
)

type UserRepository interface {
	Create(entity.User) (entity.User, error)
	GetById(string) (entity.User, bool, error)
	GetByEmail(string) (entity.User, bool, error)
}
