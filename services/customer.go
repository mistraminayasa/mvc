package services

import (
	models "mvc/datamodels"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql" // import mysql
	"github.com/kataras/iris"
)

// ResponseDelete give a message when data deleted
type ResponseDelete struct {
	Message string `json:"message"`
}

// GetCustomers Restapi
func GetCustomers(ctx iris.Context) {
	db := InitDb()
	defer db.Close()

	var customer []models.Customer

	db.Find(&customer)
	ctx.JSON(customer)
}

// GetCustomer get detail customer
func GetCustomer(ctx iris.Context) {
	db := InitDb()
	defer db.Close()

	var customer models.Customer
	id, _ := ctx.Params().GetInt("id")
	db.Where("id = ?", id).Find(&customer)
	ctx.JSON(customer)
}

// CreateCustomer end point
func CreateCustomer(ctx iris.Context) {
	db := InitDb()
	defer db.Close()

	var customer models.Customer
	ctx.ReadJSON(&customer)
	db.Create(&customer)
	db.Save(&customer)
	ctx.JSON(customer)
}

//UpdateCustomer update or edit customer endpoint
func UpdateCustomer(ctx iris.Context) {
	db := InitDb()
	defer db.Close()

	var customer models.Customer
	id, _ := ctx.Params().GetInt("id")
	db.Where("id = ?", id).First(&customer)
	ctx.ReadJSON(&customer)
	db.Save(&customer)
	ctx.JSON(customer)
}

//DeleteCustomer endpoint
func DeleteCustomer(ctx iris.Context) {
	db := InitDb()
	defer db.Close()

	var customer models.Customer
	var rd ResponseDelete
	id, _ := ctx.Params().GetInt("id")
	d := db.Where("id = ?", id).Delete(&customer)
	fmt.Println(d)

	rd.Message = "deleted"
	ctx.JSON(rd)
}
