package banks

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/helpers"
	"github.com/jeffjiang0613/question-bank/models"
	"strconv"
)

type CreateBankForm struct {
	Name          string                `json:"name"`
	Coment        string                `json:"coment"`
	QuestionTypes []models.QuestionType `json:"question_types"`
}

type UpdateBankForm struct {
	ID            uint                  `json:"id"`
	Name          string                `json:"name"`
	Coment        string                `json:"coment"`
	QuestionTypes []models.QuestionType `json:"question_types"`
}

func Create(ctx *gin.Context) {
	form := CreateBankForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		helpers.JsonFailureRespond(ctx, "参数错误")
	} else {
		var qts []models.QuestionType
		for _, qt := range form.QuestionTypes {
			qts = append(qts, models.QuestionType{
				Name:  qt.Name,
				Title: qt.Title,
				Sort:  qt.Sort,
			})
		}
		bank := models.NewBank(form.Name, form.Coment, qts)
		if err := bank.Save(); err != nil {
			helpers.JsonFailureRespond(ctx, "创建失败")
		} else {
			helpers.JsonSuccessfulRespond(ctx, "创建成功")
		}
	}
}

func List(ctx *gin.Context) {
	helpers.JsonSuccessfulRespond(ctx, models.AllBanks())
}

func Show(ctx *gin.Context) {
	bankId := ctx.Param("id")
	bank, exist := models.GetBankById(bankId)
	if !exist {
		helpers.JsonFailureRespond(ctx, "题库不存在")
	} else {
		helpers.JsonSuccessfulRespond(ctx, bank)
	}
}

func Update(ctx *gin.Context) {
	form := UpdateBankForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		helpers.JsonFailureRespond(ctx, "参数错误")
	} else {
		bank, exits := models.GetBankById(strconv.Itoa(int(form.ID)))
		if exits {
			var qts []models.QuestionType
			for _, qt := range form.QuestionTypes {
				t := models.QuestionType{
					Name:  qt.Name,
					Title: qt.Title,
					Sort:  qt.Sort,
				}
				t.ID = qt.ID
				qts = append(qts, t)
			}
			bank.Name = form.Name
			bank.Coment = form.Coment
			bank.QuestionTypes = qts
			if err := bank.Update(); err != nil {
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
	if err := models.DeleteBankById(id); err != nil {
		helpers.JsonFailureRespond(ctx, "删除失败")
	} else {
		helpers.JsonSuccessfulRespond(ctx, "删除成功")
	}
}
