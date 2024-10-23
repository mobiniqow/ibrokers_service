package basics

import "github.com/gin-gonic/gin"

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, APIError{Code: code, Message: message})
}
