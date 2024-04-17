package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name         string `gorm:"unique"`
	Description  string
	CategoryID   uint
	MealID       uint
	CategoryName string `gorm:"-"`
	MealName     string `gorm:"-"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type Meal struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type SideDish struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
}
