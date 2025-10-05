package pgsql

import (
	"errors"

	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/domain/entity"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/repository"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepoPG(db *gorm.DB) repository.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepo) GetById(username string) (entity.User, bool, error) {
	var user entity.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, false, nil
		}
		return entity.User{}, false, err
	}
	return user, true, nil
}

func (r *userRepo) GetByEmail(email string) (entity.User, bool, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, false, nil
		}
		return entity.User{}, false, err
	}
	return user, true, nil
}
