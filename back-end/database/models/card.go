package models

import "gorm.io/gorm"

// Card represents card data in the database
type Card struct {
	gorm.Model  `gorm:"primaryKey=card_id"`
	UserID      uint `gorm:"foreignKey:UserID"`
	Reserves    int
	BlackListed bool
	AccessLevel int
}

func (c *Card) BeforeSave(tx *gorm.DB) (err error) {
	c.BlackListed = c.Reserves > 3
	return
}
