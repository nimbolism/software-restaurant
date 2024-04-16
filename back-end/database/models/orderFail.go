package models

import "gorm.io/gorm"

type OrderFail struct {
	gorm.Model
	Username string
	Error    string
}
