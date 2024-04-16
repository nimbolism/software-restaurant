package models

import "gorm.io/gorm"

// Order represents order data in the database
type Order struct {
	gorm.Model
	UserID     uint       `gorm:"column:user_id"`
	User       User       `gorm:"foreignKey:UserID"`
	Foods      []Food     `gorm:"many2many:order_foods"`       // Many-to-Many Relationship with Food
	SideDishes []SideDish `gorm:"many2many:order_side_dishes"` // Many-to-Many Relationship with SideDish
	Paid       bool       //if false => reserved
}
