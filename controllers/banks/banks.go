package banks

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/helpers"
	"github.com/jeffjiang0613/question-bank/models"
)

type CreateBankForm struct {
	Name string	`json:"name"`
	Coment string	`json:"coment"`
}

type UpdateBankForm struct {
	ID uint	`json:"id"`
	Name string	`json:"name"`
	Coment string	`json:"coment"`
}

func Create(ctx *gin.Context)  {
	form := CreateBankForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		helpers.JsonFailureRespond(ctx,"参数错误")
	} else {
		bank := models.NewBank(form.Name,form.Coment)
		bank.Save()
		helpers.JsonSuccessfulRespond(ctx,models.AllBanks())
	}
}

func List(ctx *gin.Context)  {
	helpers.JsonSuccessfulRespond(ctx,models.AllBanks())
}

func Update(ctx *gin.Context)  {
	form := UpdateBankForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		helpers.JsonFailureRespond(ctx,"参数错误")
	} else {
		bank,exits := models.GetBankById(form.ID)
		if exits {
			bank.Name = form.Name
			bank.Coment = form.Coment
			bank.Save()
			helpers.JsonSuccessfulRespond(ctx,models.AllBanks())
		} else {
			helpers.JsonFailureRespond(ctx,"题库不存在")
		}
	}
}
