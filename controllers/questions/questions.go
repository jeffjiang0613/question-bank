package questions

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/helpers"
	"github.com/jeffjiang0613/question-bank/models"
)

type CreateQuestionForm struct {
	Raw    string	`gorm:"index" json:"content"`
	RawAnswer string	`gorm:"index" json:"answer"`
	BankId uint	`json:"bank_id"`
	Type    uint	`json:"type"`
}

type UpdateQuestionForm struct {
	ID uint	`json:"id"`
	Raw    string	`gorm:"index" json:"content"`
	RawAnswer string	`gorm:"index" json:"answer"`
	Type    uint	`json:"type"`
}

func Create(ctx *gin.Context)  {
	form := CreateQuestionForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		helpers.JsonFailureRespond(ctx,"参数错误")
	} else {
		question := models.NewQuestion(form.Raw,form.RawAnswer,form.BankId,form.Type)
		question.Save()
		helpers.JsonSuccessfulRespond(ctx,models.AllQuestions())
	}
}

func List(ctx *gin.Context)  {
	helpers.JsonSuccessfulRespond(ctx,models.AllQuestions())
}

func Update(ctx *gin.Context)  {
	form := UpdateQuestionForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		helpers.JsonFailureRespond(ctx,"参数错误")
	} else {
		question,exits := models.GetQuestionById(form.ID)
		if exits {
			question.UpdateRawAnswerType(form.Raw,form.RawAnswer,form.Type)
			question.Save()
			//helpers.JsonSuccessfulRespond(ctx,models.AllBanks())
		} else {
			helpers.JsonFailureRespond(ctx,"题库不存在")
		}
	}
}