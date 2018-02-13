package services

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//For connect mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//InitDb is for connecting to database
func InitDb() *gorm.DB {
	var err error
	// var customer models.Customer

	db, err := gorm.Open("mysql", "root:sekarang9@tcp(127.0.0.1:3306)/dino?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	return db
}
