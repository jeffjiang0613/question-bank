package models

import (
	"github.com/jeffjiang0613/question-bank/helpers"
)

const CountPerPage = 10

type Question struct {
	Model
	Content   string `gorm:"index" json:"-"`
	Raw       string `gorm:"index" json:"question"`
	Answer    string `gorm:"index" json:"-"`
	RawAnswer string `gorm:"index" json:"answer"`
	BankId    uint   `json:"bank_id"`
	TypeId      uint   `json:"type_id"`
}

type Question2 struct {
	Model
	Content   string `gorm:"index" json:"-"`
	Raw       string `gorm:"index" json:"question"`
	Answer    string `gorm:"index" json:"-"`
	RawAnswer string `gorm:"index" json:"answer"`
	BankId    uint   `json:"bank_id"`
	TypeId      uint   `json:"type_id"`
	BankName string	`json:"bank_name"`
	TypeName string	`json:"type_name"`
}

type Question3 struct {
	Raw       string
	RawAnswer string
	QuestionTypeTitle string
	QuestionTypeSort uint
}

type AllQuestionsResult struct {
	PageCount int	`json:"page_count"`
	Questions []Question2	`json:"questions"`
}


func NewQuestion(raw string, rawAnswer string, bankId uint, questionTypeId uint) *Question {
	q := new(Question)
	q.Raw = helpers.ComprassHtml(raw)
	q.Content = q.Raw
	q.RawAnswer = rawAnswer
	q.Answer = q.RawAnswer
	q.BankId = bankId
	q.TypeId = questionTypeId
	return q
}

func (question *Question) UpdateRawAnswerType(raw string, rawAnswer string, questionTypeId uint) {
	question.Raw = helpers.ComprassHtml(raw)
	question.Content = question.Raw
	question.RawAnswer = rawAnswer
	question.Answer = rawAnswer
	question.TypeId = questionTypeId
}

func (question *Question) UpdateRawAnswer(raw string, rawAnswer string) {
	question.Raw = helpers.ComprassHtml(raw)
	question.Content = raw
	question.RawAnswer = rawAnswer
	question.Answer = rawAnswer
}

func (question *Question) Save() (err error) {
	if DB.NewRecord(question) {
		err = DB.Create(question).Error
	} else {
		err = DB.Save(question).Error
	}
	return
}

func AllQuestions(page, perPage int) AllQuestionsResult {
	var aqr AllQuestionsResult
	var count int
	var questions []Question2
	DB.Model(&Question{}).Count(&count)
	aqr.PageCount = count / perPage
	if count % perPage != 0 {
		aqr.PageCount++
	}
	DB.Raw("select q.*,b.name bank_name,qt.name type_name from questions as q left join banks as b inner join question_types as qt on q.bank_id = b.id and q.type_id=qt.id order by id desc limit ? offset ?",perPage,page*perPage).Scan(&questions)
	aqr.Questions = questions
	return aqr
}

func AllQuestionsByTypeCount(page, perPage int) AllQuestionsResult {
	var aqr AllQuestionsResult
	var count int
	var questions []Question2
	DB.Model(&Question{}).Count(&count)
	aqr.PageCount = count / perPage
	if count % perPage != 0 {
		aqr.PageCount++
	}
	DB.Raw("select q.*,b.name bank_name,qt.name type_name from questions as q left join banks as b inner join question_types as qt on q.bank_id = b.id and q.type_id=qt.id order by id desc limit ? offset ?",perPage,page*perPage).Scan(&questions)
	aqr.Questions = questions
	return aqr
}

func GetQuestionById(id uint) (question *Question, exists bool) {
	question = new(Question)
	if DB.First(question, "id = ?", id).RecordNotFound() {
		return question, false
	}
	return question, true
}

func DeleteQuestionById(id string) error {
	if err := DB.Delete(&Question{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
