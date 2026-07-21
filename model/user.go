package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"` // 用户名
	Password string `gorm:"type:varchar(255);not null" json:"-"`                   // 密码不返回给前端
	Avatar   string `gorm:"type:varchar(255)" json:"avatar"`                       // 头像URL
	Role     int    `gorm:"default:1" json:"role"`                                 // 1:普通用户 2:管理员
}
