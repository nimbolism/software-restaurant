package models

import "gorm.io/gorm"

// this is to log failed orders - to show the user their errors
type OrderFail struct {
	gorm.Model
	Username string
	Error    string
}
