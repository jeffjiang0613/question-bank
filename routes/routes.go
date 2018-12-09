package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/controllers/questions"
)

func Route(app *gin.Engine)  {
	top := app.Group("/v1")

	questionsGroup := top.Group("/questions")
	{
		questionsGroup.POST("",questions.Create)
		questionsGroup.PATCH("",questions.Update)
		questionsGroup.GET("",questions.List)
	}
}
