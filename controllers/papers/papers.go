package papers

import (
	"fmt"
	"github.com/baliance/gooxml/document"
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/helpers"
	"github.com/jeffjiang0613/question-bank/models"
	"strings"
)

type CreatePaperForm struct {
	Datas []struct {
		ID string	`json:"id" binding:"required"`
		Count int `json:"count" binding:"required"`
	} `json:"datas"`
}

func Create(ctx *gin.Context)  {
	form := CreatePaperForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		helpers.JsonFailureRespond(ctx,"参数错误")
	} else {
		var sqls []string
		for _,value :=range form.Datas {
			if value.Count > 0 {
				sqls = append(sqls, fmt.Sprintf("select * from(select * from questions where type_id = %s limit %d)",value.ID,value.Count))
			}
		}
		if len(sqls) > 0 {
			sql := fmt.Sprintf("select q.raw raw,q.raw_answer raw_answer,qt.title question_type_title, qt.sort question_type_sort  from (%s) q left join question_types qt on q.type_id = qt.id;",strings.Join(sqls," union "))
			var q []models.Question3
			models.DB.Raw(sql).Scan(&q)
			GeneratePaper(q,ctx)
		} else {
			helpers.JsonFailureRespond(ctx,"没有出题")
		}

	}
}

func GeneratePaper(questions []models.Question3, ctx *gin.Context)  {
	doc := document.New()
	ctx.Header("content-disposition", `attachment; filename=` + "a.docx")
	doc.Save(ctx.Writer)
}

