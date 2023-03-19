package controller

import (
	"errors"
	model "go-studio/models"
	utils "go-studio/utils"

	"gorm.io/gorm"
)

// 注册
func RegisterController(db *gorm.DB, name string, password string) error {
	var count int64
	if err := db.Model(&model.User{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户名已存在")
	}

	// 创建用户
	user := &model.User{
		Name:     name,
		Password: utils.HashPassword(password),
	}
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// 登录

func Login(db *gorm.DB, name, password string) (string, error) {
	var user model.User
	if err := db.Where("name =?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("用户不存在")
		}
		return "", err
	}
	// 验证密码
	if !utils.ComparePassword(password, user.Password) {
		return "", errors.New("密码错误")
	}
	// 生成 token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
