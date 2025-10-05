package route

import (
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/delivery/http"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine, controller http.UserController) {
	r.POST("/auth/login", controller.Login)
	r.POST("/auth/register", controller.Register)
}
