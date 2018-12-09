package models

type Bank struct {
	Model
	Name string	`gorm:"unique_index" json:"name"`
	Coment string	`json:"coment"`
}

func NewBank(name string,coment string) *Bank  {
	return &Bank{
		Name:name,
		Coment:coment,
	}
}

func (bank *Bank)Save()  {
	if DB.NewRecord(bank) {
		DB.Create(bank)
	} else {
		DB.Save(bank)
	}
}

func AllBanks()[]Bank {
	var banks []Bank
	DB.Find(&banks)
	return banks
}

func GetBankById(id uint) (bank *Bank, exists bool){
	bank = new(Bank)
	if DB.First(bank,"id = ?",id).RecordNotFound() {
		return bank,false
	}
	return bank,true
}