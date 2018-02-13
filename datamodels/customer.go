package datamodels

import (
	"time"
)

// Customer model struct
type Customer struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Address1  string    `gorm:"column:address_1" sql:"type:text" json:"address_1"`
	Address2  string    `gorm:"column:address_2" sql:"type:text" json:"address_2"`
	Postcode  string    `json:"postcode"`
	Phone     string    `json:"phone"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   uint8     `json:"deleted"`
}
