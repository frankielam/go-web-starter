package entity

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	Name      string `gorm:"unique;column:name" json:"name"`
	Password  string `gorm:"column:password" json:"password"`
	Mobile    string `gorm:"column:mobile" json:"mobile"`
	Email     string `gorm:"column:email" json:"email"`
	AvatarUrl string `gorm:"column:avatar_url" json:"avatar_url"`
	Language  string `gorm:"column:language" json:"language"`
	Role      int    `gorm:"column:role" json:"role"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
