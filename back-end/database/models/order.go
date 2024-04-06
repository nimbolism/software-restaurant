package models

import "gorm.io/gorm"

// Order represents order data in the database
type Order struct {
	gorm.Model
	UserID     uint       `gorm:"foreignKey:UserID"`
	User       User       // Belongs To Relationship
	Foods      []Food     `gorm:"many2many:order_foods"`       // Many-to-Many Relationship with Food
	SideDishes []SideDish `gorm:"many2many:order_side_dishes"` // Many-to-Many Relationship with SideDish
	Quantity   int
	TotalPrice float64
	Status     string // Pending, Confirmed, Delivered, etc.
}