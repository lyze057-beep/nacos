package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);unique;not null"`
	Password string `gorm:"type:varchar(64);not null"`
	Phone    string `gorm:"type:varchar(20);unique;not null"`
}
