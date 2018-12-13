package question_types

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/helpers"
	"github.com/jeffjiang0613/question-bank/models"
)

type CreateQuestionTypeForm struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	BankId uint   `json:"bank_id"`
}

type UpdateQuestionTypeForm struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	BankId uint   `json:"bank_id"`
}

func Create(ctx *gin.Context) {
	form := CreateQuestionTypeForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		helpers.JsonFailureRespond(ctx, "参数错误")
	} else {
		qt := models.NewQuestionType(form.Name, form.Title, form.BankId)
		if err := qt.Save(); err != nil {
			helpers.JsonFailureRespond(ctx, "创建失败")
		} else {
			helpers.JsonSuccessfulRespond(ctx, "创建成功")
		}
	}
}

func List(ctx *gin.Context) {
	bankId := ctx.Param("bank_id")
	helpers.JsonSuccessfulRespond(ctx, models.AllQuestionTypes(bankId))
}

func Update(ctx *gin.Context) {
	form := UpdateQuestionTypeForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		helpers.JsonFailureRespond(ctx, "参数错误")
	} else {
		qt, exits := models.GetQuestionTypeById(form.ID)
		if exits {
			qt.Name = form.Name
			qt.BankId = form.BankId
			if err := qt.Save(); err != nil {
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
	if err := models.DeleteQuestionTypeById(id); err != nil {
		helpers.JsonFailureRespond(ctx, "删除失败")
	} else {
		helpers.JsonSuccessfulRespond(ctx, "删除成功")
	}
}
