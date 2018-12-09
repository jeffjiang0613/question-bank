package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/controllers/banks"
)

func Route(app *gin.Engine)  {
	top := app.Group("/v1")

	banksGroup := top.Group("/banks")
	{
		banksGroup.POST("",banks.Create)
		banksGroup.PATCH("",banks.Update)
		banksGroup.GET("",banks.List)
	}
}


