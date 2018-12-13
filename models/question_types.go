package models

type QuestionType struct {
	Model
	Name   string `json:"name"`
	Title  string `json:"title"`
	Sort   uint   `json:"sort"`
	BankId uint   `json:"bank_id"`
}

func NewQuestionType(name string, title string, bankId uint) *QuestionType {
	return &QuestionType{
		Name:   name,
		Title:  title,
		BankId: bankId,
	}
}

func (qt *QuestionType) Save() (err error) {
	if DB.NewRecord(qt) {
		err = DB.Create(qt).Error
	} else {
		err = DB.Save(qt).Error
	}
	return
}

func AllQuestionTypes(bankId string) []QuestionType {
	var qts []QuestionType
	DB.Find(&qts, "bank_id = ?", bankId)
	return qts
}

func GetQuestionTypeById(id uint) (qt *QuestionType, exists bool) {
	qt = new(QuestionType)
	if DB.First(qt, "id = ?", id).RecordNotFound() {
		return qt, false
	}
	return qt, true
}

func DeleteQuestionTypeById(id string) error {
	if err := DB.Delete(&QuestionType{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func DeleteQuestionTypeByBankId(id string) error {
	if err := DB.Delete(&QuestionType{}, "bank_id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func CreateQuestionTypes(bankId uint, qts []QuestionType) (err error) {
	tx := DB.Begin()
	for _, qt := range qts {
		qt.ID = 0
		qt.BankId = bankId
		err = tx.Create(&qt).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
