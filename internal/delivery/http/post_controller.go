package http

import (
	"net/http"

	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/domain/dto"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/shared/response"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostController struct {
    UseCase usecase.PostUsecase
}

func NewPostController(useCase usecase.PostUsecase) *PostController {
    return &PostController{
        UseCase: useCase,
    }
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	uid := ctx.MustGet("user_id").(string)
    
	var post dto.PostCreateRequest
	
    err := ctx.ShouldBindBodyWithJSON(&post)
    if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request body", err.Error())
        return
    }
	
	post.UserID = uuid.MustParse(uid)
    resp, err := pc.UseCase.Create(ctx.Request.Context(), post)
    if err != nil {
        response.Error(ctx, http.StatusBadRequest, "create post failed", err.Error())
        return
    }
    response.Success(ctx, http.StatusOK, "create post success", resp)
}

func (pc *PostController) UpdatePost(ctx *gin.Context) {
	uid := ctx.MustGet("user_id").(string)
	id := ctx.Param("id")
    
	var post dto.PostUpdateRequest
	
    err := ctx.ShouldBindBodyWithJSON(&post)
    if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request body", err.Error())
        return
    }
	
	post.UserID = uuid.MustParse(uid)
    post.ID = id
    resp, err := pc.UseCase.Update(ctx.Request.Context(), post)
    if err != nil {
        response.Error(ctx, http.StatusBadRequest, "update post failed", err.Error())
        return
    }
    response.Success(ctx, http.StatusOK, "update post success", resp)
}

func (pc *PostController) DeletePost(ctx *gin.Context) {
    id := ctx.Param("id")
    
    resp, err := pc.UseCase.Delete(ctx.Request.Context(), id)
    if err != nil {
        response.Error(ctx, http.StatusBadRequest, "delete post failed", err.Error())
        return
    }
    response.Success(ctx, http.StatusOK, "delete post success", resp)
}

func (pc *PostController) GetPostById(ctx *gin.Context) {
    
	id := ctx.Param("id")
	
    resp, _, err := pc.UseCase.GetById(ctx.Request.Context(), id)
    if err != nil {
        response.Error(ctx, http.StatusBadRequest, "get post failed", err.Error())
        return
    }
    response.Success(ctx, http.StatusOK, "get post success", resp)
}
