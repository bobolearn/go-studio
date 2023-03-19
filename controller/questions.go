package controller

import (
	model "go-studio/models"

	"gorm.io/gorm"
)

func CreateQuestion(db *gorm.DB, userID uint, title string, content string) (*model.Question, error) {
	question := &model.Question{
		UserID:  userID,
		Title:   title,
		Content: content,
	}
	if err := db.Create(question).Error; err != nil {
		return nil, err
	}
	return question, nil
}
func QueryAllQuestions(db *gorm.DB) ([]*model.Question, error) {
	var questions []*model.Question
	if err := db.Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}
func QueryOneQuestion(db *gorm.DB, questionID uint64) (*model.Question, error) {
	question := &model.Question{}
	if err := db.First(question, questionID).Error; err != nil {
		return nil, err
	}
	return question, nil
}

func DeleteQuestion(db *gorm.DB, questionID uint64) error {
	question := &model.Question{}
	if err := db.First(question, questionID).Error; err != nil {
		return err
	}
	if err := db.Delete(question).Error; err != nil {
		return err
	}
	return nil
}

func UpdateQuestion(db *gorm.DB, questionID uint64, title string, content string) error {
	question := &model.Question{}
	if err := db.First(question, questionID).Error; err != nil {
		return err
	}
	question.Title = title
	question.Content = content
	if err := db.Save(question).Error; err != nil {
		return err
	}
	return nil
}
