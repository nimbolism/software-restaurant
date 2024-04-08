package models

import "gorm.io/gorm"

// Food represents food data in the database
type Food struct {
	gorm.Model
	Name         string `gorm:"unique"`
	Description  string
	CategoryID   uint
	MealID       uint
	CategoryName string `gorm:"-"` // Add this line
	MealName     string `gorm:"-"` // Add this line
}

// Category represents food category data in the database
type Category struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// Meal represents meal data in the database
type Meal struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// SideDish represents side dish data in the database
type SideDish struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
}
