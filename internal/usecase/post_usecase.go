package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/domain/dto"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/domain/entity"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/repository"
)

type PostUsecase interface {
    Create(ctx context.Context, request dto.PostCreateRequest) (dto.PostResponse, error)
    Update(ctx context.Context, request dto.PostUpdateRequest) (dto.PostResponse, error)
    Delete(ctx context.Context, id string) (bool, error)
    GetById(ctx context.Context, id string) (dto.PostResponse, bool, error)	
}

type postUsecase struct {
    postRepo repository.PostRepository
    userRepo repository.UserRepository
}

func NewPostUsecase(postRepo repository.PostRepository, userRepo repository.UserRepository) PostUsecase {
    return &postUsecase{
        postRepo: postRepo,
        userRepo: userRepo,
    }
}

func (p *postUsecase) Create(ctx context.Context, request dto.PostCreateRequest) (dto.PostResponse, error) {
    post := entity.Post{
        Title:     request.Title,
        Content:   request.Content,
        UserID:    request.UserID,
    }

    created, err := p.postRepo.Create(ctx, post)
    if err != nil {
        return dto.PostResponse{}, err
    }

    return dto.PostResponse{
        ID:        created.ID.String(),
        Title:     created.Title,
        Content:   created.Content,
        UserID:    created.UserID.String(),
        CreatedAt: created.CreatedAt,
        UpdatedAt: created.UpdatedAt,
    }, nil
}

func (p *postUsecase) Update(ctx context.Context, request dto.PostUpdateRequest) (dto.PostResponse, error) {
    fmt.Println(request.ID)
    findPost, flag, err := p.postRepo.GetById(ctx, request.ID)
    
    if err != nil {
        return dto.PostResponse{}, err
    }
    if !flag {
        return dto.PostResponse{}, errors.New("post not found")
    }

    updatePost := entity.Post{
        ID:        findPost.ID,
        Title:     request.Title,
        Content:   request.Content,
        UserID:    findPost.UserID,
    }
    
    updated, err := p.postRepo.Update(ctx, updatePost)
    if err != nil {
        return dto.PostResponse{}, err
    }

    return dto.PostResponse{
        ID:        updated.ID.String(),
        Title:     updated.Title,
        Content:   updated.Content,
        UserID:    updated.UserID.String(),
        CreatedAt: updated.CreatedAt,
        UpdatedAt: updated.UpdatedAt,
    }, nil
}

func (p *postUsecase) Delete(ctx context.Context, id string) (bool, error) {
    if id == "" {
        return false, errors.New("missing id")
    }
    ok, err := p.postRepo.Delete(ctx, id)
    if err != nil {
        return false, err
    }
    return ok, nil
}

func (p *postUsecase) GetById(ctx context.Context, id string) (dto.PostResponse, bool, error) {
    res, found, err := p.postRepo.GetById(ctx, id)
    if err != nil {
        return dto.PostResponse{}, false, err
    }
    if !found {
        return dto.PostResponse{}, false, nil
    }
    return dto.PostResponse{
        ID:        res.ID.String(),
        Title:     res.Title,
        Content:   res.Content,
        UserID:    res.UserID.String(),
        CreatedAt: res.CreatedAt,
        UpdatedAt: res.UpdatedAt,
    }, true, nil
}
