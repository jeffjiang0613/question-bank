package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonSuccessfulRespond(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   data,
	})
	ctx.Abort()
}

func JsonFailureRespond(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status":  "failed",
		"message": msg,
	})
	ctx.Abort()
}
