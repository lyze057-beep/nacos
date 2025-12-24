package dao

import (
	"5/exam/nacos/config"
	"5/exam/nacos/handler/model"
	"errors"

	"gorm.io/gorm"
)

type Register struct{}
type Login struct{}

func (l *Register) CreateDao(users *model.Users) error {
	return config.DB.Create(users).Error
}
func (l *Register) GetUserByDao(username string) (*model.Users, error) {
	var user model.Users
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
