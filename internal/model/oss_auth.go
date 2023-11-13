package model

import "gorm.io/gorm"

type OssAuth struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;column:id"`
	Username string `gorm:"column:username"`
}

func (OssAuth) TableName() string {
	return "oss_auth"
}
