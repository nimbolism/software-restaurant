package models

import "gorm.io/gorm"

// Food represents food data in the database
type Food struct {
	gorm.Model
	Name        string
	Description string
	CategoryID  uint
	Category    Category // Belongs To Relationship
	MealID      uint
	Meal        Meal // Belongs To Relationship
}

// Category represents food category data in the database
type Category struct {
	gorm.Model
	Name  string
	Foods []Food `gorm:"foreignKey:CategoryID"` // Has Many Relationship
}

// Meal represents meal data in the database
type Meal struct {
	gorm.Model
	Name  string
	Foods []Food `gorm:"foreignKey:MealID"` // Has Many Relationship
}

// SideDish represents side dish data in the database
type SideDish struct {
	gorm.Model
	Name        string
	Description string
}
