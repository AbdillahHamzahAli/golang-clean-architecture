package response

import (
	"github.com/gin-gonic/gin"
)

type Meta struct {
    Code    int    `json:"code"`
    Message string `json:"message,omitempty"`
}

type Envelope struct {
    Meta   Meta `json:"meta"`
    Data   any `json:"data,omitempty"`
    Errors any `json:"errors,omitempty"`
}

func Success(ctx *gin.Context, httpStatus int, message string, data any) {
    ctx.JSON(httpStatus, Envelope{
        Meta: Meta{Code: httpStatus, Message: message},
        Data: data,
    })
}

func Error(ctx *gin.Context, httpStatus int, message string, err any) {
    ctx.JSON(httpStatus, Envelope{
        Meta:   Meta{Code: httpStatus, Message: message},
        Errors: err,
    })
}

type Pagination struct {
    Page       int   `json:"page"`
    PageSize   int   `json:"page_size"`
    TotalItems int64 `json:"total_items"`
    TotalPages int   `json:"total_pages"`
}

type PaginatedEnvelope struct {
    Meta       Meta        `json:"meta"`
    Data       any         `json:"data"`
    Pagination *Pagination `json:"pagination"`
}

func Paginated(ctx *gin.Context, httpStatus int, message string, data any, p *Pagination) {
    ctx.JSON(httpStatus, PaginatedEnvelope{
        Meta:       Meta{Code: httpStatus, Message: message},
        Data:       data,
        Pagination: p,
    })
}
