package questions

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/helpers"
	"github.com/jeffjiang0613/question-bank/models"
	"strconv"
)

type CreateQuestionForm struct {
	Raw       string `gorm:"index" json:"question"`
	RawAnswer string `gorm:"index" json:"answer"`
	BankId    uint   `json:"bank_id"`
	TypeId      uint   `json:"type_id"`
}

type UpdateQuestionForm struct {
	ID        uint   `json:"id"`
	Raw       string `gorm:"index" json:"question"`
	RawAnswer string `gorm:"index" json:"answer"`
	BankId    uint   `json:"bank_id"`
	TypeId      uint   `json:"type_id"`
}

func Create(ctx *gin.Context) {
	form := CreateQuestionForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		helpers.JsonFailureRespond(ctx, "参数错误")
	} else {
		question := models.NewQuestion(form.Raw, form.RawAnswer, form.BankId, form.TypeId)
		if question.Save() != nil {
			helpers.JsonFailureRespond(ctx, "创建失败")
		} else {
			helpers.JsonSuccessfulRespond(ctx, "创建成功")
		}
	}
}

func List(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 0
	}
	helpers.JsonSuccessfulRespond(ctx, models.AllQuestions(page, 5))
}

func Update(ctx *gin.Context) {
	form := UpdateQuestionForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		helpers.JsonFailureRespond(ctx, "参数错误")
	} else {
		question, exits := models.GetQuestionById(form.ID)
		if exits {
			question.UpdateRawAnswerType(form.Raw, form.RawAnswer, form.TypeId)
			if question.Save() != nil {
				helpers.JsonFailureRespond(ctx, "更新失败")
			} else {
				helpers.JsonSuccessfulRespond(ctx, "更新成功")
			}
		} else {
			helpers.JsonFailureRespond(ctx, "题库不存在")
		}
	}
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := models.DeleteQuestionById(id); err != nil {
		helpers.JsonFailureRespond(ctx, "删除失败")
	} else {
		helpers.JsonSuccessfulRespond(ctx, "删除成功")
	}
}
