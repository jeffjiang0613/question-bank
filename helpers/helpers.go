package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

func ComprassHtml(rawl string) string  {
	reader := strings.NewReader(rawl)
	leftString := strings.Builder{}
	var previous rune = 0
	for true {
		ch ,_ ,err := reader.ReadRune()
		if err != nil {
			return leftString.String()
		}
		if (ch == ' ' || ch == '\t' || ch == '\n') {
			if !(previous == ' ' || previous == '\t' || previous == '\n') {
				leftString.WriteRune(' ')
			}
		} else {
			leftString.WriteRune(ch)
		}
		previous = ch
	}
	return leftString.String()
}
