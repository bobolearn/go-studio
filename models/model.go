package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm: primaryKey;autoIncrement:"true";`
	Name     string `gorm: "not null"`
	Email    string `gorm: "not null"`
	Password string `gorm: "not null"`
}
type Question struct {
	gorm.Model
	Title   string `gorm: "not null"`
	Content string `gorm: "not null"`
	UserID  uint   `gorm: "not null"`
	User    User   `gorm: "constraint:OnDelete:CASCADE;OnUpdate:CASCADE;"`
}
type Answer struct {
	gorm.Model
	Content    string   `gorm: "not null"`
	UserID     uint     `gorm: "not null"`
	QuestionID uint     `gorm: "not null"`
	User       User     `gorm: "constraint:OnDelete:CASCADE;OnUpdate:CASCADE;"`
	Question   Question `gorm: "constraint:OnDelete:CASCADE;OnUpdate:CASCADE;"`
}
