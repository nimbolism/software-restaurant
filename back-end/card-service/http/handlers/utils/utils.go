package utils

import (
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"gorm.io/gorm"
)

func FindCardByUserID(db *gorm.DB, userID uint) (*models.Card, error) {
	var card models.Card
    if err := db.Where("user_id = ?", userID).First(&card).Error; err != nil {
        return nil, err
    }
    return &card, nil
} 