package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	QRCode       []byte
	PhoneNumber  string `gorm:"unique"`
	NationalCode string `gorm:"unique"`
	Password     string
}
