package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{
			"status": "running!",
		})
	}
}