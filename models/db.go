package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "datas.db")
	if err != nil {
		panic(err.Error())
	}

	if gin.Mode() != gin.ReleaseMode {
		db.LogMode(true)
	}

	DB = db
}

func init()  {
	DB.AutoMigrate(&Bank{})
}
