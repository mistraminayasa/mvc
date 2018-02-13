package datamodels

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import mysql for connect to database
)

var db *gorm.DB
var err error

func init() {

	db, err = gorm.Open("mysql", "root:sekarang9@tcp(127.0.0.1:3306)/dino?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(false)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Customer{})
}
