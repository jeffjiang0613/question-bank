package models

import "github.com/jeffjiang0613/question-bank/helpers"

type Question struct {
	Model
	Content   string `gorm:"index" json:"-"`
	Raw       string `gorm:"index" json:"question"`
	Answer    string `gorm:"index" json:"-"`
	RawAnswer string `gorm:"index" json:"answer"`
	BankId    uint   `json:"bank_id" json:"bank_id"`
	Type      uint   `json:"type"`
}

func NewQuestion(raw string, rawAnswer string, bankId uint, questionType uint) *Question {
	q := new(Question)
	q.Raw = helpers.ComprassHtml(raw)
	q.Content = q.Raw
	q.RawAnswer = q.RawAnswer
	q.Answer = rawAnswer
	q.BankId = bankId
	q.Type = questionType
	return q
}

func (question *Question) UpdateRawAnswerType(raw string, rawAnswer string, questionType uint) {
	question.Raw = helpers.ComprassHtml(raw)
	question.Content = question.Raw
	question.RawAnswer = question.RawAnswer
	question.Answer = rawAnswer
	question.Type = questionType
}

func (question *Question) UpdateRawAnswer(raw string, rawAnswer string) {
	question.Raw = helpers.ComprassHtml(raw)
	question.Content = question.Raw
	question.RawAnswer = question.RawAnswer
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

func AllQuestions(page, perPage int) []Question {
	var questions []Question
	DB.Model(&Question{}).Order("id desc", false).Limit(perPage).Offset(perPage * page).Find(&questions)
	return questions
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
