package httputil

import "github.com/gin-gonic/gin"

func NewError(ctx *gin.Context, status int, err error) {
	error := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, error)
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
