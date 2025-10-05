package usecase

import (
    "context"
    "errors"

    "github.com/AbdillahHamzahAli/golang-clean-architecture/internal/domain/dto"
    "github.com/AbdillahHamzahAli/golang-clean-architecture/internal/domain/entity"
    "github.com/AbdillahHamzahAli/golang-clean-architecture/internal/repository"
    "golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
    Register(ctx context.Context, request dto.RegisterUserRequest) (dto.UserResponse, error)
    Login(ctx context.Context, request dto.LoginUserRequest) (dto.UserResponse, error)
}

type authUsecase struct {
    repo       repository.UserRepository
    jwtUsecase JwtUsecase
}

func NewAuthUsecase(repo repository.UserRepository, jwtUsecase JwtUsecase) AuthUsecase {
    return &authUsecase{
        repo:       repo,
        jwtUsecase: jwtUsecase,
    }
}

func (u *authUsecase) Register(ctx context.Context, request dto.RegisterUserRequest) (dto.UserResponse, error) {
    _, flag, err := u.repo.GetByEmail(request.Email)
    if err != nil {
        return dto.UserResponse{}, err
    }

	if flag {
		return dto.UserResponse{}, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponse{}, err
	}

	user := entity.User{
		Username: request.Username,
		Password: string(hashedPassword),
		Email:    request.Email,
	}

	userReg, err := u.repo.Create(user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:       userReg.ID.String(),
		Username: userReg.Username,
		Email:    userReg.Email,
	}, nil

}

func (u *authUsecase) Login(ctx context.Context, request dto.LoginUserRequest) (dto.UserResponse, error) {
	check, flag, err := u.repo.GetByEmail(request.Email)
	if err != nil || !flag {
		return dto.UserResponse{}, errors.New("user does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(check.Password), []byte(request.Password))
	if err != nil {
		return dto.UserResponse{}, errors.New("wrong password")
	}

	token, err := u.jwtUsecase.GenerateUserToken(check.ID.String())

	return dto.UserResponse{
		Token: token,
	}, nil
}
