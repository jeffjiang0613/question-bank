package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/controllers/banks"
	"github.com/jeffjiang0613/question-bank/controllers/papers"
	"github.com/jeffjiang0613/question-bank/controllers/question_types"
	"github.com/jeffjiang0613/question-bank/controllers/questions"
	"github.com/jeffjiang0613/question-bank/controllers/ueditors"
)

func Route(app *gin.Engine) {
	app.Static("/static", "./static")

	top := app.Group("")
	banksGroup := top.Group("/banks")
	{
		banksGroup.GET("/:id", banks.Show)
		banksGroup.POST("", banks.Create)
		banksGroup.PATCH("", banks.Update)
		banksGroup.GET("", banks.List)
		banksGroup.DELETE("/:id", banks.Delete)
	}

	questionTypesGroup := top.Group("/question_types")
	{
		questionTypesGroup.POST("", question_types.Create)
		questionTypesGroup.PATCH("", question_types.Update)
		questionTypesGroup.GET("/:bank_id", question_types.List)
		questionTypesGroup.DELETE("/:id", question_types.Delete)
	}

	questionsGroup := top.Group("/questions")
	{
		questionsGroup.POST("", questions.Create)
		questionsGroup.PATCH("", questions.Update)
		questionsGroup.GET("", questions.List)
		questionsGroup.DELETE("/:id", questions.Delete)
	}

	papersGroup := top.Group("/papers")
	{
		papersGroup.POST("", papers.Create)
	}

	ueditorGroup := top.Group("/ueditors")
	{
		ueditorGroup.Any("", ueditors.Process)
	}

}
