package usecase

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtUsecase interface {
	GenerateUserToken(userId string) (string, error)
	ValidateUserToken(token string) (*jwt.Token, error)
}

type jwtUserClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type jwtUsecase struct {
	secretKey string
	issuer    string
}

func NewJwtUsecase() JwtUsecase {
	return &jwtUsecase{
		secretKey: getSecretKey(),
		issuer:    "template",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "Template"
	}
	return secretKey
}

func (j *jwtUsecase) GenerateUserToken(userId string) (string, error) {
	claims := jwtUserClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 120)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tx, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}
	return tx, nil
}

func (j *jwtUsecase) parseToken(t_ *jwt.Token) (any, error) {
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
	}
	return []byte(j.secretKey), nil
}

func (j *jwtUsecase) ValidateUserToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, j.parseToken)
}
