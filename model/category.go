package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"` // 板块名
	Description string `gorm:"type:varchar(255)" json:"description"`              // 板块介绍
	Sort        int    `gorm:"default:0" json:"sort"`                             // 排序权重
}
