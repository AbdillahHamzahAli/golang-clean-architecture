package route

import (
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/delivery/http"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
)

func PostRoute(r *gin.Engine, controller http.PostController, mw middleware.Middleware) {
	r.GET("/post", controller.GetPosts)
	r.POST("/post", middleware.Auth(&mw), controller.CreatePost)
	r.PUT("/post/:id", middleware.Auth(&mw), controller.UpdatePost)
	r.DELETE("/post/:id", middleware.Auth(&mw), controller.DeletePost)
	r.GET("/post/:id", middleware.Auth(&mw), controller.GetPostById)
}
