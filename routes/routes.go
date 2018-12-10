package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/controllers/banks"
	"github.com/jeffjiang0613/question-bank/controllers/papers"
	"github.com/jeffjiang0613/question-bank/controllers/questions"
	"github.com/jeffjiang0613/question-bank/controllers/ueditors"
)

func Route(app *gin.Engine)  {
	app.Static("/static","./static")

	top := app.Group("/v1")
	banksGroup := top.Group("/banks")
	{
		banksGroup.POST("",banks.Create)
		banksGroup.PATCH("",banks.Update)
		banksGroup.GET("",banks.List)
	}
	questionsGroup := top.Group("/questions")
	{
		questionsGroup.POST("",questions.Create)
		questionsGroup.PATCH("",questions.Update)
		questionsGroup.GET("",questions.List)
	}

	papersGroup := top.Group("/papers")
	{
		papersGroup.GET("",papers.Create)
	}

	ueditorGroup := top.Group("/ueditors")
	{
		ueditorGroup.Any("",ueditors.Process)
	}

}
