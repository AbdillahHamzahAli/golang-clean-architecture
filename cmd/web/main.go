package main

import (
	"fmt"
	"os"

	"github.com/AbdillahHamzahAli/golang-clean-architecture/cmd"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/config"
	deliveryhttp "github.com/AbdillahHamzahAli/golang-clean-architecture/internal/delivery/http"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/delivery/http/middleware"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/delivery/http/route"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/infrastructure/pgsql"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var (
		db *gorm.DB = config.SetUpDatabase()
	)

	server := gin.Default()
	if len(os.Args) > 1 {
		fmt.Println("Running command")
		cmd.Commands(db)
		return
	}

	// Repositories
	userRepo := pgsql.NewUserRepoPG(db)
	postRepo := pgsql.NewPostRepoPG(db)

	// Usecases
	jwtUC := usecase.NewJwtUsecase()
	authUC := usecase.NewAuthUsecase(userRepo, jwtUC)
	postUC := usecase.NewPostUsecase(postRepo, userRepo)

	// Controllers
	userController := deliveryhttp.NewUserController(authUC)
	postController := deliveryhttp.NewPostController(postUC)

	// Middleware
	mw := middleware.New(db, jwtUC)

	// Routes
	route.UserRoute(server, *userController)
	route.PostRoute(server, *postController, mw)

	err := server.Run(":8000")
	if err != nil {
		return
	}
}
