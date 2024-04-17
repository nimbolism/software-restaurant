package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model  `gorm:"primaryKey=card_id"`
	UserID      uint `gorm:"foreignKey:UserID"`
	Reserves    int
	BlackListed bool
	Verified    bool // states who can reserve
	AccessLevel int  // 1 => ordinary, 2 => helper, 3 => admin
}

// trigger for automatically black listing users
func (c *Card) BeforeSave(tx *gorm.DB) (err error) {
	c.BlackListed = c.Reserves >= 3
	return
}
