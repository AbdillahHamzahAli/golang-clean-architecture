package pgsql

import (
	"context"
	"errors"

	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/domain/entity"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/repository"
	"gorm.io/gorm"
)

type postRepo struct {
	db *gorm.DB
}

func NewPostRepoPG(db *gorm.DB) repository.PostRepository {
	return &postRepo{
		db: db,
	}
}

func (r *postRepo) Create(ctx context.Context, post entity.Post) (entity.Post, error) {
	if err := r.db.WithContext(ctx).Create(&post).Error; err != nil {
		return entity.Post{}, err
	}
	return post, nil
}

func (r *postRepo) Update(ctx context.Context, post entity.Post) (entity.Post, error) {
	if err := r.db.WithContext(ctx).Save(&post).Error; err != nil {
		return entity.Post{}, err
	}
	return post, nil
}

func (r *postRepo) Delete(ctx context.Context, id string) (bool, error) {
	if err := r.db.WithContext(ctx).Delete(&entity.Post{}, "id = ?", id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *postRepo) GetById(ctx context.Context, id string) (entity.Post, bool, error) {
	var post entity.Post
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Post{}, false, nil
		}
		return entity.Post{}, false, err
	}
	return post, true, nil
}
