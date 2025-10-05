package middleware

import (
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/usecase"
	"gorm.io/gorm"
)

type Middleware struct {
	db          *gorm.DB
	jwtUC       usecase.JwtUsecase
}

func New(db *gorm.DB, jwtUC usecase.JwtUsecase) Middleware {
	return Middleware{db, jwtUC}
}