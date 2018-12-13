package models

type Bank struct {
	Model
	Name          string         `gorm:"index" json:"name"`
	Coment        string         `json:"coment"`
	QuestionTypes []QuestionType `json:"question_types,omitempty"`
}

func NewBank(name string, coment string, qts []QuestionType) *Bank {
	return &Bank{
		Name:          name,
		Coment:        coment,
		QuestionTypes: qts,
	}
}

func (bank *Bank) Save() (err error) {
	if DB.NewRecord(bank) {
		err = DB.Create(bank).Error
	} else {
		err = DB.Save(bank).Error
	}
	return
}

func (bank *Bank) Update() (err error) {
	var b Bank
	tx := DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.First(&b, "id = ?", bank.ID).Error
	if err != nil {
		return
	}
	if b.Name != bank.Name || b.Coment != bank.Coment {
		b.Name = bank.Name
		b.Coment = bank.Coment
		err = tx.Save(&b).Error
		if err != nil {
			return
		}
	}
	var qts []QuestionType
	err = tx.Model(&b).Related(&qts).Error
	if err != nil {
		return
	}
	for _, nqt := range bank.QuestionTypes {
		if nqt.ID == 0 {
			nqt.BankId = b.ID
			err = tx.Create(&nqt).Error
			if err != nil {
				return
			}
		} else {
			for _, eqt := range qts {
				if nqt.ID == eqt.ID && (nqt.Name != eqt.Name || nqt.Title != eqt.Title || nqt.Sort != eqt.Sort) {
					eqt.Name = nqt.Name
					eqt.Title = nqt.Title
					eqt.Sort = nqt.Sort
					err = tx.Save(&eqt).Error
					if err != nil {
						return
					}
				}
			}
		}
	}
	for _, eqt := range qts {
		var needDelete bool = true
		for _, nqt := range bank.QuestionTypes {
			if nqt.ID == eqt.ID {
				needDelete = false
			}
		}
		if needDelete {
			err = tx.Delete(&eqt).Error
			return
		}
	}
	return nil
}

func AllBanks() []Bank {
	var banks []Bank
	DB.Find(&banks)
	return banks
}

func GetBankById(id string) (bank *Bank, exists bool) {
	bank = new(Bank)
	if DB.First(bank, "id = ?", id).RecordNotFound() {
		return bank, false
	}
	DB.Model(bank).Related(&bank.QuestionTypes)
	return bank, true
}

func DeleteBankById(id string) error {
	if err := DB.Delete(&Bank{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
