package controller

import (
	model "go-studio/models"

	"gorm.io/gorm"
)

func CreateAnswer(db *gorm.DB, userID uint, questionID uint, content string) error {
	answer := &model.Answer{
		UserID:     userID,
		QuestionID: questionID,
		Content:    content,
	}
	if err := db.Create(answer).Error; err != nil {
		return err
	}
	return nil
}
