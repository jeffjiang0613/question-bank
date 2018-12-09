package papers

import (
	"github.com/baliance/gooxml/document"
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/helpers"
	"github.com/jeffjiang0613/question-bank/models"
	"strings"
)

type CreatePaperForm struct {
	Datas []struct {
		Type int `json:"type" binding:"required"`
		Count int `json:"count" binding:"required"`
		Name string `json:"name"`
	} `json:"datas"`
}

func Create(ctx *gin.Context)  {
	form := CreatePaperForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		helpers.JsonFailureRespond(ctx,"参数错误")
	} else {
		var sqls []string
		var params []interface{}
		for _,value :=range form.Datas {
			if value.Count > 0 {
				sqls = append(sqls, "(select * from questions where type = ? limit ?)")
				params = append(params, value.Type)
				params = append(params, value.Count)
			}
		}
		if len(sqls) > 0 {
			sql := strings.Join(sqls," union ")+ ";"
			var q []models.Question
			models.DB.Raw(sql,params...).Scan(&q)
			GeneratePaper(q,ctx)
		} else {
			helpers.JsonFailureRespond(ctx,"没有出题")
		}

	}
}

func GeneratePaper(questions []models.Question, ctx *gin.Context)  {
	doc := document.New()
	ctx.Header("content-disposition", `attachment; filename=` + "a.docx")
	doc.Save(ctx.Writer)
}

